package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	cqtypes "github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/segmentio/fasthash/fnv1a"
)

type bulkResponse struct {
	Took   int64 `json:"took"`
	Errors bool  `json:"errors"`
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	for _, record := range records {
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
	cqSyncName := table.Columns.Index(schema.CqSyncTimeColumn.Name)
	cqSyncUnit := schema.CqSyncTimeColumn.Type.(*arrow.TimestampType).Unit
	syncTime := record.Column(cqSyncName).(*array.Timestamp).Value(0).ToTime(cqSyncUnit)
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
		if c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale {
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
	index := c.getIndexName(table.Name, syncTime)
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
		arr := col.(array.ListLike)
		elems := make([]any, 0, arr.Len())
		for j := 0; j < arr.Len(); j++ {
			from, to := arr.ValueOffsets(j)
			slc := array.NewSlice(arr.ListValues(), from, to)
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
		return col.Value(i).ToTime(col.DataType().(*arrow.TimestampType).Unit).Format("2006-01-02T15:04:05.999Z")
	case *array.DayTimeInterval:
		return col.ValueStr(i)
	case *array.MonthInterval:
		return col.ValueStr(i)
	case *array.MonthDayNanoInterval:
		return col.ValueStr(i)
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
