package test

import (
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func TestGetNestVal1(t *testing.T) {
	var find bool
	var actual interface{}

	noinfo := NumberInfo{
		ID:        123,
		DisplayNo: "789",
		ChannelID: 999,
	}
	wrapinfo := WrapperInfo{
		NumberInfo: noinfo,
		Data:       "wrap123",
	}
	target := &cominfo{
		NumberInfo: noinfo,
		Data:       "data123",
		Timestamp:  "123456789",
		Info:       wrapinfo,
	}
	find, actual = voluator.GetValue(target, "data", false)
	if !find {
		t.Fail()
	}
	if actual != "data123" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "display_number", false)
	if !find {
		t.Fail()
	}
	if actual != "789" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "channel_id", false)
	if !find {
		t.Fail()
	}
	if actual != 999 {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "wrap.data", false)
	if !find {
		t.Fail()
	}
	if actual != "wrap123" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "wrap.display_number", false)
	if !find {
		t.Fail()
	}
	if actual != "789" {
		t.Fail()
	}
}

func TestGetNestValWithCache1(t *testing.T) {
	var find bool
	var actual interface{}

	noinfo := NumberInfo{
		ID:        123,
		DisplayNo: "789",
		ChannelID: 999,
	}
	wrapinfo := WrapperInfo{
		NumberInfo: noinfo,
		Data:       "wrap123",
	}
	target := &cominfo{
		NumberInfo: noinfo,
		Data:       "data123",
		Timestamp:  "123456789",
		Info:       wrapinfo,
	}
	find, actual = voluator.GetValue(target, "data", true)
	if !find {
		t.Fail()
	}
	if actual != "data123" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "display_number", true)
	if !find {
		t.Fail()
	}
	if actual != "789" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "channel_id", true)
	if !find {
		t.Fail()
	}
	if actual != 999 {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "wrap.data", true)
	if !find {
		t.Fail()
	}
	if actual != "wrap123" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "wrap.display_number", true)
	if !find {
		t.Fail()
	}
	if actual != "789" {
		t.Fail()
	}
}
