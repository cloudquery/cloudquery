package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/goccy/go-json"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	index := c.getIndexNamePattern(table)

	// refresh index before read, to ensure all written data is available
	resp, err := c.typedClient.Indices.Refresh().Index(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to refresh index before read: %w", err)
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()

	size := 100
	resp, err = c.typedClient.Search().Index(index).Request(&search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
		Size: &size,
	}).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to read: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		if resp.StatusCode == 404 {
			return nil
		}
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to read: %s: %s", resp.Status, string(data))
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source map[string]any `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}

	sm := table.ToArrowSchema()
	for _, hit := range result.Hits.Hits {
		rb := array.NewRecordBuilder(memory.DefaultAllocator, sm)
		for i, field := range rb.Fields() {
			err := appendValue(field, hit.Source[sm.Field(i).Name])
			if err != nil {
				return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
			}
		}
		res <- rb.NewRecord()
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
