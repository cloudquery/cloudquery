package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/goccy/go-json"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.RecordBatch) error {
	index := c.getIndexNamePattern(table)

	// refresh index before read, to ensure all written data is available
	_, err := c.typedClient.Indices.Refresh().Index(index).Do(ctx)
	if err != nil {
		var esErr *types.ElasticsearchError
		if errors.As(err, &esErr) && esErr.Status == 404 {
			// index doesn't exist yet, so there are no records to read
			return nil
		}
		return fmt.Errorf("failed to refresh index before read: %w", err)
	}

	size := 100
	searchResp, err := c.typedClient.Search().Index(index).Request(&search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
		Size: &size,
	}).Do(ctx)
	if err != nil {
		var esErr *types.ElasticsearchError
		if errors.As(err, &esErr) && esErr.Status == 404 {
			return nil
		}
		return fmt.Errorf("failed to read: %w", err)
	}

	sm := table.ToArrowSchema()
	for _, hit := range searchResp.Hits.Hits {
		var source map[string]any
		if err := json.Unmarshal(hit.Source_, &source); err != nil {
			return fmt.Errorf("failed to decode hit source: %w", err)
		}
		rb := array.NewRecordBuilder(memory.DefaultAllocator, sm)
		for i, field := range rb.Fields() {
			err := appendValue(field, source[sm.Field(i).Name])
			if err != nil {
				return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
			}
		}
		res <- rb.NewRecordBatch()
	}
	return nil
}

func appendValue(builder array.Builder, value any) error {
	if value == nil {
		builder.AppendNull()
		return nil
	}
	switch bldr := builder.(type) {
	case array.ListLikeBuilder:
		lst := value.([]any)
		if lst == nil {
			bldr.AppendNull()
			return nil
		}
		bldr.Append(true)
		valBuilder := bldr.ValueBuilder()
		for _, v := range lst {
			if err := appendValue(valBuilder, v); err != nil {
				return err
			}
		}
		return nil
	case *array.StructBuilder:
		m := value.(map[string]any)
		bldr.Append(true)
		bldrType := bldr.Type().(*arrow.StructType)
		for k, v := range m {
			idx, _ := bldrType.FieldIdx(k)
			fieldBldr := bldr.FieldBuilder(idx)
			if err := appendValue(fieldBldr, v); err != nil {
				return err
			}
		}
		return nil
	case *array.MonthIntervalBuilder, *array.DayTimeIntervalBuilder, *array.MonthDayNanoIntervalBuilder:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		dec := json.NewDecoder(bytes.NewReader(b))
		return bldr.UnmarshalOne(dec)
	case *array.Int8Builder, *array.Int16Builder, *array.Int32Builder, *array.Int64Builder:
		return bldr.AppendValueFromString(fmt.Sprintf("%d", int64(value.(float64))))
	case *array.Uint8Builder, *array.Uint16Builder, *array.Uint32Builder, *array.Uint64Builder:
		return bldr.AppendValueFromString(fmt.Sprintf("%d", uint64(value.(float64))))
	}
	return builder.AppendValueFromString(fmt.Sprintf("%v", value))
}
