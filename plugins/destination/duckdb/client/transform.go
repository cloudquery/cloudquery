package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func transformRecord(sc *arrow.Schema, rec arrow.Record) arrow.Record {
	cols := make([]arrow.Array, rec.NumCols())
	for i := 0; i < int(rec.NumCols()); i++ {
		cols[i] = transformArray(rec.Column(i))
	}
	return array.NewRecord(sc, cols, rec.NumRows())
}

func transformArray(arr arrow.Array) arrow.Array {
	if arrow.TypeEqual(arr.DataType(), transformTypeForWriting(arr.DataType())) {
		return arr
	}

	switch arr := arr.(type) {
	case *types.UUIDArray, *types.InetArray, *types.MACArray, *types.JSONArray:
		return transformToStringArray(arr)
	case *array.Uint8:
		return transformUint8ToUint32Array(arr)
	case *array.Uint16:
		return transformUint16ToUint32Array(arr)
	case *array.Timestamp:
		// mismatching unit or tz
		return transformTimestamp(duckDBToArrow(arrowToDuckDB(arr.DataType())).(*arrow.TimestampType), arr)

	case *array.Struct:
		dt := arr.DataType().(*arrow.StructType)
		children := make([]arrow.ArrayData, arr.NumField())
		names := make([]string, arr.NumField())
		for i := range children {
			children[i] = transformArray(arr.Field(i)).Data()
			names[i] = dt.Field(i).Name
		}

		return array.NewStructData(array.NewData(
			transformTypeForWriting(dt), arr.Len(),
			arr.Data().Buffers(),
			children,
			arr.NullN(),
			0, // we use 0 as offset for struct arrays, as the child arrays would already be sliced properly
		))

	case array.ListLike: // this includes maps, too
		return array.MakeFromData(array.NewData(
			transformTypeForWriting(arr.DataType()), arr.Len(),
			arr.Data().Buffers(),
			[]arrow.ArrayData{transformArray(arr.ListValues()).Data()},
			arr.NullN(),
			// we use data offset for list like as the `ListValues` can be a larger array (happens when slicing)
			arr.Data().Offset(),
		))
	case *array.Date32:
		return transformDate32ToTimestamp(arr)
	case *array.Date64:
		return transformDate64ToTimestamp(arr)
	default:
		return transformToStringArray(arr)
	}
}

func transformUint16ToUint32Array(arr *array.Uint16) arrow.Array {
	bldr := array.NewUint32Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			bldr.AppendNull()
			continue
		}
		bldr.Append(uint32(arr.Value(i)))
	}
	return bldr.NewArray()
}

func transformUint8ToUint32Array(arr *array.Uint8) arrow.Array {
	bldr := array.NewUint32Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			bldr.AppendNull()
			continue
		}
		bldr.Append(uint32(arr.Value(i)))
	}
	return bldr.NewArray()
}

func transformToStringArray(arr arrow.Array) arrow.Array {
	bldr := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			bldr.AppendNull()
			continue
		}
		bldr.Append(arr.ValueStr(i))
	}
	return bldr.NewArray()
}

func transformTimestamp(dt *arrow.TimestampType, arr *array.Timestamp) arrow.Array {
	builder := array.NewTimestampBuilder(memory.DefaultAllocator, dt)
	in, out := arr.DataType().(*arrow.TimestampType).Unit, dt.Unit

	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}

		builder.Append(arrow.Timestamp(arrow.ConvertTimestampValue(in, out, int64(arr.Value(i)))))
	}

	return builder.NewArray()
}

func transformDate32ToTimestamp(arr *array.Date32) arrow.Array {
	builder := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		builder.AppendTime(arr.Value(i).ToTime())
	}
	return builder.NewArray()
}

func transformDate64ToTimestamp(arr *array.Date64) arrow.Array {
	builder := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		builder.AppendTime(arr.Value(i).ToTime())
	}
	return builder.NewArray()
}
