package voluator

import (
	"fmt"
	"strings"
	"testing"
)

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

func Benchmark_StringJoin(b *testing.B){
	h1 := "com.d9cloud.callcenter.common"
	h2 := "voiceNotifyRequest"
	for i:=0;i<b.N;i++{
		strings.Join([]string{h1, h2}, ".")
	}
}

func Benchmark_StringComb(b *testing.B){
	h1 := "com.d9cloud.callcenter.common"
	h2 := "voiceNotifyRequest"
	len1 := len(h1)
	len2 := len(h2)
	for i:=0;i<b.N;i++{
		buff := make([]byte,len1+len2+1)
		copy(buff,h1)
		buff[len1] = '.'
		copy(buff[len1+1:],h2)
		fmt.Println(buff)
	}
}
