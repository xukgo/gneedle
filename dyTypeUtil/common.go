package dyTypeUtil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const numberic_min_threshold float64 = 0.01

func checkIsBoolAndString(t1 int, t2 int) bool {
	if t1 == BOOL_TYPE && t2 == STRING_TYPE {
		return true
	}
	if t2 == BOOL_TYPE && t1 == STRING_TYPE {
		return true
	}
	return false
}

func checkIsBoolAndNumber(t1 int, t2 int) bool {
	if t1 == BOOL_TYPE && t2 == NUMBER_TYPE {
		return true
	}
	if t2 == BOOL_TYPE && t1 == NUMBER_TYPE {
		return true
	}
	return false
}

func checkIsNumberAndString(t1 int, t2 int) bool {
	if t1 == STRING_TYPE && t2 == NUMBER_TYPE {
		return true
	}
	if t2 == STRING_TYPE && t1 == NUMBER_TYPE {
		return true
	}
	return false
}

func checkIsStringAndBytes(t1 int, t2 int) bool {
	if t1 == STRING_TYPE && t2 == BYTEARRAY_TYPE {
		return true
	}
	if t2 == STRING_TYPE && t1 == BYTEARRAY_TYPE {
		return true
	}
	return false
}

func checkIsNumber(t1 int, t2 int) bool {
	if t1 == NUMBER_TYPE && t2 == NUMBER_TYPE {
		return true
	}
	return false
}

func equalBoolAndNumber(src interface{}, dest interface{}) bool {
	var err error
	var fval1 float64
	var fval2 float64
	br, ok := src.(bool)
	if !ok {
		return false
	}
	if br {
		fval1 = 1
	} else {
		fval1 = 0
	}
	fval2, err = strconv.ParseFloat(fmt.Sprintf("%v", dest), 64)
	if err != nil {
		return false
	}
	return fval1 == fval2
}

func equalStringAndBytes(src interface{}, dest interface{}) bool {
	str, ok := src.(string)
	if !ok {
		return false
	}
	bs, ok := dest.([]byte)
	if !ok {
		return false
	}
	return str == string(bs)
}
func equalBoolAndString(src interface{}, dest interface{}) bool {
	br, ok := src.(bool)
	if !ok {
		return false
	}
	str, ok := dest.(string)
	if !ok {
		return false
	}
	if br {
		if str == "1" || strings.ToLower(str) == "true" {
			return true
		} else {
			return false
		}
	} else {
		if str == "0" || strings.ToLower(str) == "false" {
			return true
		} else {
			return false
		}
	}
}
func equalNumberAndString(src interface{}, dest interface{}) bool {
	var err error
	var fval1 float64
	var fval2 float64
	fval1, err = strconv.ParseFloat(fmt.Sprintf("%v", src), 64)
	if err != nil {
		return false
	}
	str, ok := dest.(string)
	if !ok {
		return false
	}
	fval2, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}

	if math.Abs(fval1-fval2) < numberic_min_threshold {
		return true
	}

	return false
}
func equalNumber(src interface{}, dest interface{}) bool {
	switch src.(type) {
	case float32:
		return equalFloatAndNumber(float64(src.(float32)), dest)
	case float64:
		return equalFloatAndNumber(src.(float64), dest)
	case int8:
		return equalIntAndNumber(int64(src.(int8)), dest)
	case int16:
		return equalIntAndNumber(int64(src.(int16)), dest)
	case int32:
		return equalIntAndNumber(int64(src.(int32)), dest)
	case int64:
		return equalIntAndNumber(int64(src.(int64)), dest)
	case int:
		return equalIntAndNumber(int64(src.(int)), dest)
	case uint8:
		return equalUintAndNumber(uint64(src.(uint8)), dest)
	case uint16:
		return equalUintAndNumber(uint64(src.(uint16)), dest)
	case uint32:
		return equalUintAndNumber(uint64(src.(uint32)), dest)
	case uint64:
		return equalUintAndNumber(uint64(src.(uint64)), dest)
	case uint:
		return equalUintAndNumber(uint64(src.(uint)), dest)
	default:
		return false
	}
}

func equalIntAndNumber(val int64, dest interface{}) bool {
	switch dest.(type) {
	case float32:
		return math.Abs(float64(val)-float64(dest.(float32))) < numberic_min_threshold
	case float64:
		return math.Abs(float64(val)-dest.(float64)) < numberic_min_threshold
	case int8:
		return val == int64(dest.(int8))
	case int16:
		return val == int64(dest.(int16))
	case int32:
		return val == int64(dest.(int32))
	case int64:
		return val == int64(dest.(int64))
	case int:
		return val == int64(dest.(int))
	case uint8:
		return val == int64(dest.(uint8))
	case uint16:
		return val == int64(dest.(uint16))
	case uint32:
		return val == int64(dest.(uint32))
	case uint64:
		return math.Abs(float64(val)-float64(dest.(uint64))) < numberic_min_threshold
	case uint:
		return math.Abs(float64(val)-float64(dest.(uint))) < numberic_min_threshold
	default:
		return false
	}
}

func equalUintAndNumber(val uint64, dest interface{}) bool {
	switch dest.(type) {
	case float32:
		return math.Abs(float64(val)-float64(dest.(float32))) < numberic_min_threshold
	case float64:
		return math.Abs(float64(val)-dest.(float64)) < numberic_min_threshold
	case int8:
		return math.Abs(float64(val)-float64(dest.(int8))) < numberic_min_threshold
	case int16:
		return math.Abs(float64(val)-float64(dest.(int16))) < numberic_min_threshold
	case int32:
		return math.Abs(float64(val)-float64(dest.(int32))) < numberic_min_threshold
	case int64:
		return math.Abs(float64(val)-float64(dest.(int64))) < numberic_min_threshold
	case int:
		return math.Abs(float64(val)-float64(dest.(int))) < numberic_min_threshold
	case uint8:
		return val == uint64(dest.(uint8))
	case uint16:
		return val == uint64(dest.(uint16))
	case uint32:
		return val == uint64(dest.(uint32))
	case uint64:
		return val == uint64(dest.(uint64))
	case uint:
		return val == uint64(dest.(uint))
	default:
		return false
	}
}
func equalFloatAndNumber(fval float64, dest interface{}) bool {
	switch dest.(type) {
	case float32:
		return math.Abs(fval-float64(dest.(float32))) < numberic_min_threshold
	case float64:
		return math.Abs(fval-dest.(float64)) < numberic_min_threshold
	case int8:
		return math.Abs(fval-float64(dest.(int8))) < numberic_min_threshold
	case int16:
		return math.Abs(fval-float64(dest.(int16))) < numberic_min_threshold
	case int32:
		return math.Abs(fval-float64(dest.(int32))) < numberic_min_threshold
	case int64:
		return math.Abs(fval-float64(dest.(int64))) < numberic_min_threshold
	case int:
		return math.Abs(fval-float64(dest.(int))) < numberic_min_threshold
	case uint8:
		return math.Abs(fval-float64(dest.(uint8))) < numberic_min_threshold
	case uint16:
		return math.Abs(fval-float64(dest.(uint16))) < numberic_min_threshold
	case uint32:
		return math.Abs(fval-float64(dest.(uint32))) < numberic_min_threshold
	case uint64:
		return math.Abs(fval-float64(dest.(uint64))) < numberic_min_threshold
	case uint:
		return math.Abs(fval-float64(dest.(uint))) < numberic_min_threshold
	default:
		return false
	}
}
