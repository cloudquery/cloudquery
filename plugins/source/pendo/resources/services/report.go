package services

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	"github.com/cloudquery/pendo/client"
	"github.com/cloudquery/pendo/pendo"
)

func ReportsTable() *schema.Table {
	return &schema.Table{
		Name:      "reports",
		Resolver:  fetchReports,
		Transform: transformers.TransformWithStruct(&pendo.Report{}),
	}
}

func fetchReports(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c, ok := meta.(*client.Client)
	if !ok {
		return fmt.Errorf("invalid client meta: %v", meta)
	}

	reports, err := c.Pendo.GetReports(ctx)
	if err != nil {
		return fmt.Errorf("failed to get reports: %w", err)
	}

	for _, report := range reports {
		res <- report
	}

	return nil
}
