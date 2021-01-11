package meistruct

import "reflect"

// GetName 通过struct{}获取类型名称
func GetName(result interface{}) string {
	return reflect.TypeOf(result).Name()
}

// GetNameFromPt 通过*struct{}指针获取类型名称
func GetNameFromPt(result interface{}) string {
	return reflect.TypeOf(result).Elem().Name()
}
