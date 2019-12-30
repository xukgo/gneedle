/*
@Time : 2019/9/30 19:02
@Author : Hermes
@File : getValueCore
@Description:
*/
package voluator

import (
	"reflect"
	"strings"
)

//v允许是结构体或者指针
func GetValue(v interface{}, path string, cached bool) (bool, interface{}) {
	path = strings.TrimSpace(path)
	paths := strings.Split(path, ".")

	if cached {
		return getByPathsCached(v, paths)
	}
	return getByPaths(v, paths)
}

func getByPaths(v interface{}, paths []string) (bool, interface{}) {
	if v == nil || paths == nil {
		return false, nil
	}

	elem := reflect.ValueOf(v)
	elemKind := elem.Kind()
	if elemKind != reflect.Struct && elemKind != reflect.Ptr {
		return false, nil
	}
	if elemKind == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		return false, nil
	}

	elemType := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		fieldInfo := elemType.Field(i)
		if fieldInfo.Anonymous {
			br, rv := getByPaths(elem.Field(i).Interface(), paths)
			if !br {
				continue
			} else {
				return true, rv
			}
		}

		flagName := getFieldFlagName(fieldInfo)
		flag, sliceIdx := getFlagAndSliceIndex(paths[0])
		if strings.ToLower(flagName) == strings.ToLower(flag) {
			subFiled := elem.Field(i)
			var fins interface{}
			if sliceIdx >= 0 { //处理slice的
				//处理slice类型的
				subFieldKind := subFiled.Kind()
				if subFieldKind != reflect.Slice {
					return false, nil
				}
				if sliceIdx >= subFiled.Len() {
					return false, nil
				}
				fins = subFiled.Index(sliceIdx).Interface()
			} else {
				fins = subFiled.Interface()
			}

			if len(paths) == 1 {
				return true, fins
			} else {
				return getByPaths(fins, paths[1:])
			}
		}
	}

	return false, nil
}
