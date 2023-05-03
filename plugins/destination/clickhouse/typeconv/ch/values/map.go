package values

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
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

	keysArr, itemsArr := arr.Keys(), arr.Items()
	keys, err := FromArray(keysArr) // []*[]*...
	if err != nil {
		return nil, err
	}
	items, err := FromArray(itemsArr) // []*[]*...
	if err != nil {
		return nil, err
	}
	keysValue, itemsValue := reflect.ValueOf(keys), reflect.ValueOf(items)

	res := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(valueType)), arr.Len(), arr.Len()) // we do []*(type) for nullable assignment
	for i := 0; i < arr.Len(); i++ {
		var val reflect.Value
		// we need to fill in for the in-depth recursive parsing by ClickHouse SDK
		if arr.IsNull(i) {
			val = reflect.MakeMap(valueType) // zero-sized map
		} else {
			val = makeMap(valueType, keysValue.Index(i).Elem(), itemsValue.Index(i).Elem())
		}
		res.Index(i).Set(val)
	}
	return res.Interface(), nil
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

func makeMap(mapType reflect.Type, keys, items reflect.Value) reflect.Value {
	val := reflect.MakeMapWithSize(mapType, keys.Len())
	for i := 0; i < keys.Len(); i++ {
		// Arrow maps don't support nullable keys, so no need to check
		val.SetMapIndex(keys.Index(i).Elem(), items.Index(i))
	}
	return val
}
