package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
)

func (c *Client) Read(_ context.Context, sc *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	index, err := c.Meilisearch.GetIndex(schema.TableName(sc))
	if err != nil {
		return err
	}

	req := &meilisearch.SearchRequest{
		Filter:      schema.CqSourceNameColumn.Name + " = '" + sourceName + "'",
		Sort:        []string{schema.CqSyncTimeColumn.Name + ":asc"},
		HitsPerPage: 100, // default = 1, we want more
		Page:        1,   // starting from 1
	}

	for {
		resp, err := index.Search("", req)
		if err != nil {
			return err
		}

		for _, hit := range resp.Hits {
			m, ok := hit.(map[string]any)
			if !ok {
				return fmt.Errorf("unsupported format for doc: %T", hit)
			}
			row, err := docToRecord(sc, m)
			if err != nil {
				return err
			}
			res <- row
		}

		if resp.TotalPages == req.Page {
			break
		}
		req.Page++
	}

	return nil
}

func docToRecord(sc *arrow.Schema, doc map[string]any) (arrow.Record, error) {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	defer builder.Release()

	for i, fld := range builder.Fields() {
		val, ok := doc[sc.Field(i).Name]
		if !ok || val == nil {
			fld.AppendNull()
			continue
		}

		switch tBuilder := fld.(type) {
		case *array.BooleanBuilder:
			tBuilder.Append(val.(bool))
		case *array.Int8Builder:
			tBuilder.Append(int8(val.(float64)))
		case *array.Int16Builder:
			tBuilder.Append(int16(val.(float64)))
		case *array.Int32Builder:
			tBuilder.Append(int32(val.(float64)))
		case *array.Int64Builder:
			tBuilder.Append(int64(val.(float64)))
		case *array.Uint8Builder:
			tBuilder.Append(uint8(val.(float64)))
		case *array.Uint16Builder:
			tBuilder.Append(uint16(val.(float64)))
		case *array.Uint32Builder:
			tBuilder.Append(uint32(val.(float64)))
		case *array.Uint64Builder:
			tBuilder.Append(uint64(val.(float64)))
		case *array.Float32Builder:
			tBuilder.Append(float32(val.(float64)))
		case *array.Float64Builder:
			tBuilder.Append(val.(float64))
		case *array.StringBuilder:
			tBuilder.Append(val.(string))
		case *array.BinaryBuilder:
			var data []byte
			val := val.(string)
			if val != "null" {
				err := json.Unmarshal([]byte(strconv.Quote(val)), &data)
				if err != nil {
					return nil, err
				}
			}
			tBuilder.Append(data)
		case *array.TimestampBuilder:
			ts, err := arrow.TimestampFromString(val.(string), arrow.Microsecond)
			if err != nil {
				return nil, err
			}
			tBuilder.Append(ts)
		case *types.InetBuilder:
			_, ipNet, err := net.ParseCIDR(val.(string))
			if err != nil {
				return nil, err
			}
			tBuilder.Append(*ipNet)
		case *types.JSONBuilder:
			tBuilder.Append(val)
		case *types.MacBuilder:
			mac, err := net.ParseMAC(val.(string))
			if err != nil {
				return nil, err
			}
			tBuilder.Append(mac)
		case *types.UUIDBuilder:
			uid, err := uuid.Parse(val.(string))
			if err != nil {
				return nil, err
			}
			tBuilder.Append(uid)
		case *array.ListBuilder:
			tBuilder.Append(true)
			switch listBuilder := tBuilder.ValueBuilder().(type) {
			case *array.StringBuilder:
				// string array
				for _, v := range val.([]any) {
					listBuilder.Append(v.(string))
				}
			case *array.Int64Builder:
				// int array
				for _, v := range val.([]any) {
					listBuilder.Append(int64(v.(float64)))
				}
			case *types.UUIDBuilder:
				// uuid array
				for _, v := range val.([]any) {
					uid, err := uuid.Parse(v.(string))
					if err != nil {
						return nil, err
					}
					listBuilder.Append(uid)
				}
			case *types.InetBuilder:
				// inet/cidr array
				for _, v := range val.([]any) {
					_, ipNet, err := net.ParseCIDR(v.(string))
					if err != nil {
						return nil, err
					}
					listBuilder.Append(*ipNet)
				}
			case *types.MacBuilder:
				// macaddr array
				for _, v := range val.([]any) {
					mac, err := net.ParseMAC(v.(string))
					if err != nil {
						return nil, err
					}
					listBuilder.Append(mac)
				}
			default:
				return nil, fmt.Errorf("unsupported list builder: %T", listBuilder)
			}
		default:
			return nil, fmt.Errorf("unsupported builder: %T", tBuilder)
		}
	}
	return builder.NewRecord(), nil
}
