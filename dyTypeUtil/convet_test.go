package dyTypeUtil

import (
	"testing"
)

func Test_Convert_float32_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(-75.21), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75.21) {
		t.Fail()
	}
}

func Test_Convert_float32_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(-75.21), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float32_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(-75.21), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float32_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(-75.21), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float32_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(-75.21), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float32_int(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(-75.21), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float32_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(75.21), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float32_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(75.21), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float32_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(75.21), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float32_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(75.21), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float32_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(float32(75.21), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float64_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(-75.21), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75.21) {
		t.Fail()
	}
}

func Test_Convert_float64_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(-75.21), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float64_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(-75.21), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float64_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(-75.21), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float64_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(-75.21), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float64_int(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(-75.21), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -75) {
		t.Fail()
	}
}

func Test_Convert_float64_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(75.21), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float64_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(75.21), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float64_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(75.21), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float64_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(75.21), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_float64_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(float64(75.21), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 75) {
		t.Fail()
	}
}

func Test_Convert_int8_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_int(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(-7), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int8_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(7), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int8_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(7), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int8_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(7), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int8_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(7), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int8_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(int8(7), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int64_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_int(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(-7), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int64_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(7), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int64_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(7), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int64_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(7), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int64_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(7), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int64_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(int64(7), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_int(t *testing.T) {
	val, err := ConvertBasicTypeData(int(-7), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, -7) {
		t.Fail()
	}
}

func Test_Convert_int_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(int(7), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(int(7), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(int(7), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(int(7), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_int_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(int(7), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_int(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint8_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(uint8(7), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint64_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint64_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(77), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 77) {
		t.Fail()
	}
}

func Test_Convert_uint64_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_int(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7) {
		t.Fail()
	}
}

func Test_Convert_uint64_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint64_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(uint64(7999), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 7999) {
		t.Fail()
	}
}

func Test_Convert_uint_float32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), float32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_float64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), float64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_int8(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), int8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_int16(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), int16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_int32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), int32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_int64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), int64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_int(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), int(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_uint8(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), uint8(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_uint16(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), uint16(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_uint32(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), uint32(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_uint64(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), uint64(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}

func Test_Convert_uint_uint(t *testing.T) {
	val, err := ConvertBasicTypeData(uint(17), uint(1))
	if err != nil {
		t.Fail()
	}
	if !Equal(val, 17) {
		t.Fail()
	}
}
