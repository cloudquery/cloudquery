package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet/file"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.parquet", table.Name))
	if err != nil {
		return err
	}

	defer os.Remove(f.Name())
	sc := table.ToArrowSchema()
	fName := f.Name()
	if err := f.Close(); err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString("copy " + sanitizeID(table.Name) + " (")
	for i, col := range sc.Fields() {
		sb.WriteString(sanitizeID(col.Name))
		if i < len(sc.Fields())-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") to '" + f.Name() + "' (FORMAT PARQUET)")

	if err := c.exec(ctx, sb.String()); err != nil {
		return err
	}
	f, err = os.Open(fName)
	if err != nil {
		return err
	}

	rdr, err := file.NewParquetReader(f)
	if err != nil {
		return fmt.Errorf("failed to create new parquet reader: %w", err)
	}
	arrProps := pqarrow.ArrowReadProperties{
		Parallel:  false,
		BatchSize: 1024,
	}
	fr, err := pqarrow.NewFileReader(rdr, arrProps, memory.DefaultAllocator)
	if err != nil {
		return fmt.Errorf("failed to create new parquet file reader: %w", err)
	}
	rr, err := fr.GetRecordReader(ctx, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to get parquet record reader: %w", err)
	}

	for rr.Next() {
		for _, r := range slice(reverseTransformRecord(sc, rr.Record())) {
			res <- r
		}
	}
	if rr.Err() != nil && rr.Err() != io.EOF {
		return fmt.Errorf("failed to read parquet record: %w", rr.Err())
	}

	return nil
}

func slice(r arrow.Record) []arrow.Record {
	res := make([]arrow.Record, r.NumRows())
	for i := int64(0); i < r.NumRows(); i++ {
		res[i] = r.NewSlice(i, i+1)
	}
	return res
}

func reverseTransformRecord(sc *arrow.Schema, rec arrow.Record) arrow.Record {
	cols := make([]arrow.Array, rec.NumCols())

	for i := 0; i < int(rec.NumCols()); i++ {
		cols[i] = reverseTransformArray(sc.Field(i).Type, rec.Column(i))
	}

	return array.NewRecord(sc, cols, rec.NumRows())
}

func reverseTransformArray(dt arrow.DataType, arr arrow.Array) arrow.Array {
	if arrow.TypeEqual(dt, arr.DataType()) {
		return arr
	}

	switch dt := dt.(type) {
	case *types.UUIDType:
		return array.NewExtensionArrayWithStorage(dt, arr.(*array.FixedSizeBinary))
	case *types.InetType, *types.MACType:
		return reverseTransformFromString(dt, arr.(*array.String))
	case *arrow.Uint16Type:
		return reverseTransformUint16(arr.(*array.Uint32))
	case *arrow.Uint8Type:
		return reverseTransformUint8(arr.(*array.Uint32))
	case *arrow.TimestampType:
		return transformTimestamp(dt, arr.(*array.Timestamp))
	case *arrow.Date32Type:
		// We save date types as Timestamp
		return reverseTransformDate32(arr.(*array.Timestamp))
	case *arrow.Date64Type:
		// We save date types as Timestamp
		return reverseTransformDate64(arr.(*array.Timestamp))
	case *arrow.StructType:
		if sarr, ok := arr.(*array.Binary); ok {
			return reverseTransformStruct(dt, sarr)
		}

		arr := arr.(*array.Struct)
		children := make([]arrow.ArrayData, arr.NumField())
		for i := range children {
			// struct fields can be odd when read from parquet, but the data is intact
			child := array.MakeFromData(arr.Data().Children()[i])
			children[i] = reverseTransformArray(dt.Field(i).Type, child).Data()
		}

		return array.NewStructData(array.NewData(
			dt, arr.Len(),
			arr.Data().Buffers(),
			children,
			arr.NullN(),
			0, // we use 0 as offset for struct arrays, as the child arrays would already be sliced properly
		))
	case arrow.ListLikeType: // also handles maps
		if mapdt, ok := dt.(*arrow.MapType); ok {
			if sarr, ok := arr.(*array.Binary); ok {
				return reverseTransformMap(mapdt, sarr)
			}
		}

		return array.MakeFromData(array.NewData(
			dt, arr.Len(),
			arr.Data().Buffers(),
			[]arrow.ArrayData{reverseTransformArray(dt.Elem(), arr.(array.ListLike).ListValues()).Data()},
			arr.NullN(),
			// we use data offset for list like as the `ListValues` can be a larger array (happens when slicing)
			arr.Data().Offset(),
		))
	case *types.JSONType:
		jsonArray := arr.(*array.Binary)
		jsonBuilder := types.NewJSONBuilder(memory.DefaultAllocator)
		for i := 0; i < jsonArray.Len(); i++ {
			if arr.IsNull(i) {
				jsonBuilder.AppendNull()
			} else {
				jsonBuilder.AppendBytes(jsonArray.Value(i))
			}
		}
		return jsonBuilder.NewJSONArray()
	case *arrow.BinaryType:
		binaryArray := arr.(*array.Binary)
		binaryBuilder := array.NewBinaryBuilder(memory.DefaultAllocator, dt)
		for i := 0; i < binaryArray.Len(); i++ {
			if binaryArray.IsNull(i) {
				binaryBuilder.AppendNull()
			} else {
				binaryBuilder.Append(binaryArray.Value(i))
			}
		}
		return binaryBuilder.NewLargeBinaryArray()
	case *arrow.LargeBinaryType:
		largeBinaryArray := arr.(*array.Binary)
		largeBinaryBuilder := array.NewBinaryBuilder(memory.DefaultAllocator, dt)
		for i := 0; i < largeBinaryArray.Len(); i++ {
			if largeBinaryArray.IsNull(i) {
				largeBinaryBuilder.AppendNull()
			} else {
				largeBinaryBuilder.Append(largeBinaryArray.Value(i))
			}
		}
		return largeBinaryBuilder.NewLargeBinaryArray()
	default:
		return reverseTransformFromString(dt, arr.(*array.String))
	}
}

func reverseTransformFromString(dt arrow.DataType, arr *array.String) arrow.Array {
	builder := array.NewBuilder(memory.DefaultAllocator, dt)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		if err := builder.AppendValueFromString(arr.Value(i)); err != nil {
			panic(fmt.Errorf("failed to append from string value %q: %w", arr.Value(i), err))
		}
	}

	return builder.NewArray()
}

func reverseTransformStruct(dt *arrow.StructType, arr *array.Binary) arrow.Array {
	bldr := array.NewStructBuilder(memory.DefaultAllocator, dt)
	defer bldr.Release()
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			bldr.AppendNull()
			continue
		}
		if err := bldr.AppendValueFromString(arr.ValueString(i)); err != nil {
			panic(err)
		}
	}
	return bldr.NewStructArray()
}

func reverseTransformMap(dt *arrow.MapType, arr *array.Binary) arrow.Array {
	bldr := array.NewMapBuilder(memory.DefaultAllocator, dt.KeyType(), dt.ItemType(), dt.KeysSorted)
	defer bldr.Release()
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			bldr.AppendNull()
			continue
		}
		if err := bldr.AppendValueFromString(arr.ValueString(i)); err != nil {
			panic(err)
		}
	}
	return bldr.NewMapArray()
}
func reverseTransformUint8(arr *array.Uint32) arrow.Array {
	builder := array.NewUint8Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		builder.Append(uint8(arr.Value(i)))
	}

	return builder.NewArray()
}

func reverseTransformUint16(arr *array.Uint32) arrow.Array {
	builder := array.NewUint16Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		builder.Append(uint16(arr.Value(i)))
	}

	return builder.NewArray()
}

func reverseTransformDate32(arr *array.Timestamp) arrow.Array {
	builder := array.NewDate32Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		builder.Append(arrow.Date32FromTime(arr.Value(i).ToTime(arrow.Microsecond)))
	}

	return builder.NewArray()
}

func reverseTransformDate64(arr *array.Timestamp) arrow.Array {
	builder := array.NewDate64Builder(memory.DefaultAllocator)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		builder.Append(arrow.Date64FromTime(arr.Value(i).ToTime(arrow.Microsecond)))
	}

	return builder.NewArray()
}
