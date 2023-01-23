package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	readSQL    = "SELECT * FROM `%s.%s.%s` WHERE `_cq_source_name` = @cq_source_name order by _cq_sync_time asc"
	readCypher = "MATCH (t:%s {_cq_source_name: $cq_source_name}) RETURN t ORDER BY t._cq_sync_time ASC"
)

func (*Client) createResultsArray(table *schema.Table, node *neo4j.Node) []any {
	results := make([]any, 0, len(table.Columns))
	for _, col := range table.Columns {
		if node.Props[col.Name] == nil {
			results = append(results, nil)
			continue
		}
		switch col.Type {
		case schema.TypeBool:
			r := node.Props[col.Name].(bool)
			results = append(results, &r)
		case schema.TypeInt:
			r := node.Props[col.Name].(int64)
			results = append(results, &r)
		case schema.TypeFloat:
			r := node.Props[col.Name].(float64)
			results = append(results, &r)
		case schema.TypeUUID:
			r := node.Props[col.Name].(string)
			results = append(results, &r)
		case schema.TypeString:
			r := node.Props[col.Name].(string)
			results = append(results, &r)
		case schema.TypeByteArray:
			r := node.Props[col.Name].([]byte)
			results = append(results, &r)
		case schema.TypeStringArray:
			r := make([]string, len(node.Props[col.Name].([]any)))
			for i, v := range node.Props[col.Name].([]any) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeTimestamp:
			r := node.Props[col.Name].(time.Time)
			results = append(results, &r)
		case schema.TypeJSON:
			r := node.Props[col.Name].(string)
			results = append(results, &r)
		case schema.TypeUUIDArray:
			r := make([]string, len(node.Props[col.Name].([]any)))
			for i, v := range node.Props[col.Name].([]any) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeCIDR:
			r := node.Props[col.Name].(string)
			results = append(results, &r)
		case schema.TypeCIDRArray:
			r := make([]string, len(node.Props[col.Name].([]any)))
			for i, v := range node.Props[col.Name].([]any) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeMacAddr:
			r := node.Props[col.Name].(string)
			results = append(results, &r)
		case schema.TypeMacAddrArray:
			r := make([]string, len(node.Props[col.Name].([]any)))
			for i, v := range node.Props[col.Name].([]any) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeInet:
			r := node.Props[col.Name].(string)
			results = append(results, &r)
		case schema.TypeInetArray:
			r := make([]string, len(node.Props[col.Name].([]any)))
			for i, v := range node.Props[col.Name].([]any) {
				r[i] = v.(string)
			}
			results = append(results, r)
		case schema.TypeIntArray:
			r := make([]int64, len(node.Props[col.Name].([]any)))
			for i, v := range node.Props[col.Name].([]any) {
				r[i] = v.(int64)
			}
			results = append(results, r)
		default:
			panic(fmt.Sprintf("unsupported type for col %v: %v", col.Name, col.Type))
		}
	}
	return results
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	stmt := fmt.Sprintf(readCypher, table.Name)

	session := c.client.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		r, err := tx.Run(ctx, stmt, map[string]any{"cq_source_name": sourceName})
		if err != nil {
			return nil, err
		}
		records, err := r.Collect(ctx)
		if err != nil {
			return nil, err
		}
		for _, record := range records {
			values := record.Values
			for _, value := range values {
				node := value.(neo4j.Node)
				result := c.createResultsArray(table, &node)
				res <- result
			}
		}
		return nil, nil
	})
	return nil
}
