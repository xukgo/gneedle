package dyTypeUtil

import (
	"reflect"
	"strings"
)

//动态判断是否匹配，基础数据类型自动转换，
func Equal(src interface{}, dest interface{}) bool {
	srcKind := getKind(src)
	destKind := getKind(dest)

	//类型一样直接匹配，除非是浮点数，因为浮点判断是算最小差
	if srcKind == destKind && srcKind != reflect.Float32 && srcKind != reflect.Float64 {
		//if srcKind == reflect.Interface || srcKind == reflect.Ptr|| srcKind == reflect.Slice|| srcKind == reflect.Map{
		//	br1 := reflect.ValueOf(src).IsNil()
		//	br2 := reflect.ValueOf(dest).IsNil()
		//	if br1 == br2{
		//		return true
		//	}
		//}
		return src == dest
	}

	if srcKind == reflect.Invalid{
		rval := reflect.Zero(reflect.TypeOf(dest))
		ri := rval.Interface()
		kt := getValueKindType(dest)
		if kt >= UNSUPPORT_TYPE && (destKind != reflect.Interface && destKind != reflect.Ptr&& destKind != reflect.Slice&& destKind != reflect.Map ){
			return false
		}
		return Equal(ri,dest)
	}
	if destKind == reflect.Invalid{
		rval := reflect.Zero(reflect.TypeOf(src))
		ri := rval.Interface()
		kt := getValueKindType(src)
		if kt >= UNSUPPORT_TYPE && (srcKind != reflect.Interface && srcKind != reflect.Ptr&& srcKind != reflect.Slice&& srcKind != reflect.Map ){
			return false
		}
		return Equal(src,ri)
	}

	srcKindType := getValueKindType(src)
	if srcKindType >= UNSUPPORT_TYPE {
		return false
	}
	destKindType := getValueKindType(dest)
	if destKindType >= UNSUPPORT_TYPE {
		return false
	}

	if checkIsBoolAndString(srcKindType, destKindType) {
		if srcKindType == BOOL_TYPE {
			return equalBoolAndString(src, dest)
		} else {
			return equalBoolAndString(dest, src)
		}
	}
	if checkIsStringAndBytes(srcKindType, destKindType) {
		if srcKindType == STRING_TYPE {
			return equalStringAndBytes(src, dest)
		} else {
			return equalStringAndBytes(dest, src)
		}
	}

	if checkIsBoolAndNumber(srcKindType, destKindType) {
		if srcKindType == BOOL_TYPE {
			return equalBoolAndNumber(src, dest)
		} else {
			return equalBoolAndNumber(dest, src)
		}
	}

	if checkIsNumberAndString(srcKindType, destKindType) {
		if srcKindType == NUMBER_TYPE {
			str := dest.(string)
			if strings.EqualFold(str,"true"){
				return equalNumber(src, 1)
			}else if strings.EqualFold(str,"false"){
				return equalNumber(src, 0)
			}else{
				return equalNumberAndString(src, dest)
			}
		} else {
			str := src.(string)
			if strings.EqualFold(str,"true"){
				return equalNumber(1,dest)
			}else if strings.EqualFold(str,"false"){
				return equalNumber(0,dest)
			}else{
				return equalNumberAndString(dest, src)
			}
		}
	}

	if checkIsNumber(srcKindType, destKindType) {
		return equalNumber(src, dest)
	}

	return false
}

//var numberLevelMap map[reflect.Kind]int = make(map[reflect.Kind]int)
//var levelTypeMap map[int]reflect.Kind = make(map[int]reflect.Kind)
//
//func init() {
//	sysbit := getSystemBit()
//	numberLevelMap[reflect.Int8] = 2
//	numberLevelMap[reflect.Uint8] = 2
//	numberLevelMap[reflect.Int16] = 3
//	numberLevelMap[reflect.Uint16] = 3
//	numberLevelMap[reflect.Int32] = 4
//	numberLevelMap[reflect.Uint32] = 4
//	numberLevelMap[reflect.Int64] = 5
//	numberLevelMap[reflect.Uint64] = 5
//	numberLevelMap[reflect.Float32] = 6
//	numberLevelMap[reflect.Float64] = 7
//
//	if sysbit == 64 {
//		numberLevelMap[reflect.Int] = numberLevelMap[reflect.Int64]
//		numberLevelMap[reflect.Uint] = numberLevelMap[reflect.Int64]
//	} else {
//		numberLevelMap[reflect.Int] = numberLevelMap[reflect.Int32]
//		numberLevelMap[reflect.Uint] = numberLevelMap[reflect.Int32]
//	}
//
//	levelTypeMap[2] = reflect.Int8
//	levelTypeMap[3] = reflect.Int16
//	levelTypeMap[4] = reflect.Int32
//	levelTypeMap[5] = reflect.Int64
//	levelTypeMap[6] = reflect.Float32
//	levelTypeMap[7] = reflect.Float64
//}
