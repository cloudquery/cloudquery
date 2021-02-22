package keyvals

import (
	"fmt"
	"reflect"
)

// ToMap creates a map of key-value pairs from a variadic key-value pair slice.
//
// The implementation bellow is from go-kit's JSON logger.
func ToMap(kvs []interface{}) map[string]interface{} {
	m := map[string]interface{}{}

	if len(kvs) == 0 {
		return m
	}

	if len(kvs)%2 == 1 {
		kvs = append(kvs, nil)
	}

	for i := 0; i < len(kvs); i += 2 {
		merge(m, kvs[i], kvs[i+1])
	}

	return m
}

func merge(dst map[string]interface{}, k, v interface{}) {
	var key string

	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	dst[key] = v
}

func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(panicVal)
			}
		}
	}()

	s = str.String()

	return
}
