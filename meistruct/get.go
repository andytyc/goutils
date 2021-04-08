package meistruct

import "reflect"

// GetName 获取对象的struct的名称{*Struct}
func GetName(result interface{}) string {
	return reflect.TypeOf(result).Name()
}

// GetNameFromPt 获取对象指针的struct的名称{&Struct}
func GetNameFromPt(result interface{}) string {
	return reflect.TypeOf(result).Elem().Name()
}
