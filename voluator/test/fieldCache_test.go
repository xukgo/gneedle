package test

import (
	"fmt"
	"github.com/xukgo/gneedle/voluator"
	"testing"
)

func TestGetClassPath(t *testing.T) {
	model := new(cominfo)
	fmt.Println(voluator.GetClassPath(model))
}

func Benchmark_GetClassPath(b *testing.B) {
	model := new(cominfo)

	for i := 0; i < b.N; i++ {
		path :=(voluator.GetClassPath(model))
		if len(path) == 0{
			b.Fail()
		}
	}
}