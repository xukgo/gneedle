package dyTypeUtil

import (
	"fmt"
	"strconv"
)

//支持基础数据类型自动转换，不支持bool byteArray，只支持数值和string
func ConvertBasicTypeData(src interface{}, dest interface{}) (interface{}, error) {
	srcKind := getKind(src)
	destKind := getKind(dest)
	if srcKind == destKind {
		return src, nil
	}

	srcKindType := getValueKindType(src)
	if srcKindType == UNSUPPORT_TYPE || srcKindType == BOOL_TYPE || srcKindType == BYTEARRAY_TYPE {
		return nil, fmt.Errorf("unsupport type")
	}
	destKindType := getValueKindType(dest)
	if destKindType == UNSUPPORT_TYPE || destKindType == BOOL_TYPE || destKindType == BYTEARRAY_TYPE {
		return nil, fmt.Errorf("unsupport type")
	}

	if checkIsNumberAndString(srcKindType, destKindType) {
		if srcKindType == NUMBER_TYPE {
			return fmt.Sprintf("%v", src), nil
		} else {
			return convertStringToNumber(src.(string), dest)
		}
	}

	if checkIsNumber(srcKindType, destKindType) {
		return convertNumberToNumber(src, dest)
	}

	return nil, fmt.Errorf("unknow error")
}

func convertStringToNumber(src string, dest interface{}) (interface{}, error) {
	switch dest.(type) {
	case float32:
		return strconv.ParseFloat(src, 32)
	case float64:
		return strconv.ParseFloat(src, 64)
	case int8:
		val, err := strconv.ParseInt(src, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8(val), nil
	case int16:
		val, err := strconv.ParseInt(src, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16(val), nil
	case int32:
		val, err := strconv.ParseInt(src, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(val), nil
	case int64:
		val, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return nil, err
		}
		return val, nil
	case int:
		val, err := strconv.ParseInt(src, 10, 32)
		if err != nil {
			return nil, err
		}
		return int(val), nil
	case uint8:
		val, err := strconv.ParseUint(src, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8(val), nil
	case uint16:
		val, err := strconv.ParseUint(src, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16(val), nil
	case uint32:
		val, err := strconv.ParseUint(src, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(val), nil
	case uint64:
		val, err := strconv.ParseUint(src, 10, 64)
		if err != nil {
			return nil, err
		}
		return val, nil
	case uint:
		val, err := strconv.ParseUint(src, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint(val), nil
	default:
		return nil, fmt.Errorf("unsupport type")
	}
}

func convertNumberToNumber(src interface{}, dest interface{}) (interface{}, error) {
	switch src.(type) {
	case float32:
		return convertFloatToNumber(float64(src.(float32)), dest)
	case float64:
		return convertFloatToNumber(src.(float64), dest)
	case int8:
		return convertIntToNumber(int64(src.(int8)), dest)
	case int16:
		return convertIntToNumber(int64(src.(int16)), dest)
	case int32:
		return convertIntToNumber(int64(src.(int32)), dest)
	case int64:
		return convertIntToNumber(int64(src.(int64)), dest)
	case int:
		return convertIntToNumber(int64(src.(int)), dest)
	case uint8:
		return convertUintToNumber(uint64(src.(uint8)), dest)
	case uint16:
		return convertUintToNumber(uint64(src.(uint16)), dest)
	case uint32:
		return convertUintToNumber(uint64(src.(uint32)), dest)
	case uint64:
		return convertUintToNumber(uint64(src.(uint64)), dest)
	case uint:
		return convertUintToNumber(uint64(src.(uint)), dest)
	default:
		return nil, fmt.Errorf("unsupport type")
	}
}

func convertUintToNumber(val uint64, dest interface{}) (interface{}, error) {
	switch dest.(type) {
	case float32:
		return float32(val), nil
	case float64:
		return float64(val), nil
	case int8:
		return int8(val), nil
	case int16:
		return int16(val), nil
	case int32:
		return int32(val), nil
	case int64:
		return int64(val), nil
	case int:
		return int(val), nil
	case uint8:
		return uint8(val), nil
	case uint16:
		return uint16(val), nil
	case uint32:
		return uint32(val), nil
	case uint64:
		return uint64(val), nil
	case uint:
		return uint(val), nil
	default:
		return nil, fmt.Errorf("unsupport type")
	}
}

func convertIntToNumber(val int64, dest interface{}) (interface{}, error) {
	switch dest.(type) {
	case float32:
		return float32(val), nil
	case float64:
		return float64(val), nil
	case int8:
		return int8(val), nil
	case int16:
		return int16(val), nil
	case int32:
		return int32(val), nil
	case int64:
		return int64(val), nil
	case int:
		return int(val), nil
	case uint8:
		return uint8(val), nil
	case uint16:
		return uint16(val), nil
	case uint32:
		return uint32(val), nil
	case uint64:
		return uint64(val), nil
	case uint:
		return uint(val), nil
	default:
		return nil, fmt.Errorf("unsupport type")
	}
}

func convertFloatToNumber(fval float64, dest interface{}) (interface{}, error) {
	switch dest.(type) {
	case float32:
		return float32(fval), nil
	case float64:
		return fval, nil
	case int8:
		return int8(fval), nil
	case int16:
		return int16(fval), nil
	case int32:
		return int32(fval), nil
	case int64:
		return int64(fval), nil
	case int:
		return int(fval), nil
	case uint8:
		return uint8(fval), nil
	case uint16:
		return uint16(fval), nil
	case uint32:
		return uint32(fval), nil
	case uint64:
		return uint64(fval), nil
	case uint:
		return uint(fval), nil
	default:
		return nil, fmt.Errorf("unsupport type")
	}
}
