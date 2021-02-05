package meistruct

import "reflect"

// GetName | get the name of the struct, it is a object
func GetName(result interface{}) string {
	return reflect.TypeOf(result).Name()
}

// GetNameFromPt | get the name of the struct, it is a ptr of object
func GetNameFromPt(result interface{}) string {
	return reflect.TypeOf(result).Elem().Name()
}
