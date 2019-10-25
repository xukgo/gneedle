package dyEqual

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

	if math.Abs(fval1-fval2) < 0.01 {
		return true
	}

	return false
}
func equalNumber(src interface{}, dest interface{}) bool {
	var err error
	var fval1 float64
	var fval2 float64
	fval1, err = strconv.ParseFloat(fmt.Sprintf("%v", src), 64)
	if err != nil {
		return false
	}
	fval2, err = strconv.ParseFloat(fmt.Sprintf("%v", dest), 64)
	if err != nil {
		return false
	}
	if math.Abs(fval1-fval2) < 0.01 {
		return true
	}

	return false
}
