package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func transformRecord(sc *arrow.Schema, rec arrow.Record) arrow.Record {
	cols := make([]arrow.Array, rec.NumCols())
	for i := 0; i < int(rec.NumCols()); i++ {
		cols[i] = transformArray(rec.Column(i))
	}
	return array.NewRecord(sc, cols, rec.NumRows())
}

func transformArray(arr arrow.Array) arrow.Array {
	dt := arr.DataType()
	switch {
	case arrow.TypeEqual(dt, types.ExtensionTypes.UUID) ||
		arrow.TypeEqual(dt, types.ExtensionTypes.Inet) ||
		arrow.TypeEqual(dt, types.ExtensionTypes.MAC) ||
		arrow.TypeEqual(dt, types.ExtensionTypes.JSON) ||
		dt.ID() == arrow.STRUCT:
		return transformToStringArray(arr)
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint8):
		return transformUint8ToUint32Array(arr.(*array.Uint8))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint16):
		return transformUint16ToUint32Array(arr.(*array.Uint16))
	case arrow.IsListLike(dt.ID()):
		child := transformArray(arr.(*array.List).ListValues()).Data()
		newType := arrow.ListOf(child.DataType())
		return array.NewListData(array.NewData(newType, arr.Len(), arr.Data().Buffers(), []arrow.ArrayData{child}, arr.NullN(), arr.Data().Offset()))
	default:
		return arr
	}
}

func transformUint16ToUint32Array(arr *array.Uint16) arrow.Array {
	bldr := array.NewUint32Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			bldr.Append(uint32(arr.Value(i)))
		} else {
			bldr.AppendNull()
		}
	}
	return bldr.NewArray()
}

func transformUint8ToUint32Array(arr *array.Uint8) arrow.Array {
	bldr := array.NewUint32Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			bldr.Append(uint32(arr.Value(i)))
		} else {
			bldr.AppendNull()
		}
	}
	return bldr.NewArray()
}

func transformToStringArray(arr arrow.Array) arrow.Array {
	bldr := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			bldr.Append(arr.ValueStr(i))
		} else {
			bldr.AppendNull()
		}
	}
	return bldr.NewArray()
}
