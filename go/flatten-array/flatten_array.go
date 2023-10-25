package flatten

import "reflect"

func Flatten(nested interface{}) []interface{} {
	var result []interface{}
	value := reflect.ValueOf(nested)
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			result = append(result, Flatten(value.Index(i).Interface())...)
		}
	} else {
		result = append(result, nested)
	}
	return result
}
