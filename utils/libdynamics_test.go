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

package utils

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/ericlagergren/decimal"
)

func TestCloneDynamicStringsSliceOpt(t *testing.T) {
	in := []*DynamicStringSliceOpt{
		{
			Value:     []string{"VAL_1", "VAL_2"},
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     []string{"VAL_3", "VAL_4"},
			FilterIDs: []string{"fltr2"},
		},
	}

	clone := CloneDynamicStringSliceOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicStringOpt(t *testing.T) {
	in := []*DynamicStringOpt{
		{
			Value:     "VAL_1",
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     "VAL_2",
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicStringOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicInterfaceOpt(t *testing.T) {
	in := []*DynamicInterfaceOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     "VAL_2",
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicInterfaceOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicBoolOpt(t *testing.T) {
	in := []*DynamicBoolOpt{
		{
			Value:     true,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     false,
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicBoolOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicIntOpt(t *testing.T) {
	in := []*DynamicIntOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicIntOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicFloat64Opt(t *testing.T) {
	in := []*DynamicFloat64Opt{
		{
			Value:     1.3,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2.7,
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicFloat64Opt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicDurationOpt(t *testing.T) {
	in := []*DynamicDurationOpt{
		{
			Value:     time.Second * 2,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     time.Second * 5,
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicDurationOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

func TestCloneDynamicDecimalBigOpt(t *testing.T) {
	in := []*DynamicDecimalBigOpt{
		{
			Value:     new(decimal.Big).SetUint64(10),
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     new(decimal.Big).SetUint64(2),
			FilterIDs: []string{"fltr2"},
		},
	}
	clone := CloneDynamicDecimalBigOpt(in)
	if !reflect.DeepEqual(in, clone) {
		t.Error("Expected objects to match")
	}
}

// Equals
func TestDynamicStringSliceOptEqual(t *testing.T) {
	v1 := []*DynamicStringSliceOpt{
		{
			Value:     []string{"VAL_1", "VAL_2"},
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     []string{"VAL_3", "VAL_4"},
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicStringSliceOpt{
		{
			Value:     []string{"VAL_1", "VAL_2"},
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     []string{"VAL_3", "VAL_4"},
			FilterIDs: []string{"fltr2"},
		},
	}

	//Test if equal
	if !DynamicStringSliceOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	//Test if different
	v1[0].Value = append(v1[0].Value, "VAL_3")
	if DynamicStringSliceOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicStringSliceOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicStringSliceOpt{
		Value:     []string{"VAL_1", "VAL_2"},
		FilterIDs: []string{"fltr1"},
	})
	if DynamicStringSliceOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicStringOptEqual(t *testing.T) {
	v1 := []*DynamicStringOpt{
		{
			Value:     "VAL_1",
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     "VAL_2",
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicStringOpt{
		{
			Value:     "VAL_1",
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     "VAL_2",
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicStringOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = "VAL_3"
	if DynamicStringOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicStringOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicStringOpt{
		Value:     "NEW_VAL",
		FilterIDs: []string{"fltr1"},
	})
	if DynamicStringOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicBoolOptEquals(t *testing.T) {
	v1 := []*DynamicBoolOpt{
		{
			Value:     true,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     false,
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicBoolOpt{
		{
			Value:     true,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     false,
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicBoolOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = false
	if DynamicBoolOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicBoolOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicBoolOpt{
		Value:     true,
		FilterIDs: []string{"fltr1"},
	})
	if DynamicBoolOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicIntOptEqual(t *testing.T) {
	v1 := []*DynamicIntOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicIntOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicIntOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = 2
	if DynamicIntOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicIntOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicIntOpt{
		Value:     2,
		FilterIDs: []string{"fltr1"},
	})
	if DynamicIntOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicFloat64OptEqual(t *testing.T) {
	v1 := []*DynamicFloat64Opt{
		{
			Value:     1.2,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2.6,
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicFloat64Opt{
		{
			Value:     1.2,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2.6,
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicFloat64OptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = 2.8
	if DynamicFloat64OptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicFloat64OptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicFloat64Opt{
		Value:     3.5,
		FilterIDs: []string{"fltr1"},
	})
	if DynamicFloat64OptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicDurationOptEquals(t *testing.T) {
	v1 := []*DynamicDurationOpt{
		{
			Value:     time.Second * 2,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     time.Second * 5,
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicDurationOpt{
		{
			Value:     time.Second * 2,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     time.Second * 5,
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicDurationOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = time.Second * 11
	if DynamicDurationOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicDurationOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicDurationOpt{
		Value:     time.Second * 2,
		FilterIDs: []string{"fltr1"},
	})
	if DynamicDurationOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicDecimalBigOptEquals(t *testing.T) {
	v1 := []*DynamicDecimalBigOpt{
		{
			Value:     new(decimal.Big).SetUint64(10),
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     new(decimal.Big).SetUint64(2),
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicDecimalBigOpt{
		{
			Value:     new(decimal.Big).SetUint64(10),
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     new(decimal.Big).SetUint64(2),
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicDecimalBigOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = new(decimal.Big).SetUint64(16)
	if DynamicDecimalBigOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicDecimalBigOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicDecimalBigOpt{
		Value:     new(decimal.Big).SetUint64(10),
		FilterIDs: []string{"fltr1"},
	})
	if DynamicDecimalBigOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

func TestDynamicInterfaceOptEqual(t *testing.T) {
	v1 := []*DynamicInterfaceOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	v2 := []*DynamicInterfaceOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	if !DynamicInterfaceOptEqual(v1, v2) {
		t.Error("Expected both slices to be the same")
	}

	v1[0].Value = 2
	if DynamicInterfaceOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different filters
	v1[0].FilterIDs = append(v1[0].FilterIDs, "new_fltr")
	if DynamicInterfaceOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}

	//Test if different lengths
	v1 = append(v1, &DynamicInterfaceOpt{
		Value:     2,
		FilterIDs: []string{"fltr1"},
	})
	if DynamicInterfaceOptEqual(v1, v2) {
		t.Error("Expected slices to differ")
	}
}

// Dynamic to Map
func TestDynamicStringSliceOptsToMap(t *testing.T) {
	in := []*DynamicStringSliceOpt{
		{
			Value:     []string{"VAL_1", "VAL_2"},
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     []string{"VAL_3", "VAL_4"},
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string][]string{
		"fltr1;fltr2": {"VAL_1", "VAL_2"},
		"fltr2":       {"VAL_3", "VAL_4"},
	}

	if mp := DynamicStringSliceOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicStringOptsToMap(t *testing.T) {
	in := []*DynamicStringOpt{
		{
			Value:     "VAL_1",
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     "VAL_2",
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]string{
		"fltr1;fltr2": "VAL_1",
		"fltr2":       "VAL_2",
	}

	if mp := DynamicStringOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicBoolOptsToMap(t *testing.T) {
	in := []*DynamicBoolOpt{
		{
			Value:     true,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     false,
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]bool{
		"fltr1;fltr2": true,
		"fltr2":       false,
	}

	if mp := DynamicBoolOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicIntOptsToMap(t *testing.T) {
	in := []*DynamicIntOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]int{
		"fltr1;fltr2": 1,
		"fltr2":       2,
	}

	if mp := DynamicIntOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicFloat64OptsToMap(t *testing.T) {
	in := []*DynamicFloat64Opt{
		{
			Value:     2.2,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     5.1,
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]float64{
		"fltr1;fltr2": 2.2,
		"fltr2":       5.1,
	}

	if mp := DynamicFloat64OptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicDurationOptsToMap(t *testing.T) {
	in := []*DynamicDurationOpt{
		{
			Value:     time.Second * 2,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     time.Second * 8,
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]string{
		"fltr1;fltr2": "2s",
		"fltr2":       "8s",
	}

	if mp := DynamicDurationOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicDecimalBigOptsToMap(t *testing.T) {
	in := []*DynamicDecimalBigOpt{
		{
			Value:     new(decimal.Big).SetUint64(21),
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     new(decimal.Big).SetUint64(3),
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]string{
		"fltr1;fltr2": "21",
		"fltr2":       "3",
	}

	if mp := DynamicDecimalBigOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

func TestDynamicInterfaceOptsToMap(t *testing.T) {
	in := []*DynamicInterfaceOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	exp := map[string]interface{}{
		"fltr1;fltr2": 1,
		"fltr2":       2,
	}

	if mp := DynamicInterfaceOptsToMap(in); !reflect.DeepEqual(exp, mp) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(mp))
	}
}

// Map to Dynamic

func TestMapToDynamicStringSliceOpts(t *testing.T) {
	exp := []*DynamicStringSliceOpt{
		{
			Value:     []string{"VAL_1", "VAL_2"},
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     []string{"VAL_3", "VAL_4"},
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string][]string{
		"fltr1;fltr2": {"VAL_1", "VAL_2"},
		"fltr2":       {"VAL_3", "VAL_4"},
	}
	rcv := MapToDynamicStringSliceOpts(in)
	sort.Slice(rcv, func(i, j int) bool {
		return rcv[i].Value[i] < rcv[j].Value[j]
	})
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}
}

func TestMapToDynamicStringOpts(t *testing.T) {
	exp := []*DynamicStringOpt{
		{
			Value:     "VAL_1",
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     "VAL_2",
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]string{
		"fltr1;fltr2": "VAL_1",
		"fltr2":       "VAL_2",
	}

	rcv := MapToDynamicStringOpts(in)
	sort.Slice(rcv, func(i, j int) bool {
		return rcv[i].Value < rcv[j].Value
	})
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}
}

func TestMapToDynamicBoolOpts(t *testing.T) {
	exp := []*DynamicBoolOpt{
		{
			Value:     true,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     false,
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]bool{
		"fltr1;fltr2": true,
		"fltr2":       false,
	}

	rcv := MapToDynamicBoolOpts(in)
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}
}

func TestMapToDynamicIntOpts(t *testing.T) {
	exp := []*DynamicIntOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]int{
		"fltr1;fltr2": 1,
		"fltr2":       2,
	}

	rcv := MapToDynamicIntOpts(in)
	sort.Slice(rcv, func(i, j int) bool {
		return rcv[i].Value < rcv[j].Value
	})
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}
}

func TestMapToDynamicFloat64Opts(t *testing.T) {
	exp := []*DynamicFloat64Opt{
		{
			Value:     2.2,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     5.1,
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]float64{
		"fltr1;fltr2": 2.2,
		"fltr2":       5.1,
	}

	rcv := MapToDynamicFloat64Opts(in)
	sort.Slice(rcv, func(i, j int) bool {
		return rcv[i].Value < rcv[j].Value
	})
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}
}

func TestMapToDynamicDecimalBigOpts(t *testing.T) {
	exp := []*DynamicDecimalBigOpt{
		{
			Value:     new(decimal.Big).SetUint64(3),
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     new(decimal.Big).SetUint64(21),
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]string{
		"fltr1;fltr2": "3",
		"fltr2":       "21",
	}

	rcv, err := MapToDynamicDecimalBigOpts(in)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}

	in = map[string]string{
		"fltr1;fltr2": "21c",
		"fltr2":       "3c",
	}
	errExp := "can't convert <21c> to decimal"
	_, err = MapToDynamicDecimalBigOpts(in)
	if err == nil || err.Error() != errExp {
		t.Errorf("Expected error: %s", errExp)
	}
}

func TestMapToDynamicDurationOpts(t *testing.T) {
	exp := []*DynamicDurationOpt{
		{
			Value:     time.Second * 2,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     time.Second * 8,
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]string{
		"fltr1;fltr2": "2s",
		"fltr2":       "8s",
	}
	rcv, err := MapToDynamicDurationOpts(in)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}

	in = map[string]string{
		"fltr1;fltr2": MetaUnlimited,
		"fltr2":       MetaUnlimited,
	}
}

func TestMapToDynamicInterfaceOpts(t *testing.T) {
	exp := []*DynamicInterfaceOpt{
		{
			Value:     1,
			FilterIDs: []string{"fltr1", "fltr2"},
		},
		{
			Value:     2,
			FilterIDs: []string{"fltr2"},
		},
	}

	in := map[string]interface{}{
		"fltr1;fltr2": 1,
		"fltr2":       2,
	}
	rcv := MapToDynamicInterfaceOpts(in)
	if !reflect.DeepEqual(exp, rcv) {
		t.Errorf("Expected %v \n but received \n %v", ToJSON(exp), ToJSON(rcv))
	}
}