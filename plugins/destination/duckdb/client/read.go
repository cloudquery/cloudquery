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
)

func (c *Client) Read(ctx context.Context, table *schema.Table, _ string, res chan<- arrow.Record) error {
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

	_, err = c.db.ExecContext(ctx, sb.String())
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
	// transform arrays first
	transformed := reverseTransformRecord(sc, rec)

	// slice after
	records := make([]arrow.Record, transformed.NumRows())
	for i := int64(0); i < transformed.NumRows(); i++ {
		records[i] = transformed.NewSlice(i, i+1)
	}

	return records
}

func reverseTransformRecord(sc *arrow.Schema, rec arrow.Record) arrow.Record {
	cols := make([]arrow.Array, rec.NumCols())

	for i := 0; i < int(rec.NumCols()); i++ {
		cols[i] = reverseTransformArray(sc.Field(i).Type, rec.Column(i))
	}

	return array.NewRecord(sc, cols, rec.NumRows())
}

func reverseTransformArray(dt arrow.DataType, arr arrow.Array) arrow.Array {
	switch dt := dt.(type) {
	case *types.UUIDType:
		return array.NewExtensionArrayWithStorage(dt, arr.(*array.FixedSizeBinary))
	case *types.InetType, *types.MACType, *types.JSONType, *arrow.StructType:
		return reverseTransformFromString(dt, arr.(*array.String))
	case *arrow.Uint16Type:
		return reverseTransformUint16(arr.(*array.Uint32))
	case *arrow.Uint8Type:
		return reverseTransformUint8(arr.(*array.Uint32))
	case *arrow.TimestampType:
		return reverseTransformTimestamp(dt, arr.(*array.Timestamp))
	case listLike:
		child := reverseTransformArray(dt.Elem(), arr.(array.ListLike).ListValues()).Data()
		return array.NewListData(array.NewData(dt, arr.Len(), arr.Data().Buffers(), []arrow.ArrayData{child}, arr.NullN(), arr.Data().Offset()))
	default:
		return arr
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

func reverseTransformTimestamp(dt *arrow.TimestampType, arr *array.Timestamp) arrow.Array {
	builder := array.NewTimestampBuilder(memory.DefaultAllocator, dt)
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		t := arr.Value(i).ToTime(arr.DataType().(*arrow.TimestampType).Unit)
		switch dt.Unit {
		case arrow.Second:
			builder.Append(arrow.Timestamp(t.Unix()))
		case arrow.Millisecond:
			builder.Append(arrow.Timestamp(t.UnixMilli()))
		case arrow.Microsecond:
			builder.Append(arrow.Timestamp(t.UnixMicro()))
		case arrow.Nanosecond:
			builder.Append(arrow.Timestamp(t.UnixNano()))
		default:
			panic(fmt.Errorf("unsupported timestamp unit: %s", dt.Unit))
		}
	}
	return builder.NewTimestampArray()
}
