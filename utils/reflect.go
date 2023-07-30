package utils

import "reflect"

//判断interface类型,反射

func ReflectType(payload interface{}, p reflect.Kind) bool {
	pay := reflect.TypeOf(payload)
	kind := pay.Kind()
	if kind == p {
		return true
	}
	return false
}
