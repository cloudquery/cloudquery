package client

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/samber/lo"
)

func sliceToSet(slice []string) map[string]struct{} {
	return lo.Associate(slice, func(s string) (string, struct{}) {
		return s, struct{}{}
	})
}

func extractCqIDs(msgs message.WriteInserts) ([]string, error) {
	cqIDs := make([]string, 0)
	for _, msg := range msgs {
		batchCQIDs, err := getCQIDs(msg.Record)
		if err != nil {
			return nil, err
		}
		cqIDs = append(cqIDs, batchCQIDs...)
	}
	return cqIDs, nil
}

func getCQIDs(r arrow.RecordBatch) ([]string, error) {
	cqIDColIndex, err := getColumnIndexWithName(r, CQIDColumn)
	if err != nil {
		return nil, err
	}
	results := make([]string, r.NumRows())
	for i := range int(r.NumRows()) {
		results[i] = r.Column(cqIDColIndex).ValueStr(i)
	}
	return results, nil
}

func getColumnIndexWithName(r arrow.RecordBatch, name string) (int, error) {
	for i := range int(r.NumCols()) {
		if r.Schema().Field(i).Name == name {
			return i, nil
		}
	}
	return -1, fmt.Errorf("column %s not found", name) // should never happen
}
