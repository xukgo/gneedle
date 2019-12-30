/*
@Time : 2019/9/30 14:18
@Author : Hermes
@File : core
@Description:
*/
package voluator

import (
	"fmt"
	"reflect"
	"strings"
)

func AssignValue(instance interface{}, path string, val interface{}, cached bool) (bool, error) {
	path = strings.TrimSpace(path)
	paths := strings.Split(path, ".")
	if cached {
		return assignByPathsCached(instance, paths, val)
	}
	return assignByPaths(instance, paths, val)
}

//instance必须是指针类型的变量，比如指针 接口 slice等才能走下去
func assignByPaths(instance interface{}, paths []string, val interface{}) (bool, error) {
	if instance == nil || paths == nil {
		return false, nil
	}

	elem := reflect.ValueOf(instance)
	elemKind := elem.Kind()
	if elemKind != reflect.Ptr && elemKind != reflect.Interface {
		return false, fmt.Errorf("not support assign type:%s", elemKind)
	}

	elem = elem.Elem()
	elemType := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		fieldInfo := elemType.Field(i)
		if fieldInfo.Anonymous {
			acField := elem.Field(i)
			ptrElem := getValuePtr(acField)
			br, err := assignByPaths(ptrElem, paths, val)
			if err != nil {
				return false, err
			}
			if !br {
				continue
			} else {
				return true, nil
			}
		}

		flagName := getFieldFlagName(fieldInfo)
		if strings.ToLower(flagName) == strings.ToLower(paths[0]) {
			acField := elem.Field(i)
			ptrElem := getValuePtr(acField)
			if len(paths) == 1 {
				err := assignFieldValue(acField, &fieldInfo, val)
				if err != nil {
					return true, err
				}
				return true, nil
			} else {
				return assignByPaths(ptrElem, paths[1:], val)
			}
		}
	}

	return false, nil
}
