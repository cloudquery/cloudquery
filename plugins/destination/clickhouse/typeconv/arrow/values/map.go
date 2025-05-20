package values

import (
	"fmt"
	"reflect"

	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/ch/types"
)

func buildMap(builder *array.MapBuilder, value any) error {
	// check if the value is actually stored as string (JSON), if so - the stored value is a single marshaled JSON entry
	colType, err := types.ColumnType(builder.Type())
	if err != nil {
		return err
	}

	if colType == "String" {
		return buildFromString(builder, value)
	}

	return buildMapFromReflect(builder, reflect.ValueOf(value))
}

func buildMapFromReflect(builder *array.MapBuilder, value reflect.Value) error {
	for value.Kind() == reflect.Pointer {
		if value.IsNil() {
			builder.AppendNull()
			return nil
		}
		value = value.Elem()
	}

	// non-nil map pointer
	if value.Kind() != reflect.Map {
		return fmt.Errorf("unexpected value type %q for map", value.Kind())
	}
	if value.IsNil() {
		builder.AppendNull()
		return nil
	}

	builder.Append(true)
	keyBuilder, itemBuilder := builder.KeyBuilder(), builder.ItemBuilder()
	it := value.MapRange()
	for it.Next() {
		if err := buildValue(keyBuilder, it.Key().Interface()); err != nil {
			return err
		}
		if err := buildValue(itemBuilder, it.Value().Interface()); err != nil {
			return err
		}
	}

	return nil
}
