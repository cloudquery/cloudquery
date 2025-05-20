package client

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/typeconv/ch/values"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/util"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) DeleteRecord(ctx context.Context, messages message.WriteDeleteRecords) error {
	if len(messages) == 0 {
		return nil
	}

	for _, msg := range messages {
		sql := generateDelete(msg.DeleteRecord)
		params, err := extractPredicateValues(msg.DeleteRecord.WhereClause)
		if err != nil {
			return err
		}

		if err = c.conn.Exec(ctx, sql, params...); err != nil {
			return err
		}
	}

	return nil
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
			primitiveValues, err := values.FromArray(col)
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

func generateDelete(msg message.DeleteRecord) string {
	var sb strings.Builder

	sb.WriteString("DELETE FROM ")
	sb.WriteString(util.SanitizeID(msg.TableName))
	sb.WriteString(" WHERE ")
	if len(msg.WhereClause) == 0 {
		sb.WriteString("1")
	} else {
		counter := 1
		for i, predicateGroup := range msg.WhereClause {
			if len(predicateGroup.Predicates) == 0 {
				continue
			}
			sb.WriteString("(")
			for i, predicate := range predicateGroup.Predicates {
				if i > 0 {
					sb.WriteString(" ")
					sb.WriteString(predicateGroup.GroupingType)
					sb.WriteString(" ")
				}
				sb.WriteString(util.SanitizeID(predicate.Column))
				sb.WriteString(" = $")
				sb.WriteString(strconv.Itoa(counter))
				counter++
			}
			sb.WriteString(")")
			if i < len(msg.WhereClause)-1 {
				sb.WriteString(" AND ")
			}
		}
	}

	return sb.String()
}
