package meistruct

import (
	"reflect"

	"github.com/andytyc/goutils/meimap"
)

// Bson2Map | Struct => Map
// Sturct object convert to a map according to bson tag of struct
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
