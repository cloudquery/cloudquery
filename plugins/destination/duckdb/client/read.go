package client

import (
	"context"
	"fmt"
	"io"
	"os"

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
	// sc := table.ToArrowSchema()
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

	_, err = c.db.Exec("copy " + table.Name + " to '" + f.Name() + "' (FORMAT PARQUET)")
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
		castRecs := convertToSingleRowRecords(sc, rec)
		for _, r := range castRecs {
			res <- r
		}
	}
	if rr.Err() != nil && rr.Err() != io.EOF {
		return fmt.Errorf("failed to read parquet record: %w", rr.Err())
	}

	// Create a new scanner to read the file
	// scanner := bufio.NewScanner(f)

	// // Loop through the scanner, reading line by line
	// for scanner.Scan() {
	// 	line := scanner.Bytes()
	// 	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	// 	if err := bldr.UnmarshalJSON(line); err != nil {
	// 		return err
	// 	}
	// 	res <- bldr.NewRecord()
	// }

	// // Check for errors
	// if err := scanner.Err(); err != nil {
	// 	return fmt.Errorf("error reading temporary json file: %s", err)
	// }

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
		col := rec.Column(i)
		switch  {
		case arrow.TypeEqual(sc.Field(i).Type, types.ExtensionTypes.UUID):
			cols[i] = reverseTransformUUID(col.(*array.FixedSizeBinary))
		case arrow.TypeEqual(sc.Field(i).Type, types.ExtensionTypes.Inet):
			cols[i] = reverseTransformInet(col.(*array.String))
		case arrow.TypeEqual(sc.Field(i).Type, types.ExtensionTypes.MAC):
			cols[i] = reverseTransformMAC(col.(*array.String))
		case arrow.TypeEqual(sc.Field(i).Type, arrow.FixedWidthTypes.Timestamp_us):
			cols[i] = reverseTransformTimestamp(sc.Field(i).Type.(*arrow.TimestampType), col.(*array.Timestamp))
		default:
			cols[i] = col
		}
	}
	return array.NewRecord(sc, cols, -1)
}

func reverseTransformArray(f arrow.Field, col arrow.Array) arrow.Array {
	switch  {
	case arrow.TypeEqual(f.Type, types.ExtensionTypes.UUID):
		return reverseTransformUUID(col.(*array.FixedSizeBinary))
	case arrow.TypeEqual(f.Type, types.ExtensionTypes.Inet):
		return reverseTransformInet(col.(*array.String))
	case arrow.TypeEqual(f.Type, types.ExtensionTypes.MAC):
		return reverseTransformMAC(col.(*array.String))
	case arrow.TypeEqual(f.Type, arrow.FixedWidthTypes.Timestamp_us):
		return reverseTransformTimestamp(f.Type.(*arrow.TimestampType), col.(*array.Timestamp))
	case arrow.TypeEqual(f.Type, &arrow.StructType{}):
		// col.(*array.Struct).
		return nil
	default:
		return col
	}
}

func reverseTransformMAC(col *array.String) arrow.Array {
	bldr := types.NewMACBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.ExtensionTypes.MAC))
	for i := 0; i < col.Len(); i++ {
		if !col.IsValid(i) {
			bldr.AppendNull()
		} else {
			bldr.AppendValueFromString(col.Value(i))
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
			bldr.AppendValueFromString(col.Value(i))
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
				bldr.Append(arrow.Timestamp(t.UnixNano()))
			case arrow.Nanosecond:
				bldr.Append(arrow.Timestamp(t.UnixNano()))
			default:
				panic(fmt.Errorf("unsupported timestamp unit: %s", dtype.Unit))
			}
		}
	}
	return bldr.NewTimestampArray()
}