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

func AssignValue(instance interface{}, path string, val interface{}) (bool, error) {
	path = strings.TrimSpace(path)
	paths := strings.Split(path, ".")
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

func getValuePtr(field reflect.Value) interface{} {
	kind := field.Kind()
	if kind != reflect.Ptr && kind != reflect.Interface {
		field = field.Addr()
	}
	return field.Interface()
}

//结构相同则直接赋值，不同则尝试转换赋值，转换仅限于基础数据类型的int string
func assignFieldValue(value reflect.Value, field *reflect.StructField, setVal interface{}) error {
	if !value.CanSet() {
		return fmt.Errorf("%s can not set value", field.Name)
	}
	//fmt.Println("offset", field.Offset)
	//fmt.Println("PkgPath", field.PkgPath)
	targetValue := reflect.ValueOf(setVal)
	targetValType := targetValue.Type()

	//fmt.Println(targetValType)
	//fmt.Println(field.Type)
	if targetValType == field.Type {
		value.Set(targetValue)
		return nil
	} else {
		err := convertAssign(field.Type, value, setVal)
		if err != nil {
			return err
		}
		return nil
	}
}
