package test

import (
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func TestValAssign2(t *testing.T) {
	var find bool
	var err error
	target := new(SessionCache)
	find, err = voluator.AssignValue(target, "isMeeting", true, true)
	if !find {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
}