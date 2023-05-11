package table_options

import (
	"reflect"
)

type changeCaseFunc func(string) string

func changeCaseForObject(obj any, changeCase changeCaseFunc) {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() == reflect.String {
				nk := changeCase(k.String())
				v := iter.Value()
				changeCaseForObject(v.Interface(), changeCase)
				value.SetMapIndex(k, reflect.Value{})
				value.SetMapIndex(reflect.ValueOf(nk), v)
			}
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			changeCaseForObject(value.Index(i).Interface(), changeCase)
		}
	}
}
