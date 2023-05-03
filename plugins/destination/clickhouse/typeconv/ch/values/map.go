package values

import (
	"encoding/json"
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
		val := reflect.New(valueType)
		// we need to fill in for the in-depth recursive parsing by ClickHouse SDK
		if arr.IsNull(i) {
			res.Index(i).Set(val)
			continue
		}

		rowKeys, rowItems := keysValue.Index(i).Elem(), itemsValue.Index(i).Elem() // ->[]*
		for idx := 0; idx < rowKeys.Len(); idx++ {
			// this is matched exactly by the items
			reflect.Ma
		}
		from, to := arr.ValueOffsets(i)
		elems[i], err = FromArray(array.NewSlice(arr.ListValues(), from, to))
		if err != nil {
			return nil, err
		}
	}
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
