package main

import (
	"reflect"
)

func walk(v interface{}, fn func(input string)) {
	val := getValue(v)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		// struct have method NumField() and Field()
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	case reflect.Array, reflect.Slice:
		// slice and array have method Len() and Index()
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		// map have method MapKeys() and MapIndex()
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	}
}

func getValue(v interface{}) reflect.Value {
	val := reflect.ValueOf(v)

	// if val.Kind() is pointer, we need get the underlying value of val
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
