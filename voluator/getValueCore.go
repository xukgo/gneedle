/*
@Time : 2019/9/30 19:02
@Author : Hermes
@File : getValueCore
@Description:
*/
package voluator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//get的路径允许a.b.c和数组a.b[2]这样的，这里只判断语法
func CheckGetPathValid(path string) bool {
	path = strings.TrimSpace(path)
	if path == "" {
		return false
	}

	arr := strings.Split(path, ".")
	for _, item := range arr {
		if !checkIsTagNodeFormat(item) {
			return false
		}
	}
	return true
}

//格式是数字英文下划线和允许以[0]这样结尾
func checkIsTagNodeFormat(str string) bool {
	matched, _ := regexp.MatchString(`^([0-9a-zA-Z_]{1,})(\[[0-9]+\])?$`, str)
	return matched
}

//v允许是结构体或者指针
func GetValue(v interface{}, path string) (bool, interface{}) {
	path = strings.TrimSpace(path)
	paths := strings.Split(path, ".")
	return getByPaths(v, paths)
}

func getByPaths(v interface{}, paths []string) (bool, interface{}) {
	if v == nil || paths == nil {
		return false, nil
	}

	elem := reflect.ValueOf(v)
	elemKind := elem.Kind()
	if elemKind != reflect.Struct {
		elem = elem.Elem()
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

func getFlagAndSliceIndex(str string) (string, int) {
	idx1 := strings.Index(str, "[")
	if idx1 < 0 {
		return str, -1
	}
	idx2 := strings.Index(str, "]")

	sliceIdxStr := str[idx1+1 : idx2]
	sliceIdx, err := strconv.Atoi(sliceIdxStr)
	if err != nil {
		return str[:idx1], -1
	}
	return str[:idx1], sliceIdx
}
