package services

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	"github.com/cloudquery/pendo/client"
	"github.com/cloudquery/pendo/pendo"
)

func GuidesTable() *schema.Table {
	return &schema.Table{
		Name:      "guides",
		Resolver:  fetchGuides,
		Transform: transformers.TransformWithStruct(&pendo.Guide{}),
	}
}

func fetchGuides(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c, ok := meta.(*client.Client)
	if !ok {
		return fmt.Errorf("invalid client meta: %v", meta)
	}

	guides, err := c.Pendo.GetGuides(ctx)
	if err != nil {
		return fmt.Errorf("failed to get guides: %w", err)
	}

	for _, guide := range guides {
		res <- guide
	}

	return nil
}
