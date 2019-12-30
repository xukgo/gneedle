package voluator

import "testing"

func TestCheckTagNodeFormat(t *testing.T) {
	if !checkIsTagNodeFormat("bjfsfs1g465d4gd") {
		t.Fail()
	}
	if !checkIsTagNodeFormat("name") {
		t.Fail()
	}
	if !checkIsTagNodeFormat("aaa[3]") {
		t.Fail()
	}
	if !checkIsTagNodeFormat("name2_xukuan09[3]") {
		t.Fail()
	}
	if checkIsTagNodeFormat("=bjfsfs1g465d4gd") {
		t.Fail()
	}
	if checkIsTagNodeFormat("http[aa]") {
		t.Fail()
	}
}
