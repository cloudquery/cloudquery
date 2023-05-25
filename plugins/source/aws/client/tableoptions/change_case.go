package tableoptions

import (
	"reflect"

	"golang.org/x/exp/slices"
)

type changeCaseFunc func(string) string

// skipFields is a list of fields that should not be changed. This is useful for fields that are
// maps, where case needs to be preserved. Right now skipFields only supports top level fields,
// but recursive support could be added if needed later.
func changeCaseForObject(obj any, changeCase changeCaseFunc, skipFields ...string) {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() == reflect.String {
				nk := changeCase(k.String())
				v := iter.Value()
				if slices.Contains(skipFields, k.String()) {
					continue
				}
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
