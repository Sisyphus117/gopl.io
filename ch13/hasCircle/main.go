package hascircle

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	t reflect.Type
	p unsafe.Pointer
}

func HasCircle(v reflect.Value) bool {
	mp := make(map[comparison]struct{})
	return search(v, mp)
}

func search(v reflect.Value, mp map[comparison]struct{}) bool {
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		c := comparison{v.Type(), vptr}
		if _, has := mp[c]; has {
			return true
		}
		mp[c] = struct{}{}
	}
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return false
		}
		return search(v.Elem(), mp)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			res := search(v.Index(i), mp)
			if res {
				return true
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			res := search(v.MapIndex(key), mp)
			if res {
				return true
			}
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			res := search(v.Field(i), mp)
			if res {
				return true
			}
		}
	}
	return false
}
