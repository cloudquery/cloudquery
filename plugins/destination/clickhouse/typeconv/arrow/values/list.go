package values

import (
	"fmt"
	"reflect"

	"github.com/apache/arrow-go/v18/arrow/array"
)

func buildList(builder array.ListLikeBuilder, values any) error {
	slice := reflect.Indirect(reflect.ValueOf(values))
	if slice.IsNil() {
		builder.AppendNull()
		return nil
	}

	switch slice.Kind() {
	case reflect.Array, reflect.Slice:
	// expected, continue
	default:
		return fmt.Errorf("unsupported type %T for %s", values, builder.Type().String())
	}

	builder.Append(true)
	for i := 0; i < slice.Len(); i++ {
		if err := buildValue(builder.ValueBuilder(), slice.Index(i).Interface()); err != nil {
			return err
		}
	}

	return nil
}
