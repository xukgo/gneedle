/*
@Time : 2019/9/30 14:18
@Author : Hermes
@File : constAssign_test
@Description:
*/
package test

import (
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func TestValAssignOneLayerString(t *testing.T) {
	var find bool
	var err error
	target := new(rootResponse)
	find, err = voluator.AssignValue(target, "data", []byte("helloworld"), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "helloworld" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", "mydata", false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "mydata" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", int8(9), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "9" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", int8(-79), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "-79" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", int64(22222222222222222), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "22222222222222222" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", int64(-46464646), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "-46464646" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", uint64(777777), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "777777" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", uint8(22), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "22" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", float64(799.99999), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "799.99999" {
		t.Fail()
	}
}

func TestValAssignOneLayerInt(t *testing.T) {
	var find bool
	var err error
	target := new(rootResponse)
	find, err = voluator.AssignValue(target, "id", "123", false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.ID != 123 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "id", int8(-79), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.ID != -79 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "id", int64(-46464646), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.ID != -46464646 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "id", uint8(22), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.ID != 22 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "id", uint64(4687664648), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.ID != 4687664648 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "id", float64(799.99999), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.ID != 799 {
		t.Fail()
	}
}

func TestValAssignOneLayerFloat(t *testing.T) {
	var find bool
	var err error
	target := new(rootResponse)
	find, err = voluator.AssignValue(target, "gap", "-998.897", false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Gap != -998.897 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "gap", int8(-79), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Gap != -79 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "gap", int64(-46464646), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Gap != -46464646 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "gap", uint8(22), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Gap != 22 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "gap", uint64(4687664648), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Gap != 4687664648 {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "gap", float64(799.99999), false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Gap != 799.99999 {
		t.Fail()
	}
}

func TestValAssignMultiLayer(t *testing.T) {
	var find bool
	var err error
	for i := 0; i < 1; i++ {
		target := new(rootResponse)
		find, err = voluator.AssignValue(target, "subject.application.AppName", "myname", false)
		if !find {
			t.Fail()
		}
		if err != nil {
			t.Error(err)
		}
		if target.Subject.ApplicationInfo.AppName != "myname" {
			t.Fail()
		}
	}

	target := new(rootResponse)
	find, err = voluator.AssignValue(target, "subject.application.AppName", 12.34, false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Subject.ApplicationInfo.AppName != "12.34" {
		t.Fail()
	}
}

func TestValAssignStringSlice(t *testing.T) {
	var find bool
	var err error
	target := new(rootResponse)
	find, err = voluator.AssignValue(target, "sa", []string{"aaa", "bbb", "ccc"}, false)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.StrArr[0] != "aaa" {
		t.Fail()
	}
	if target.StrArr[1] != "bbb" {
		t.Fail()
	}
	if target.StrArr[2] != "ccc" {
		t.Fail()
	}
}
