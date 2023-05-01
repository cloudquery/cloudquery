package json

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
)

var (
	ErrMismatchFields = errors.New("json: number of records mismatch")
)

// Writer wraps json.Encoder and writes arrow.Record based on a schema.
type Writer struct {
	schema *arrow.Schema
	w      *json.Encoder
}

func NewWriter(w io.Writer, schema *arrow.Schema) *Writer {
	ww := &Writer{
		schema: schema,
		w:      json.NewEncoder(w),
	}
	ww.w.SetEscapeHTML(false)

	return ww
}

// Write writes a single Record as one row to the CSV file
func (w *Writer) Write(record arrow.Record) error {
	if !record.Schema().Equal(w.schema) {
		return ErrMismatchFields
	}
	nc := int(record.NumCols())
	nr := int(record.NumRows())
	tmp := make(map[string]any, nc)

	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			switch record.Column(j).DataType().ID() {
			case arrow.TIMESTAMP:
				tmp[w.schema.Field(j).Name] = record.Column(j).GetOneForMarshal(i)
				if tmp[w.schema.Field(j).Name] != nil {
					strTimestamp := tmp[w.schema.Field(j).Name].(string)
					// I believe the real fix should be in how duckdb is handling timestamp as it seems like a bug
					// so this is just a workaround.
					// in anycase we do want a capability of custom marshaling where needed.
					if ind := strings.Index(strTimestamp, "."); ind == -1 {
						tmp[w.schema.Field(j).Name] = strTimestamp + ".000000"
					} else if len(strTimestamp[ind+1:]) < 6 {
						tmp[w.schema.Field(j).Name] = strTimestamp + "000000"[len(strTimestamp[ind+1:]):]
					}
				}
			default:
				tmp[w.schema.Field(j).Name] = record.Column(j).GetOneForMarshal(i)
			}
		}
		if err := w.w.Encode(tmp); err != nil {
			return err
		}
		tmp = make(map[string]any, nc)
	}

	return nil
}
