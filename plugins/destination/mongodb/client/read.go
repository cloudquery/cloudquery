package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (*Client) createResultsArray(res bson.M, table *schema.Table) []any {
	results := make([]any, 0, len(table.Columns))
	for _, col := range table.Columns {
		val := res[col.Name]
		if val == nil {
			results = append(results, nil)
			continue
		}
		switch col.Type {
		case schema.TypeBool:
			r := (val).(bool)
			results = append(results, r)
		case schema.TypeInt:
			r := (val).(int64)
			results = append(results, r)
		case schema.TypeFloat:
			r := (val).(float64)
			results = append(results, r)
		case schema.TypeUUID:
			r := (val).(string)
			results = append(results, r)
		case schema.TypeString:
			r := (val).(string)
			results = append(results, r)
		case schema.TypeByteArray:
			r := (val).(primitive.Binary).Data
			results = append(results, r)
		case schema.TypeStringArray:
			r := make([]string, len((val).(primitive.A)))
			for i, v := range (val).(primitive.A) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeTimestamp:
			r := (val).(primitive.DateTime).Time().UTC()
			results = append(results, r)
		case schema.TypeJSON:
			r := (val).(primitive.M)
			results = append(results, r)
		case schema.TypeUUIDArray:
			r := make([]string, len((val).(primitive.A)))
			for i, v := range (val).(primitive.A) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeCIDR:
			r := (val).(string)
			results = append(results, r)
		case schema.TypeCIDRArray:
			r := make([]string, len((val).(primitive.A)))
			for i, v := range (val).(primitive.A) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeMacAddr:
			r := (val).(string)
			results = append(results, r)
		case schema.TypeMacAddrArray:
			r := make([]string, len((val).(primitive.A)))
			for i, v := range (val).(primitive.A) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeInet:
			r := (val).(string)
			results = append(results, r)
		case schema.TypeInetArray:
			r := make([]string, len((val).(primitive.A)))
			for i, v := range (val).(primitive.A) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeIntArray:
			r := make([]int64, len((val).(primitive.A)))
			for i, v := range (val).(primitive.A) {
				r[i] = v.(int64)
			}
			results = append(results, r)
		}
	}
	return results
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	cur, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Find(
		ctx,
		bson.M{"_cq_source_name": sourceName},
		options.Find().SetSort(bson.M{"_cq_sync_time": 1}))
	if err != nil {
		return fmt.Errorf("failed to read table %s: %w", table.Name, err)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		values := c.createResultsArray(result, table)
		res <- values
	}
	return nil
}
