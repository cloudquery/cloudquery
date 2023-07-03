package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/segmentio/fasthash/fnv1a"
)

type bulkResponse struct {
	Took   int64 `json:"took"`
	Errors bool  `json:"errors"`
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return fmt.Errorf("failed to write messages: %w", err)
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	for _, msg := range msgs {
		table := msg.GetTable()
		record := msg.Record
		err := c.writeRecord(ctx, table, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) writeRecord(ctx context.Context, table *schema.Table, record arrow.Record) error {
	var buf bytes.Buffer
	pks := pkIndexes(table) // do some work up front to avoid doing it for every resource
	// get the sync time from the first resource in the batch (here we assume that all resources in the batch
	// have the same sync time. At the moment this assumption holds.)
	syncTime := time.Now()
	for r := 0; r < int(record.NumRows()); r++ {
		doc := map[string]any{}
		for i, col := range record.Columns() {
			doc[table.Columns[i].Name] = c.getValueForElasticsearch(col, r)
		}
		data, err := json.Marshal(doc)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		var meta []byte
		hasPrimaryKeys := len(table.PrimaryKeys()) > 0

		if hasPrimaryKeys {
			docID := fmt.Sprint(resourceID(record, r, pks))
			meta = []byte(fmt.Sprintf(`{"index":{"_id":"%s"}}%s`, docID, "\n"))
		} else {
			meta = []byte(`{"index":{}}` + "\n")
		}
		data = append(data, "\n"...)
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}
	index := c.getIndexName(table, syncTime)
	resp, err := c.client.Bulk(bytes.NewReader(buf.Bytes()),
		c.client.Bulk.WithContext(ctx),
		c.client.Bulk.WithIndex(index),
	)
	if err != nil {
		return fmt.Errorf("failed to create bulk request: %w", err)
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return fmt.Errorf("bulk request failed: %s", resp.String())
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read bulk response: %w", err)
	}
	var bulkResp bulkResponse
	err = json.Unmarshal(b, &bulkResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bulk response: %w", err)
	}
	if bulkResp.Errors {
		return fmt.Errorf("bulk request had errors: %s", string(b))
	}
	return err
}

func (c *Client) getValueForElasticsearch(col arrow.Array, i int) any {
	if col.IsNull(i) {
		return nil
	}
	switch col := col.(type) {
	case *array.Struct:
		m := map[string]any{}
		fields := col.DataType().(*arrow.StructType).Fields()
		for f, field := range fields {
			m[field.Name] = c.getValueForElasticsearch(col.Field(f), i)
		}
		return m
	case *cqtypes.JSONArray:
		return col.ValueStr(i)
	case array.ListLike:
		elems := make([]any, 0, col.Len())
		for j := 0; j < col.Len(); j++ {
			from, to := col.ValueOffsets(j)
			slc := array.NewSlice(col.ListValues(), from, to)
			for k := 0; k < slc.Len(); k++ {
				if slc.IsNull(k) {
					elems = append(elems, nil)
				} else {
					elems = append(elems, c.getValueForElasticsearch(slc, k))
				}
			}
		}
		return elems
	case *array.Timestamp:
		u := col.DataType().(*arrow.TimestampType).Unit
		switch u {
		case arrow.Nanosecond:
			return col.Value(i).ToTime(u).Format("2006-01-02T15:04:05.99999999Z")
		case arrow.Microsecond:
			return col.Value(i).ToTime(u).Format("2006-01-02T15:04:05.999999Z")
		case arrow.Millisecond:
			return col.Value(i).ToTime(u).Format("2006-01-02T15:04:05.999Z")
		case arrow.Second:
			return col.Value(i).ToTime(u).Format("2006-01-02T15:04:05Z")
		}
		panic(fmt.Sprintf("unsupported timestamp unit: %s", u))
	case *array.Time32:
		u := col.DataType().(*arrow.Time32Type).Unit
		switch u {
		case arrow.Second:
			format := "15:04:05"
			return padRightWithZero(col.Value(i).ToTime(u).Format(format), len(format))
		case arrow.Millisecond:
			format := "15:04:05.999"
			return padRightWithZero(col.Value(i).ToTime(u).Format(format), len(format))
		}
		panic(fmt.Sprintf("unsupported time32 unit: %s", u))
	case *array.Time64:
		u := col.DataType().(*arrow.Time64Type).Unit
		switch u {
		case arrow.Microsecond:
			format := "15:04:05.999999"
			return padRightWithZero(col.Value(i).ToTime(u).Format(format), len(format))
		case arrow.Nanosecond:
			format := "15:04:05.999999999"
			return padRightWithZero(col.Value(i).ToTime(u).Format(format), len(format))
		}
		panic(fmt.Sprintf("unsupported time64 unit: %s", u))
	}
	return col.GetOneForMarshal(i)
}

func pkIndexes(table *schema.Table) []int {
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// if no PK is defined, use all columns for the ID which is based on the indices returned by this function
		pks = table.Columns.Names()
	}
	inds := make([]int, 0, len(pks))
	for _, col := range pks {
		inds = append(inds, table.Columns.Index(col))
	}
	return inds
}

// elasticsearch IDs are limited to 512 bytes, so we hash the resource PK to make sure it's within the limit
func resourceID(record arrow.Record, i int, pkIndexes []int) uint64 {
	parts := make([]string, 0, len(pkIndexes))
	for _, pkIndex := range pkIndexes {
		parts = append(parts, record.Column(pkIndex).ValueStr(i))
	}
	h1 := fnv1a.HashString64(strings.Join(parts, "-"))
	return h1
}

func padRightWithZero(s string, length int) string {
	count := length - len(s)
	if count <= 0 {
		return s
	}
	return s + strings.Repeat("0", count)
}
