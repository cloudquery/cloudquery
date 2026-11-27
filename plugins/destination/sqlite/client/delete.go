package client

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/v2/typeconv"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) DeleteRecord(ctx context.Context, messages message.WriteDeleteRecords) error {
	for _, msg := range messages {
		sql, err := generateDelete(msg.DeleteRecord)
		if err != nil {
			return err
		}

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

var groupingOperators = map[string]string{
	"AND": "AND",
	"OR":  "OR",
}

func groupingOperator(groupingType string) (string, error) {
	operator, ok := groupingOperators[strings.ToUpper(groupingType)]
	if !ok {
		return "", fmt.Errorf("unsupported grouping type %q", groupingType)
	}
	return operator, nil
}

func generateDelete(msg message.DeleteRecord) (string, error) {
	var sb strings.Builder

	sb.WriteString("DELETE FROM ")
	sb.WriteString(identifier(msg.TableName))
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
					operator, err := groupingOperator(predicateGroup.GroupingType)
					if err != nil {
						return "", err
					}
					sb.WriteString(" ")
					sb.WriteString(operator)
					sb.WriteString(" ")
				}
				sb.WriteString(identifier(predicate.Column))
				sb.WriteString(" = ?")
			}
			sb.WriteString(")")
			if i < len(msg.WhereClause)-1 {
				sb.WriteString(" AND ")
			}
		}
	}

	return sb.String(), nil
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
