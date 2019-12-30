package voluator

import (
	"reflect"
	"sync"
)

const ANONYMOUS_FIELD_KEY = "_@Ams"

var pathCacheDict sync.Map
var pathCacheDictLocker sync.Mutex

type MinorPath struct {
	Index int
}

func NewMinorPath(index int) MinorPath {
	model := MinorPath{
		Index: index,
	}
	return model
}

func initPathCache(v interface{}) {
	if v == nil {
		return
	}

	elem := reflect.ValueOf(v)
	elemKind := elem.Kind()
	if elemKind != reflect.Struct && elemKind != reflect.Ptr {
		return
	}
	if elemKind == reflect.Ptr {
		elem = elem.Elem()
	}
	if elem.Kind() != reflect.Struct {
		return
	}

	instanceClassPath := GetClassPath(v)
	_, find := pathCacheDict.Load(instanceClassPath)
	if find {
		return
	}

	elemType := elem.Type()

	dict := make(map[string]MinorPath)
	for i := 0; i < elem.NumField(); i++ {
		fieldInfo := elemType.Field(i)
		if fieldInfo.Anonymous {
			initPathCache(elem.Field(i).Interface())
			minorPath := NewMinorPath(i)
			dict[ANONYMOUS_FIELD_KEY] = minorPath
			continue
		}

		var fins interface{} = nil
		subFiled := elem.Field(i)
		//slice处理
		if fieldInfo.Type.Kind() == reflect.Slice {
			SubCount := subFiled.Len()
			if SubCount > 0 {
				fins = subFiled.Index(0).Interface()
			}
		} else {
			fins = subFiled.Interface()
		}
		if fins != nil {
			initPathCache(fins)
		}
		flagName := getFieldFlagName(fieldInfo)
		minorPath := NewMinorPath(i)
		dict[flagName] = minorPath
	}

	pathCacheDict.Store(instanceClassPath, dict)
	//fmt.Println("struct name:", key)
	//fmt.Println("struct children:", dict)
}

func lockLoadPathMap(instance interface{}) map[string]MinorPath {
	instanceClassPath := GetClassPath(instance)
	obj, find := pathCacheDict.Load(instanceClassPath)
	if find {
		return obj.(map[string]MinorPath)
	}

	pathCacheDictLocker.Lock()
	initPathCache(instance)
	pathCacheDictLocker.Unlock()

	obj, _ = pathCacheDict.Load(instanceClassPath)
	return obj.(map[string]MinorPath)
}
