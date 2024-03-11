package cmd

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type syncSummary struct {
	Resources uint64
	Errors    uint64
	Warnings  uint64
	SyncID    string
}

func generateSummaryTable() *schema.Table {
	tableName := "cloudquery_sync_summary"
	t := schema.Tables{{
		Name:      tableName,
		Transform: transformers.TransformWithStruct(&syncSummary{}),
		Columns:   []schema.Column{},
	}}

	if err := transformers.TransformTables(t); err != nil {
		panic(err)
	}

	return t[0]
}
