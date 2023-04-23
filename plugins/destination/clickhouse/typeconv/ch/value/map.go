package value

import (
	"reflect"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/definitions"
)

func mapValue(arr *array.Map) (any, error) {
	// we need to check if we really want the map or string
	def, err := definitions.MapType(arr.DataType().(*arrow.MapType))
	if err != nil {
		return nil, err
	}
	if def == "String" {
		return valueStrData(arr), nil
	}

	// we really do need map here, so we need to construct the proper type to be scanned as well

	// keys won't be nullable, but the FromArray output will render them as such
	// We'll account for that later
	keys, err := FromArray(arr.Keys())
	if err != nil {
		return nil, err
	}

	items, err := FromArray(arr.Items())
	if err != nil {
		return nil, err
	}

	// keys & items are []*[]*(type)
	keySlices, err := toSlice(keys)
	if err != nil {
		return nil, err
	}
	itemSlices, err := toSlice(items)
	if err != nil {
		return nil, err
	}

	keyScanType, err := scanType(arr.DataType().(*arrow.MapType).KeyField())
	if err != nil {
		return nil, err
	}

	itemScanType, err := scanType(arr.DataType().(*arrow.MapType).ItemField())
	if err != nil {
		return nil, err
	}

	mapType := reflect.MapOf(keyScanType, itemScanType)
	result := reflect.MakeSlice(reflect.PointerTo(mapType), arr.Len(), arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			continue
		}

		row := reflect.New()

	}
}

func scanType(field arrow.Field) (reflect.Type, error) {
	fieldType, err := definitions.FieldType(field)
	if err != nil {
		return nil, err
	}

	col, err := column.Type(fieldType).Column(field.Name, time.UTC)
	if err != nil {
		return nil, err
	}

	return col.ScanType(), nil
}
