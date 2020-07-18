/*
@Time : 2019/9/30 18:41
@Author : Hermes
@File : getvalue_test
@Description:
*/
package test

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func intReponseModel() rootResponse {
	ninfo := NumberInfo{
		ID:        999,
		DisplayNo: "123456789",
		ChannelID: 774,
	}

	blockInfo1 := notifyBlockInfo{
		ID:    101,
		Name:  "blname1",
		Level: 201,
	}
	blockInfo2 := notifyBlockInfo{
		ID:    102,
		Name:  "blname2",
		Level: 202,
	}
	blockInfo3 := notifyBlockInfo{
		ID:    103,
		Name:  "blname3",
		Level: 203,
	}

	appinfo := applicationInfo{
		AppId:                4001,
		AppName:              "axb",
		NotifyBlockInfoArray: []notifyBlockInfo{blockInfo1, blockInfo2, blockInfo3},
		Callback:             "http://ip:port/url",
	}
	model := rootResponse{
		ID:   9201,
		Gap:  -123.4567,
		Data: "mydata",
		Subject: subject{
			NumberInfo:      ninfo,
			ApplicationInfo: appinfo,
		},
		Timestamp: "2019100110121235",
		StrArr:    []string{"aaa", "bbb", "ccc"},
	}

	return model
}

func TestGetSliceItem(t *testing.T) {
	var arr []string
	val,find := voluator.GetSliceIndexItem(arr,0)
	if find{
		t.FailNow()
	}

	val,find = voluator.GetSliceIndexItem(arr,-1)
	if find{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(arr,1)
	if find{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(nil,-1)
	if find{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(nil,0)
	if find{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(nil,1)
	if find{
		t.FailNow()
	}

	arr = []string{"0","1","2"}
	val,find = voluator.GetSliceIndexItem(arr,0)
	if val != "0"{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(arr,1)
	if val != "1"{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(arr,2)
	if val != "2"{
		t.FailNow()
	}
	val,find = voluator.GetSliceIndexItem(arr,3)
	if find{
		t.FailNow()
	}
}
func TestGetOneLayer(t *testing.T) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	find, actual = voluator.GetValue(target, "id", false)
	if !find {
		t.Fail()
	}
	if actual != 9201 {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "data", false)
	if !find {
		t.Fail()
	}
	if actual != "mydata" {
		t.Fail()
	}

	find, actual = voluator.GetValue(&target, "id", false)
	if !find {
		t.Fail()
	}
	if actual != 9201 {
		t.Fail()
	}
	find, actual = voluator.GetValue(&target, "data", false)
	if !find {
		t.Fail()
	}
	if actual != "mydata" {
		t.Fail()
	}
}

func TestGetMultiLayerSliceItem(t *testing.T) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	find, actual = voluator.GetValue(target, "sa[1]", false)
	if !find {
		t.Fail()
	}
	if actual != "bbb" {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "subject.application.NotifyBlock[0].id", false)
	if !find {
		t.Fail()
	}
	if actual != 101 {
		t.Fail()
	}

	find, actual = voluator.GetValue(target, "subject.application.NotifyBlock[2].name", false)
	if !find {
		t.Fail()
	}
	if actual != "blname3" {
		t.Fail()
	}

}

func TestCheckGetPathValid(t *testing.T) {
	if !voluator.CheckGetPathValid("aaa1") {
		t.Fail()
	}
	if !voluator.CheckGetPathValid("a2aa.bbb2") {
		t.Fail()
	}
	if !voluator.CheckGetPathValid("aaa.bbb[3].ccc[2]") {
		t.Fail()
	}
	if !voluator.CheckGetPathValid("aaa.bbb.ccc[2]") {
		t.Fail()
	}

	if voluator.CheckGetPathValid("aa*a.bbb.ccc[2]") {
		t.Fail()
	}
	if voluator.CheckGetPathValid("aaa.bbb.ccc[p]") {
		t.Fail()
	}
}

func TestGetMultiLayer(t *testing.T) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	find, actual = voluator.GetValue(target, "subject.application.AppID", false)
	if !find {
		t.Fail()
	}
	if actual != 4001 {
		t.Fail()
	}
	find, actual = voluator.GetValue(target, "subject.application.AppName", false)
	if !find {
		t.Fail()
	}
	if actual != "axb" {
		t.Fail()
	}

	for i := 0; i < 1000000; i++ {
		find, actual = voluator.GetValue(target, "subject.application.Callback", false)
		if !find {
			t.Fail()
		}
		if actual != "http://ip:port/url" {
			t.Fail()
		}
	}
}

func BenchmarkGet1(b *testing.B) {
	var find bool
	var actual interface{}

	target := intReponseModel()
	for i := 0; i < b.N; i++ {
		find, actual = voluator.GetValue(target, "subject.application.Callback", false)
		if !find {
			b.Fail()
		}
		if actual != "http://ip:port/url" {
			b.Fail()
		}
	}
}

func BenchmarkGetByGjson(b *testing.B) {
	target := intReponseModel()
	bs, _ := jsoniter.Marshal(target)
	json := string(bs)
	for i := 0; i < b.N; i++ {
		value := gjson.Get(json, "subject.application.Callback")
		if !value.Exists() {
			b.Fail()
		}
		if value.String() != "http://ip:port/url" {
			b.Fail()
		}
	}
}
