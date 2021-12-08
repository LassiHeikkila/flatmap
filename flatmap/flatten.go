package flatmap

import (
	"fmt"
	"reflect"
)

// Flatten takes a structure and turns into a flat map[string]string.
//
// Within the "thing" parameter, only primitive values are allowed. Structs are
// not supported. Therefore, it can only be slices, maps, primitives, and
// any combination of those together.
//
// See the tests for examples of what inputs are turned into.
func Flatten(thing map[string]interface{}) Map {
	result := make(map[string]interface{})

	for k, raw := range thing {
		flatten(result, k, reflect.ValueOf(raw))
	}

	return Map(result)
}

func flatten(result Map, prefix string, v reflect.Value) {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Invalid:
		result[prefix] = nil
	case reflect.Bool:
		if v.Bool() {
			result[prefix] = true
		} else {
			result[prefix] = false
		}
	case reflect.Int:
		result[prefix] = int(v.Int())
	case reflect.Int8:
		result[prefix] = int8(v.Int())
	case reflect.Int16:
		result[prefix] = int16(v.Int())
	case reflect.Int32:
		result[prefix] = int32(v.Int())
	case reflect.Int64:
		result[prefix] = int64(v.Int())
	case reflect.Uint:
		result[prefix] = uint(v.Uint())
	case reflect.Uint8:
		result[prefix] = uint8(v.Uint())
	case reflect.Uint16:
		result[prefix] = uint16(v.Uint())
	case reflect.Uint32:
		result[prefix] = uint32(v.Uint())
	case reflect.Uint64:
		result[prefix] = uint64(v.Uint())
	case reflect.Float32:
		result[prefix] = float32(v.Float())
	case reflect.Float64:
		result[prefix] = v.Float()
	case reflect.Map:
		flattenMap(result, prefix, v)
	case reflect.Slice:
		flattenSlice(result, prefix, v)
	case reflect.String:
		result[prefix] = v.String()
	case reflect.Complex64:
		result[prefix] = complex64(v.Complex())
	case reflect.Complex128:
		result[prefix] = v.Complex()
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.UnsafePointer:
		result[prefix] = v.Pointer()
	case reflect.Uintptr:
		result[prefix] = v.Uint()
	default:
		// should just ignore the rest
		return
	}
}

func flattenMap(result Map, prefix string, v reflect.Value) {
	for _, k := range v.MapKeys() {
		if k.Kind() == reflect.Interface {
			k = k.Elem()
		}

		if k.Kind() != reflect.String {
			panic(fmt.Sprintf("%s: map key is not string: %s", prefix, k))
		}

		flatten(result, fmt.Sprintf("%s.%s", prefix, k.String()), v.MapIndex(k))
	}
}

func flattenSlice(result Map, prefix string, v reflect.Value) {
	prefix = prefix + "."

	result[prefix+"#"] = v.Len()
	for i := 0; i < v.Len(); i++ {
		flatten(result, fmt.Sprintf("%s%d", prefix, i), v.Index(i))
	}
}
