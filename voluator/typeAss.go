/*
@Time : 2019/9/30 16:35
@Author : Hermes
@File : typeAss
@Description:
*/
package voluator

import (
	"fmt"
	"reflect"
	"strconv"
)

func convertAssign(valType reflect.Type, value reflect.Value, targetVal interface{}) error {
	switch valType.Kind() {
	//字符串允许int float 二进制流转换
	case reflect.String:
		switch tval := targetVal.(type) {
		case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			value.SetString(fmt.Sprintf("%v", tval))
			return nil
		case []byte:
			value.SetString(string(targetVal.([]byte)))
			return nil
		default:
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}

	//int类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch tval := targetVal.(type) {
		case int:
			value.SetInt(int64(tval))
			return nil
		case int8:
			value.SetInt(int64(tval))
			return nil
		case int16:
			value.SetInt(int64(tval))
			return nil
		case int32:
			value.SetInt(int64(tval))
			return nil
		case int64:
			value.SetInt(int64(tval))
			return nil
		case uint:
			value.SetInt(int64(tval))
			return nil
		case uint8:
			value.SetInt(int64(tval))
			return nil
		case uint16:
			value.SetInt(int64(tval))
			return nil
		case uint32:
			value.SetInt(int64(tval))
			return nil
		case uint64:
			value.SetInt(int64(tval))
			return nil
		case float32:
			value.SetInt(int64(tval))
			return nil
		case float64:
			value.SetInt(int64(tval))
			return nil
		case string:
			mval, err := strconv.ParseInt(tval, 10, 64)
			if err != nil {
				return err
			}
			value.SetInt(mval)
			return nil
		default:
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}

	//uint类型
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch tval := targetVal.(type) {
		case int:
			value.SetUint(uint64(tval))
			return nil
		case int8:
			value.SetUint(uint64(tval))
			return nil
		case int16:
			value.SetUint(uint64(tval))
			return nil
		case int32:
			value.SetUint(uint64(tval))
			return nil
		case int64:
			value.SetUint(uint64(tval))
			return nil
		case uint:
			value.SetUint(uint64(tval))
			return nil
		case uint8:
			value.SetUint(uint64(tval))
			return nil
		case uint16:
			value.SetUint(uint64(tval))
			return nil
		case uint32:
			value.SetUint(uint64(tval))
			return nil
		case uint64:
			value.SetUint(uint64(tval))
			return nil
		case float32:
			value.SetUint(uint64(tval))
			return nil
		case float64:
			value.SetUint(uint64(tval))
			return nil
		case string:
			mval, err := strconv.ParseUint(tval, 10, 64)
			if err != nil {
				return err
			}
			value.SetUint(mval)
			return nil
		default:
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}

	case reflect.Float32, reflect.Float64:
		switch tval := targetVal.(type) {
		case int:
			value.SetFloat(float64(tval))
			return nil
		case int8:
			value.SetFloat(float64(tval))
			return nil
		case int16:
			value.SetFloat(float64(tval))
			return nil
		case int32:
			value.SetFloat(float64(tval))
			return nil
		case int64:
			value.SetFloat(float64(tval))
			return nil
		case uint:
			value.SetFloat(float64(tval))
			return nil
		case uint8:
			value.SetFloat(float64(tval))
			return nil
		case uint16:
			value.SetFloat(float64(tval))
			return nil
		case uint32:
			value.SetFloat(float64(tval))
			return nil
		case uint64:
			value.SetFloat(float64(tval))
			return nil
		case float32:
			value.SetFloat(float64(tval))
			return nil
		case float64:
			value.SetFloat(float64(tval))
			return nil
		case string:
			mval, err := strconv.ParseFloat(tval, 64)
			if err != nil {
				return err
			}
			value.SetFloat(mval)
			return nil
		default:
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}
	}
	return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
}
