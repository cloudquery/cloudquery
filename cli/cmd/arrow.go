package cmd

import (
	"bytes"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/ipc"
)


func RecordToBytes(record arrow.Record) ([]byte, error) {
	var buf bytes.Buffer
	wr := ipc.NewWriter(&buf, ipc.WithSchema(record.Schema()))
	if err := wr.Write(record); err != nil {
		return nil, err
	}
	if err := wr.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NewRecordFromBytes(b []byte) (arrow.Record, error) {
	rdr, err := ipc.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	for rdr.Next() {
		rec := rdr.Record()
		rec.Retain()
		return rec, nil
	}
	return nil, nil
}