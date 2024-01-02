package services

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	"github.com/cloudquery/pendo/client"
	"github.com/cloudquery/pendo/pendo"
)

func PagesTable() *schema.Table {
	return &schema.Table{
		Name:      "pages",
		Resolver:  fetchPages,
		Transform: transformers.TransformWithStruct(&pendo.Page{}),
	}
}

func fetchPages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c, ok := meta.(*client.Client)
	if !ok {
		return fmt.Errorf("invalid client meta: %v", meta)
	}

	pages, err := c.Pendo.GetPages(ctx)
	if err != nil {
		return fmt.Errorf("failed to get pages: %w", err)
	}

	for _, page := range pages {
		res <- page
	}

	return nil
}
