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
)

//func AssignValueCached(instance interface{}, path string, val interface{}) (bool, error) {
//	path = strings.TrimSpace(path)
//	paths := strings.Split(path, ".")
//	return assignByPathsCached(instance, paths, val)
//}

//instance必须是指针类型的变量，比如指针 接口 slice等才能走下去
func assignByPathsCached(instance interface{}, paths []string, val interface{}) (bool, error) {
	if instance == nil || paths == nil {
		return false, nil
	}

	elem := reflect.ValueOf(instance)
	elemKind := elem.Kind()
	if elemKind != reflect.Ptr && elemKind != reflect.Interface {
		return false, fmt.Errorf("not support assign type:%s", elemKind)
	}
	elem = elem.Elem()

	//instanceClassPath := GetClassPath(instance)
	//instancePathMapObj, find := pathCacheDict.Load(instanceClassPath)
	//if !find {
	//	initPathCache(instance)
	//	instancePathMapObj, find = pathCacheDict.Load(instanceClassPath)
	//	if !find {
	//		return false, fmt.Errorf("initPathCache then cannot find class load info")
	//	}
	//}

	instancePathMap := lockLoadPathMap(instance)
	flag, _ := getFlagAndSliceIndex(paths[0])
	mpath, ok := instancePathMap[flag]

	if !ok {
		mpath, ok = instancePathMap[ANONYMOUS_FIELD_KEY]
		if !ok {
			return false, nil
		}
		ptrElem := getValuePtr(elem.Field(mpath.Index))
		return assignByPaths(ptrElem, paths, val)
	}

	if len(paths) > 1 {
		ptrElem := getValuePtr(elem.Field(mpath.Index))
		return assignByPathsCached(ptrElem, paths[1:], val)
	}

	elemType := elem.Type()
	fieldInfo := elemType.Field(mpath.Index)
	acField := elem.Field(mpath.Index)
	err := assignFieldValue(acField, &fieldInfo, val)
	if err != nil {
		return true, err
	}
	return true, nil
}
