/*
@Time : 2019/9/30 16:35
@Author : Hermes
@File : typeAss
@Description:
*/
package voluator

import (
	"fmt"
	"github.com/xukgo/gneedle/dyTypeUtil"
	"reflect"
)

func convertAssign(valType reflect.Type, value reflect.Value, targetVal interface{}) error {
	switch valType.Kind() {
	//字符串允许int float 二进制流转换
	case reflect.String:
		tval,err := dyTypeUtil.ConvertToString(targetVal)
		if err != nil{
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}
		value.SetString(tval)
		return nil
	//int类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tval,err := dyTypeUtil.ConvertToInt64(targetVal)
		if err != nil{
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}
		value.SetInt(tval)
		return nil
	//uint类型
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tval,err := dyTypeUtil.ConvertToUint64(targetVal)
		if err != nil{
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}
		value.SetUint(tval)
		return nil
	case reflect.Float32, reflect.Float64:
		tval,err := dyTypeUtil.ConvertToFloat64(targetVal)
		if err != nil{
			return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
		}
		value.SetFloat(tval)
		return nil
	}
	return fmt.Errorf("source value type can not convert to target value type:%s->%s", reflect.ValueOf(targetVal).Type(), valType)
}
