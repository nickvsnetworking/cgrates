/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package ers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cgrates/cgrates/agents"
	"github.com/cgrates/cgrates/config"
	"github.com/cgrates/cgrates/ees"
	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/utils"
	"github.com/nats-io/nats.go"
)

// NewNatsER return a new amqp event reader
func NewNatsER(cfg *config.CGRConfig, cfgIdx int,
	rdrEvents, partialEvents chan *erEvent, rdrErr chan error,
	fltrS *engine.FilterS, rdrExit chan struct{}) (_ EventReader, err error) {
	rdr := &NatsER{
		cgrCfg:        cfg,
		cfgIdx:        cfgIdx,
		fltrS:         fltrS,
		rdrEvents:     rdrEvents,
		partialEvents: partialEvents,
		rdrExit:       rdrExit,
		rdrErr:        rdrErr,
	}
	if concReq := rdr.Config().ConcurrentReqs; concReq != -1 {
		rdr.cap = make(chan struct{}, concReq)
		for i := 0; i < concReq; i++ {
			rdr.cap <- struct{}{}
		}
	}
	if err = rdr.processOpts(); err != nil {
		return
	}
	if err = rdr.createPoster(); err != nil {
		return
	}
	return rdr, nil
}

// NatsER implements EventReader interface for amqp message
type NatsER struct {
	// sync.RWMutex
	cgrCfg *config.CGRConfig
	cfgIdx int // index of config instance within ERsCfg.Readers
	fltrS  *engine.FilterS

	rdrEvents     chan *erEvent // channel to dispatch the events created to
	partialEvents chan *erEvent // channel to dispatch the partial events created to
	rdrExit       chan struct{}
	rdrErr        chan error
	cap           chan struct{}

	subject      string
	queueID      string
	jetStream    bool
	consumerName string
	opts         []nats.Option
	jsOpts       []nats.JSOpt

	poster *ees.NatsEE
}

// Config returns the curent configuration
func (rdr *NatsER) Config() *config.EventReaderCfg {
	return rdr.cgrCfg.ERsCfg().Readers[rdr.cfgIdx]
}

// Serve will start the gorutines needed to watch the nats subject
func (rdr *NatsER) Serve() (err error) {
	// Connect to a server
	var nc *nats.Conn
	var js nats.JetStreamContext

	if nc, err = nats.Connect(rdr.Config().SourcePath, rdr.opts...); err != nil {
		return
	}
	ch := make(chan *nats.Msg)
	if !rdr.jetStream {
		if _, err = nc.ChanQueueSubscribe(rdr.subject, rdr.queueID, ch); err != nil {
			return
		}
	} else {
		js, err = nc.JetStream(rdr.jsOpts...)
		if err != nil {
			return
		}
		if _, err = js.QueueSubscribe(rdr.subject, rdr.queueID, func(msg *nats.Msg) {
			ch <- msg
		}, nats.Durable(rdr.consumerName)); err != nil {
			return
		}
	}
	go func() {
		for {
			if rdr.Config().ConcurrentReqs != -1 {
				<-rdr.cap // do not try to read if the limit is reached
			}
			select {
			case <-rdr.rdrExit:
				utils.Logger.Info(
					fmt.Sprintf("<%s> stop monitoring nats path <%s>",
						utils.ERs, rdr.Config().SourcePath))
				nc.Drain()
				if rdr.poster != nil {
					rdr.poster.Close()
				}
				return
			case msg := <-ch:
				go func(msg *nats.Msg) {
					if err := rdr.processMessage(msg.Data); err != nil {
						utils.Logger.Warning(
							fmt.Sprintf("<%s> processing message %s error: %s",
								utils.ERs, string(msg.Data), err.Error()))
					}
					if rdr.poster != nil { // post it
						if err := ees.ExportWithAttempts(rdr.poster, msg.Data, utils.EmptyString); err != nil {
							utils.Logger.Warning(
								fmt.Sprintf("<%s> writing message %s error: %s",
									utils.ERs, string(msg.Data), err.Error()))
						}
					}
					if rdr.Config().ConcurrentReqs != -1 {
						rdr.cap <- struct{}{}
					}
				}(msg)
			}
		}
	}()
	return
}

func (rdr *NatsER) processMessage(msg []byte) (err error) {
	var decodedMessage map[string]interface{}
	if err = json.Unmarshal(msg, &decodedMessage); err != nil {
		return
	}
	agReq := agents.NewAgentRequest(
		utils.MapStorage(decodedMessage), nil,
		nil, nil, nil, rdr.Config().Tenant,
		rdr.cgrCfg.GeneralCfg().DefaultTenant,
		utils.FirstNonEmpty(rdr.Config().Timezone,
			rdr.cgrCfg.GeneralCfg().DefaultTimezone),
		rdr.fltrS, nil) // create an AgentRequest
	var pass bool
	if pass, err = rdr.fltrS.Pass(agReq.Tenant, rdr.Config().Filters,
		agReq); err != nil || !pass {
		return
	}
	if err = agReq.SetFields(rdr.Config().Fields); err != nil {
		return
	}
	cgrEv := utils.NMAsCGREvent(agReq.CGRRequest, agReq.Tenant, utils.NestingSep, agReq.Opts)
	rdrEv := rdr.rdrEvents
	if _, isPartial := cgrEv.APIOpts[utils.PartialOpt]; isPartial {
		rdrEv = rdr.partialEvents
	}
	rdrEv <- &erEvent{
		cgrEvent: cgrEv,
		rdrCfg:   rdr.Config(),
	}
	return
}

func (rdr *NatsER) createPoster() (err error) {
	processedOpt := getProcessOptions(rdr.Config().Opts)
	if len(processedOpt) == 0 &&
		len(rdr.Config().ProcessedPath) == 0 {
		return
	}
	rdr.poster, err = ees.NewNatsEE(&config.EventExporterCfg{
		ID: rdr.Config().ID,
		ExportPath: utils.FirstNonEmpty(
			rdr.Config().ProcessedPath, rdr.Config().SourcePath),
		Opts:           processedOpt,
		Attempts:       rdr.cgrCfg.GeneralCfg().PosterAttempts,
		FailedPostsDir: rdr.cgrCfg.GeneralCfg().FailedPostsDir,
	}, rdr.cgrCfg.GeneralCfg().NodeID,
		rdr.cgrCfg.GeneralCfg().ConnectTimeout, nil)
	return
}

func (rdr *NatsER) processOpts() (err error) {
	rdr.subject = utils.IfaceAsString(rdr.Config().Opts[utils.NatsSubject])
	rdr.queueID = utils.FirstNonEmpty(utils.IfaceAsString(rdr.Config().Opts[utils.NatsQueueID]),
		rdr.cgrCfg.GeneralCfg().NodeID)
	rdr.consumerName = utils.FirstNonEmpty(utils.IfaceAsString(rdr.Config().Opts[utils.NatsConsumerName]),
		utils.CGRateSLwr)
	if useJetStreamVal, has := rdr.Config().Opts[utils.NatsJetStream]; has {
		if rdr.jetStream, err = utils.IfaceAsBool(useJetStreamVal); err != nil {
			return
		}
	}
	if rdr.jetStream {
		if maxWaitVal, has := rdr.Config().Opts[utils.NatsJetStreamMaxWait]; has {
			var maxWait time.Duration
			if maxWait, err = utils.IfaceAsDuration(maxWaitVal); err != nil {
				return
			}
			rdr.jsOpts = []nats.JSOpt{nats.MaxWait(maxWait)}
		}
	}
	rdr.opts, err = ees.GetNatsOpts(rdr.Config().Opts,
		rdr.cgrCfg.GeneralCfg().NodeID,
		rdr.cgrCfg.GeneralCfg().ConnectTimeout)
	return
}
