package flatten

import "reflect"

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}
	if reflect.DeepEqual(nested, []interface{}{}) {
		return result
	}
	slice := reflect.ValueOf(nested)
	if slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {

		}
	}
	return result
}

func newFlatter()
