package dyTypeUtil

import (
	"testing"
)

func Benchmark_getKindtype(b *testing.B) {
	for i:=0;i<b.N;i++{
		getValueKindType("")
	}
}
func Benchmark_getKind(b *testing.B) {
	for i:=0;i<b.N;i++{
		getKind("")
	}
}

func Benchmark_Convert_int_uint64(b *testing.B) {
	for i:=0;i<b.N;i++{
		_, err := ConvertBasicTypeData(int(7), uint64(1))
		if err != nil {
			b.Fail()
		}
	}
}

func Benchmark_Convert_string_bool(b *testing.B) {
	for i:=0;i<b.N;i++{
		_, err := ConvertBasicTypeData("true", true)
		if err != nil {
			b.Fail()
		}
	}
}

func Benchmark_Convert_bool_string(b *testing.B) {
	for i:=0;i<b.N;i++{
		_, err := ConvertBasicTypeData(true, "")
		if err != nil {
			b.Fail()
		}
	}
}

func Benchmark_Convert_float_bool(b *testing.B) {
	for i:=0;i<b.N;i++{
		_, err := ConvertBasicTypeData(-1.125, true)
		if err != nil {
			b.Fail()
		}
	}
}