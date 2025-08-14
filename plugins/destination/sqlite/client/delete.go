package client

import (
	"context"
	"reflect"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/v2/typeconv"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) DeleteRecord(ctx context.Context, messages message.WriteDeleteRecords) error {
	for _, msg := range messages {
		sql := generateDelete(msg.DeleteRecord)

		params, err := extractPredicateValues(msg.DeleteRecord.WhereClause)
		if err != nil {
			return err
		}

		if _, err = c.db.ExecContext(ctx, sql, params...); err != nil {
			return err
		}
	}
	return nil
}

func generateDelete(msg message.DeleteRecord) string {
	var sb strings.Builder

	sb.WriteString("DELETE FROM ")
	sb.WriteString("\"" + msg.TableName + "\"")
	sb.WriteString(" WHERE ")
	if len(msg.WhereClause) == 0 {
		sb.WriteString("1")
	} else {
		for i, predicateGroup := range msg.WhereClause {
			if len(predicateGroup.Predicates) == 0 {
				continue
			}
			sb.WriteString("(")
			for j, predicate := range predicateGroup.Predicates {
				if j > 0 {
					sb.WriteString(" ")
					sb.WriteString(predicateGroup.GroupingType)
					sb.WriteString(" ")
				}
				sb.WriteString("\"" + predicate.Column + "\"")
				sb.WriteString(" = ?")
			}
			sb.WriteString(")")
			if i < len(msg.WhereClause)-1 {
				sb.WriteString(" AND ")
			}
		}
	}

	return sb.String()
}

func extractPredicateValues(where message.PredicateGroups) ([]any, error) {
	predicateCount := 0
	for _, predicateGroup := range where {
		predicateCount += len(predicateGroup.Predicates)
	}
	results := make([]any, predicateCount)
	counter := 0
	for _, predicateGroup := range where {
		for _, predicate := range predicateGroup.Predicates {
			col := predicate.Record.Column(0)
			primitiveValues, err := typeconv.FromArray(col)
			if err != nil {
				return nil, err
			}
			unpacked := unpackArray(primitiveValues)
			results[counter] = unpacked[0]
			counter++
		}
	}
	return results, nil
}

func unpackArray(s any) []any {
	v := reflect.ValueOf(s)
	r := make([]any, v.Len())
	for i := range v.Len() {
		r[i] = v.Index(i).Interface()
	}
	return r
}
