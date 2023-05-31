package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/apache/arrow/go/v13/parquet/file"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/google/uuid"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- arrow.Record) error {
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
	sb.WriteString("copy " + table.Name + "(")
	for i, col := range sc.Fields() {
		sb.WriteString("\"" + col.Name + "\"")
		if i < len(sc.Fields())-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") to '" + f.Name() + "' (FORMAT PARQUET)")

	_, err = c.db.Exec(sb.String())
	if err != nil {
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
		rec := rr.Record()
		rec.Retain()
		castRecs := convertToSingleRowRecords(sc, rec)
		for _, r := range castRecs {
			res <- r
		}
	}
	if rr.Err() != nil && rr.Err() != io.EOF {
		return fmt.Errorf("failed to read parquet record: %w", rr.Err())
	}

	return nil
}

func convertToSingleRowRecords(sc *arrow.Schema, rec arrow.Record) []arrow.Record {
	records := make([]arrow.Record, rec.NumRows())
	for i := int64(0); i < rec.NumRows(); i++ {
		records[i] = reverseTransformRecord(sc, rec.NewSlice(i, i+1))
	}
	return records
}

func reverseTransformRecord(sc *arrow.Schema, rec arrow.Record) arrow.Record {
	cols := make([]arrow.Array, rec.NumCols())
	for i := 0; i < int(rec.NumCols()); i++ {
		cols[i] = reverseTransformArray(sc.Field(i).Type, rec.Column(i))
	}
	return array.NewRecord(sc, cols, -1)
}

func reverseTransformArray(dt arrow.DataType, col arrow.Array) arrow.Array {
	switch {
	case arrow.TypeEqual(dt, types.ExtensionTypes.UUID):
		return reverseTransformUUID(col.(*array.FixedSizeBinary))
	case arrow.TypeEqual(dt, types.ExtensionTypes.Inet):
		return reverseTransformInet(col.(*array.String))
	case arrow.TypeEqual(dt, types.ExtensionTypes.MAC):
		return reverseTransformMAC(col.(*array.String))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint16):
		return reverseTransformUint16(col.(*array.Uint32))
	case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint8):
		return reverseTransformUint8(col.(*array.Uint32))
	case arrow.TypeEqual(dt, types.ExtensionTypes.JSON):
		return reverseTransformJSON(col.(*array.String))
	case arrow.TypeEqual(col.DataType(), arrow.FixedWidthTypes.Timestamp_us):
		return reverseTransformTimestamp(dt.(*arrow.TimestampType), col.(*array.Timestamp))
	case dt.ID() == arrow.STRUCT:
		return reverseTransformStruct(dt.(*arrow.StructType), col.(*array.String))
	case arrow.IsListLike(dt.ID()):
		child := reverseTransformArray(dt.(*arrow.ListType).Elem(), col.(*array.List).ListValues()).Data()
		return array.NewListData(array.NewData(dt, col.Len(), col.Data().Buffers(), []arrow.ArrayData{child}, col.NullN(), col.Data().Offset()))
	default:
		return col
	}
}

func reverseTransformStruct(dt *arrow.StructType, col *array.String) arrow.Array {
	bldr := array.NewStructBuilder(memory.DefaultAllocator, dt)
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			if err := bldr.AppendValueFromString(col.Value(i)); err != nil {
				panic(fmt.Errorf("failed to append json %s value: %w", col.Value(i), err))
			}
		}
	}

	return bldr.NewArray()
}

func reverseTransformJSON(col *array.String) arrow.Array {
	bldr := types.NewJSONBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.ExtensionTypes.JSON))
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			if err := bldr.AppendValueFromString(col.Value(i)); err != nil {
				panic(fmt.Errorf("failed to append json %s value: %w", col.Value(i), err))
			}
		}
	}

	return bldr.NewArray()
}

func reverseTransformUint8(col *array.Uint32) arrow.Array {
	bldr := array.NewUint8Builder(memory.DefaultAllocator)
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			bldr.Append(uint8(col.Value(i)))
		}
	}

	return bldr.NewArray()
}

func reverseTransformUint16(col *array.Uint32) arrow.Array {
	bldr := array.NewUint16Builder(memory.DefaultAllocator)
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			bldr.Append(uint16(col.Value(i)))
		}
	}

	return bldr.NewArray()
}

func reverseTransformMAC(col *array.String) arrow.Array {
	bldr := types.NewMACBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.ExtensionTypes.MAC))
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			if err := bldr.AppendValueFromString(col.Value(i)); err != nil {
				panic(err)
			}
		}
	}

	return bldr.NewMACArray()
}

func reverseTransformInet(col *array.String) arrow.Array {
	bldr := types.NewInetBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.ExtensionTypes.Inet))
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			if err := bldr.AppendValueFromString(col.Value(i)); err != nil {
				panic(err)
			}
		}
	}

	return bldr.NewInetArray()
}

func reverseTransformUUID(col *array.FixedSizeBinary) arrow.Array {
	bldr := types.NewUUIDBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.ExtensionTypes.UUID))
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			u, err := uuid.FromBytes(col.Value(i))
			if err != nil {
				panic(err)
			}
			bldr.Append(u)
		}
	}

	return bldr.NewUUIDArray()
}

func reverseTransformTimestamp(dtype *arrow.TimestampType, col *array.Timestamp) arrow.Array {
	bldr := array.NewTimestampBuilder(memory.DefaultAllocator, dtype)
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			t := col.Value(i).ToTime(col.DataType().(*arrow.TimestampType).Unit)
			switch dtype.Unit {
			case arrow.Second:
				bldr.Append(arrow.Timestamp(t.Unix()))
			case arrow.Millisecond:
				bldr.Append(arrow.Timestamp(t.UnixMilli()))
			case arrow.Microsecond:
				bldr.Append(arrow.Timestamp(t.UnixMicro()))
			case arrow.Nanosecond:
				bldr.Append(arrow.Timestamp(t.UnixNano()))
			default:
				panic(fmt.Errorf("unsupported timestamp unit: %s", dtype.Unit))
			}
		}
	}
	return bldr.NewTimestampArray()
}
