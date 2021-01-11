package meistruct

import (
	"reflect"

	"github.com/andytyc/goutils/meimap"
)

// Bson2Map 对象struct根据bson转换至字典
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
					return nil // 在对象转换键值对时，发现_id为空，请核查
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
