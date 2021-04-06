package meistruct

import (
	"reflect"

	"github.com/andytyc/goutils/meimap"
)

// Bson2Map | Struct => Map
// 通过Struct的字段tag标签(bson)，来转换为字典键值对
// 场景：mongodb使用时，对应的bson字段
func Bson2Map(obj interface{}, checks, dels []string, idpass bool) map[string]interface{} {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		bsonVal, ok := t.Field(i).Tag.Lookup("bson")
		if ok && bsonVal != "-" {
			if bsonVal == "_id" {
				if idpass {
					continue
				}
				if !v.Field(i).IsZero() {
					data["_id"] = v.Field(i).Interface()
				} else {
					return nil // if idpass is false, _id is not allow to be empty.
				}
			} else {
				data[bsonVal] = v.Field(i).Interface()
			}
		}
	}

	checkData := meimap.StringPickByKeys(data, checks)
	delData := meimap.StringDelByKeys(checkData, dels)

	return delData
}
