package voluator

import (
	"reflect"
)

func getByPathsCached(v interface{}, paths []string) (bool, interface{}) {
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

	//instanceClassPath := GetClassPath(v)
	//instancePathMapObj, find := pathCacheDict.Load(instanceClassPath)
	//if !find {
	//	initPathCache(v)
	//	instancePathMapObj, find = pathCacheDict.Load(instanceClassPath)
	//	if !find {
	//		return false, nil
	//	}
	//}

	flag, sliceIdx := getFlagAndSliceIndex(paths[0])
	instancePathMap := lockLoadPathMap(v)
	//instancePathMap := instancePathMapObj.(map[string]MinorPath)
	mpath, ok := instancePathMap[flag]

	var fins interface{}
	if sliceIdx < 0 {
		if !ok {
			mpath, ok = instancePathMap[ANONYMOUS_FIELD_KEY]
			if !ok {
				return false, nil
			}
			return getByPathsCached(elem.Field(mpath.Index).Interface(), paths)
		} else {
			fins = elem.Field(mpath.Index).Interface()
		}
	} else {
		if !ok {
			return false, nil
		}
		subFiled := elem.Field(mpath.Index)
		subFieldKind := subFiled.Kind()
		if subFieldKind != reflect.Slice {
			return false, nil
		}
		if sliceIdx >= subFiled.Len() {
			return false, nil
		}
		fins = subFiled.Index(sliceIdx).Interface()
	}

	if len(paths) == 1 {
		return true, fins
	} else {
		return getByPathsCached(fins, paths[1:])
	}
}
