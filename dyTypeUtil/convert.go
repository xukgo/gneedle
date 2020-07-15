package dyTypeUtil

import (
	"fmt"
	"math"
	"strconv"
)

var DEFAULT_EMPTY_BYTESLICE = make([]byte,0)

func ConvertToString(src interface{})(string,error){
	v,err := ConvertBasicTypeData(src,"")
	if err != nil{
		return "",err
	}
	return v.(string),nil
}
//func ConvertToByteArray(src interface{})([]byte,error){
//	v,err := ConvertBasicTypeData(src, DEFAULT_EMPTY_BYTESLICE)
//	if err != nil{
//		return nil,err
//	}
//	return v.([]byte),nil
//}

func ConvertToBool(src interface{})(bool,error){
	v,err := ConvertBasicTypeData(src,true)
	if err != nil{
		return false,err
	}
	return v.(bool),nil
}

func ConvertToInt(src interface{})(int,error){
	v,err := ConvertBasicTypeData(src,int(0))
	if err != nil{
		return 0,err
	}
	return v.(int),nil
}

func ConvertToInt64(src interface{})(int64,error){
	v,err := ConvertBasicTypeData(src,int64(0))
	if err != nil{
		return 0,err
	}
	return v.(int64),nil
}

func ConvertToUint64(src interface{})(uint64,error){
	v,err := ConvertBasicTypeData(src,uint64(0))
	if err != nil{
		return 0,err
	}
	return v.(uint64),nil
}
func ConvertToFloat32(src interface{})(float32,error){
	v,err := ConvertBasicTypeData(src,float32(0))
	if err != nil{
		return 0,err
	}
	return v.(float32),nil
}
func ConvertToFloat64(src interface{})(float64,error){
	v,err := ConvertBasicTypeData(src,float64(0))
	if err != nil{
		return 0,err
	}
	return v.(float64),nil
}

//支持基础数据类型自动转换，不支持 byteArray，只支持数值和string/bool
func ConvertBasicTypeData(src interface{}, dest interface{}) (interface{}, error) {
	srcKind := getKind(src)
	destKind := getKind(dest)
	if srcKind == destKind {
		return src, nil
	}

	srcKindType := getValueKindType(src)
	if srcKindType >= UNSUPPORT_TYPE{
		return nil, fmt.Errorf("unsupport type")
	}
	destKindType := getValueKindType(dest)
	if destKindType >= UNSUPPORT_TYPE{
		return nil, fmt.Errorf("unsupport type")
	}

	//byteArray只允许和string转换
	if srcKindType == BYTEARRAY_TYPE && destKindType != STRING_TYPE{
		return nil, fmt.Errorf("unsupport type")
	}
	if destKindType == BYTEARRAY_TYPE && srcKindType != STRING_TYPE{
		return nil, fmt.Errorf("unsupport type")
	}

	if checkIsStringAndBytes(srcKindType, destKindType) {
		if srcKindType == BYTEARRAY_TYPE {
			return converBytesToString(src.([]byte)), nil
		} else {
			return convertStringToBytes(src.(string)), nil
		}
	}

	if checkIsBoolAndString(srcKindType, destKindType) {
		if srcKindType == BOOL_TYPE {
			return convertBoolToString(src.(bool)), nil
		} else {
			return convertStringToBool(src.(string)), nil
		}
	}

	if checkIsBoolAndNumber(srcKindType, destKindType) {
		if srcKindType == BOOL_TYPE {
			return convertBoolToNumber(src.(bool), dest)
		} else {
			return convertNumberToBool(src), nil
		}
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

func convertStringToBytes(s string) []byte {
	return []byte(s)
}

func converBytesToString(bytes []byte) string {
	if len(bytes) == 0{
		return ""
	}
	return string(bytes)
}

func convertStringToBool(str string) bool {
	slen := len(str)
	if slen == 1{
		if str[0] == 't' || str[0] == 'T'{
			return true
		}
		return false
	}else if slen == 4{
		if str[0] != 't' && str[0] != 'T'{
			return false
		}
		if str[1] != 'r' && str[1] != 'R'{
			return false
		}
		if str[2] != 'u' && str[2] != 'U'{
			return false
		}
		if str[3] != 'e' && str[3] != 'E'{
			return false
		}
		return true
	}else {
		return false
	}
}

func convertBoolToString(br bool) string {
	if br {
		return "true"
	}
	return "false"
}

func convertStringToNumber(src string, dest interface{}) (interface{}, error) {
	switch dest.(type) {
	case float32:
		val, err :=  strconv.ParseFloat(src, 32)
		if err != nil {
			return nil, err
		}
		return float32(val), nil
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

func convertBoolToNumber(src bool, dest interface{}) (interface{}, error) {
	switch dest.(type) {
	case float32:
		if src {
			return float32(1), nil
		} else {
			return float32(0), nil
		}
	case float64:
		if src {
			return float64(1), nil
		} else {
			return float64(0), nil
		}
	case int8:
		if src {
			return int8(1), nil
		} else {
			return int8(0), nil
		}
	case int16:
		if src {
			return int16(1), nil
		} else {
			return int16(0), nil
		}
	case int32:
		if src {
			return int32(1), nil
		} else {
			return int32(0), nil
		}
	case int64:
		if src {
			return int64(1), nil
		} else {
			return int64(0), nil
		}
	case int:
		if src {
			return int(1), nil
		} else {
			return int(0), nil
		}
	case uint8:
		if src {
			return uint8(1), nil
		} else {
			return uint8(0), nil
		}
	case uint16:
		if src {
			return uint16(1), nil
		} else {
			return uint16(0), nil
		}
	case uint32:
		if src {
			return uint32(1), nil
		} else {
			return uint32(0), nil
		}
	case uint64:
		if src {
			return uint64(1), nil
		} else {
			return uint64(0), nil
		}
	case uint:
		if src {
			return uint(1), nil
		} else {
			return uint(0), nil
		}
	default:
		return nil, fmt.Errorf("unsupport type")
	}
}

func convertNumberToBool(src interface{}) bool {
	switch src.(type) {
	case float32:
		if math.Abs(float64(src.(float32))) < numberic_min_threshold {
			return false
		}
		return true
	case float64:
		if math.Abs(src.(float64)) < numberic_min_threshold {
			return false
		}
		return true
	case int8:
		if src.(int8) == 0 {
			return false
		}
		return true
	case int16:
		if src.(int16) == 0 {
			return false
		}
		return true
	case int32:
		if src.(int32) == 0 {
			return false
		}
		return true
	case int64:
		if src.(int64) == 0 {
			return false
		}
		return true
	case int:
		if src.(int) == 0 {
			return false
		}
		return true
	case uint8:
		if src.(uint8) == 0 {
			return false
		}
		return true
	case uint16:
		if src.(uint16) == 0 {
			return false
		}
		return true
	case uint32:
		if src.(uint32) == 0 {
			return false
		}
		return true
	case uint64:
		if src.(uint64) == 0 {
			return false
		}
		return true
	case uint:
		if src.(uint) == 0 {
			return false
		}
		return true
	default:
		return false
	}
}
