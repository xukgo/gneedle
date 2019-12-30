/*
@Time : 2019/9/30 18:41
@Author : Hermes
@File : getvalue_test
@Description:
*/
package test

import (
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func TestGetWithCacheOneLayer(t *testing.T) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	find, actual = voluator.GetValue(target, "id", true)
	if !find {
		t.Fail()
	}
	if actual != 9201 {
		t.Fail()
	}
	find, actual = voluator.GetValue(&target, "id", true)
	if !find {
		t.Fail()
	}
	if actual != 9201 {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "data", true)
	if !find {
		t.Fail()
	}
	if actual != "mydata" {
		t.Fail()
	}

	find, actual = voluator.GetValue(&target, "data", true)
	if !find {
		t.Fail()
	}
	if actual != "mydata" {
		t.Fail()
	}
}

func TestGetMultiLayerSliceItemWithCache(t *testing.T) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	find, actual = voluator.GetValue(target, "sa[1]", true)
	if !find {
		t.Fail()
	}
	if actual != "bbb" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "subject.application.NotifyBlock[0].id", true)
	if !find {
		t.Fail()
	}
	if actual != 101 {
		t.Fail()
	}

	find, actual = voluator.GetValue(target, "subject.application.NotifyBlock[2].name", true)
	if !find {
		t.Fail()
	}
	if actual != "blname3" {
		t.Fail()
	}

}

func TestGetMultiLayerWithCache(t *testing.T) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	find, actual = voluator.GetValue(target, "subject.application.AppID", true)
	if !find {
		t.Fail()
	}
	if actual != 4001 {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "subject.application.AppName", true)
	if !find {
		t.Fail()
	}
	if actual != "axb" {
		t.Fail()
	}

	for i := 0; i < 1000000; i++ {
		find, actual = voluator.GetValue(target, "subject.application.Callback", true)
		if !find {
			t.Fail()
		}
		if actual != "http://ip:port/url" {
			t.Fail()
		}
	}
}

func BenchmarkGetWithCache(b *testing.B) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	for i := 0; i < b.N; i++ {
		find, actual = voluator.GetValue(target, "subject.application.Callback", true)
		if !find {
			b.Fail()
		}
		if actual != "http://ip:port/url" {
			b.Fail()
		}
	}
}
