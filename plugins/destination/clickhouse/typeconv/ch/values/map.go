package values

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
)

func mapValue(arr *array.Map) (any, error) {
	// check if we really need to construct map
	colType, err := types.ColumnType(arr.DataType())
	if err != nil {
		return nil, err
	}
	if colType == "String" {
		return marshalValuesToStrings(arr)
	}

	// Need to create slice of the proper type.
	// We could infer in from elements, but sometimes array is empty
	col, err := column.Type(colType).Column("tmp", time.UTC)
	if err != nil {
		return nil, err
	}
	valueType := col.ScanType()
	if valueType.Kind() != reflect.Map {
		return nil, fmt.Errorf("unexpected reflect type for map: %q", valueType.String())
	}

	return makeMapSlice(valueType, arr)
}

func makeMapSlice(mapType reflect.Type, arr *array.Map) (any, error) {
	res := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(mapType)), arr.Len(), arr.Len()) // we do []*(type) for nullable assignment
	for i := 0; i < arr.Len(); i++ {
		val := reflect.New(mapType)
		if arr.IsNull(i) {
			// we need to fill in for the in-depth recursive parsing by ClickHouse SDK
			res.Index(i).Set(val)
			continue
		}
		start, end := arr.ValueOffsets(i)
		mapVal, err := makeMap(mapType, array.NewSlice(arr.ListValues(), start, end))
		if err != nil {
			return nil, err
		}
		val.Elem().Set(*mapVal)
		res.Index(i).Set(val)
	}
	return res.Interface(), nil
}

func makeMap(mapType reflect.Type, arr arrow.Array) (*reflect.Value, error) {
	data, err := FromArray(arr)
	if err != nil {
		return nil, err
	}
	// we do know that this is []*map[string]any (map is implemented as list of structs(key, item))
	actualData := data.([]*map[string]any)

	const (
		keyField  = "key"
		itemField = "value"
	)

	value := reflect.MakeMapWithSize(mapType, len(actualData))
	for _, elem := range actualData {
		// elem should NEVER be nil (at least key has to be filled in)
		value.SetMapIndex(reflect.ValueOf((*elem)[keyField]).Elem(), mapItemValue(reflect.ValueOf((*elem)[itemField])))
	}
	return &value, nil
}

// mapItemValue adds logic to unwrap value stored in map, if the value is a pointer to map or slice
func mapItemValue(item reflect.Value) reflect.Value {
	if item.Kind() != reflect.Pointer {
		return item
	}
	switch item.Type().Elem().Kind() {
	case reflect.Map, reflect.Slice:
		if item.IsNil() {
			return reflect.New(item.Type().Elem()).Elem()
		}
		return item.Elem()
	default:
		return item
	}
}
