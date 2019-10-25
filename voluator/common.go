/*
@Time : 2019/9/30 19:15
@Author : Hermes
@File : common
@Description:
*/
package voluator

import (
	"reflect"
	"strings"
)

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
