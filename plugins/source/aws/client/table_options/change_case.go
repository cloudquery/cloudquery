package table_options

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/caser"
)

func changeCaseForObject(obj any) {
	csr := caser.New()
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() == reflect.String {
				nk := csr.ToPascal(k.String())
				v := iter.Value()
				changeCaseForObject(v.Interface())
				value.SetMapIndex(k, reflect.Value{})
				value.SetMapIndex(reflect.ValueOf(nk), v)
			}
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			changeCaseForObject(value.Index(i).Interface())
		}
	}
}
