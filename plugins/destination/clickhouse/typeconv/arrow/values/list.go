package values

import (
	"fmt"
	"reflect"

	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildListValues(builder array.ListLikeBuilder, values any) error {
	if values == nil {
		return nil
	}

	slice := reflect.Indirect(reflect.ValueOf(values))
	switch slice.Kind() {
	case reflect.Array, reflect.Slice:
	// expected, continue
	default:
		return fmt.Errorf("unsupported type %T for %s", values, builder.Type().String())
	}

	for i := 0; i < slice.Len(); i++ {
		builder.Append(true)
		if err := buildValue(builder.ValueBuilder(), slice.Index(i).Interface()); err != nil {
			return err
		}
	}

	return nil
}
