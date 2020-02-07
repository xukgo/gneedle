/*
@Time : 2019/9/30 19:15
@Author : Hermes
@File : common
@Description:
*/
package voluator

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func combindPathString(pkg, name string)string{
	len1 := len(pkg)
	len2 := len(name)
	buff := make([]byte,len1+len2+1)
	copy(buff, pkg)
	buff[len1] = '.'
	copy(buff[len1+1:], name)
	return string(buff)
}

func GetClassPath(v interface{}) string {
	tp := reflect.TypeOf(v)

	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	//return tp.String() //这个是包名+结构体名，比较短
	return combindPathString(tp.PkgPath(), tp.Name()) //这个是全部包名路径，有点长
}

func getFieldFlagName(fieldInfo reflect.StructField) string {
	jsonTag := fieldInfo.Tag.Get("json")
	var itemName string
	if jsonTag != "" && jsonTag != "-" {
		arr := strings.Split(jsonTag, ",")
		itemName = arr[0]
	}
	if itemName == "" {
		itemName = fieldInfo.Name
	}
	return strings.TrimSpace(itemName)
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
		log.Printf("getFlagAndSliceIndex parse url format error,%s\r\n",str)
		return str[:idx1], -1
	}
	return str[:idx1], sliceIdx
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
