package reflection

import (
	"reflect"
)

// func walk(x any, fn func(string)) {
// 	val := getValue(x)
//
// 	numberOfValues := 0
// 	var getField func(int) reflect.Value
//
// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
//
// 	case reflect.Struct:
// 		numberOfValues = val.NumField()
// 		getField = val.Field
//
// 	case reflect.Slice, reflect.Array:
// 		numberOfValues = val.Len()
// 		getField = val.Index
//
// 	case reflect.Map:
// 		for _, key := range val.MapKeys() {
// 			walk(val.MapIndex(key).Interface(), fn)
// 		}
// 	}
//
// 	for i := range numberOfValues {
// 		walk(getField(i).Interface(), fn)
// 	}
// }

func walk(x any, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		for i := range val.NumField() {
			walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}

	case reflect.Func:
		for _, ret := range val.Call(nil) {
			walkValue(ret)
		}

	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
