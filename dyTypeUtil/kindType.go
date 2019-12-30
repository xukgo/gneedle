package dyTypeUtil

import (
	"reflect"
)

const UNSUPPORT_TYPE = -1
const BOOL_TYPE = 1
const NUMBER_TYPE = 2
const STRING_TYPE = 3
const BYTEARRAY_TYPE = 4

func getValueKindType(v interface{}) int {
	switch v.(type) {
	case bool:
		return BOOL_TYPE
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return NUMBER_TYPE
	case string:
		return STRING_TYPE
	case []byte:
		return BYTEARRAY_TYPE
	default:
		return UNSUPPORT_TYPE
	}
}

func getKind(v interface{}) reflect.Kind {
	elem := reflect.ValueOf(v)
	//fmt.Println(elem)
	elemType := elem.Type()
	//fmt.Println(elemType)
	k := elemType.Kind()
	//fmt.Println(k)
	return k
}

//获取系统的位数 32或64
func getSystemBit() int {
	bit := 32 << (^uint(0) >> 63)
	return bit
}
