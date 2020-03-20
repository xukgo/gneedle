package test

import (
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func TestSetNestValCached1(t *testing.T) {
	var find bool
	var err error

	noinfo := NumberInfo{}
	wrapinfo := WrapperInfo{}
	target := &cominfo{
		NumberInfo: noinfo,
		Data:       "",
		Timestamp:  "",
		Info:       wrapinfo,
	}

	find, err = voluator.AssignValue(target, "data", nil, true)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "data", "123abc", true)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Data != "123abc" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "display_number", "123abc", true)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.DisplayNo != "123abc" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "wrap.data", "123abc", true)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Info.Data != "123abc" {
		t.Fail()
	}

	find, err = voluator.AssignValue(target, "wrap.channel_id", 999, true)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
	if target.Info.ChannelID != 999 {
		t.Fail()
	}
}
