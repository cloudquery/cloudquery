package values

import (
	"encoding/json"
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
		return marshalValuesToStrings(arr), nil
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

func marshalValuesToStrings(arr *array.Map) []*string {
	data := marshalValue[json.RawMessage](arr)
	res := make([]*string, len(data))
	for i, elem := range data {
		if elem == nil {
			continue
		}
		str := string(*elem)
		res[i] = &str
	}
	return res
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
		mapVal, err := makeMapWithList(mapType, array.NewSlice(arr.ListValues(), start, end))
		//keys := array.NewSlice(arr.Keys(), start, end)
		//items := array.NewSlice(arr.Items(), start, end)
		//mapVal, err := makeMapKV(mapType, keys, items)
		if err != nil {
			return nil, err
		}
		val.Elem().Set(*mapVal)
		res.Index(i).Set(val)
	}
	return res.Interface(), nil
}

func makeMapKV(mapType reflect.Type, keyArr, itemArr arrow.Array) (*reflect.Value, error) {
	keysVal, err := FromArray(keyArr)
	if err != nil {
		return nil, err
	}
	keys := reflect.ValueOf(keysVal)

	itemsVal, err := FromArray(itemArr)
	if err != nil {
		return nil, err
	}
	items := reflect.ValueOf(itemsVal)

	if keys.Len() != items.Len() {
		return nil, fmt.Errorf("keys len (%d) != items len (%d)", keys.Len(), items.Len())
	}

	val := reflect.MakeMapWithSize(mapType, keys.Len())
	for i := 0; i < keys.Len(); i++ {
		key, item := keys.Index(i).Elem(), items.Index(i)
		// we unwrap the item only if it's nested type: map or slice/array
		if item.Kind() == reflect.Pointer {
			switch item.Type().Elem().Kind() {
			case reflect.Map, reflect.Slice, reflect.Array:
				item = item.Elem()
			case reflect.Invalid: // we encountered the nil value, still we need to unwrap

			}
		}
		// Arrow maps don't support nullable keys, so no need to check
		// BEWARE: reflect.Value.SetMapIndex deletes the value from map if it's the zero value (nil)
		val.SetMapIndex(key, item)
	}
	return &val, nil
}

func makeMapWithList(mapType reflect.Type, arr arrow.Array) (*reflect.Value, error) {
	data, err := FromArray(arr)
	if err != nil {
		return nil, err
	}
	// we do know that this is []*[]*map[string]any (map is implemented as list of structs(key, item))
	actualData := data.([]*map[string]any)

	const (
		keyField  = "key"
		itemField = "value"
	)

	value := reflect.MakeMapWithSize(mapType, len(actualData))
	for _, elem := range actualData {
		// elem should NEVER be nil (at least key has to be filled in)
		key := reflect.ValueOf((*elem)[keyField]).Elem()
		item := reflect.ValueOf((*elem)[itemField])
		if item.Kind() == reflect.Pointer {
			switch item.Type().Elem().Kind() {
			case reflect.Map, reflect.Slice, reflect.Array:
				if item.IsNil() {
					item = reflect.New(item.Type().Elem()).Elem()
				} else {
					item = item.Elem()
				}
			}
		}
		value.SetMapIndex(key, item)
	}
	return &value, nil
}
