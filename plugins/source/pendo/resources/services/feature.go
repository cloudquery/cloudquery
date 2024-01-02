package services

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	"github.com/cloudquery/pendo/client"
	"github.com/cloudquery/pendo/pendo"
)

func FeaturesTable() *schema.Table {
	return &schema.Table{
		Name:      "features",
		Resolver:  fetchFeatures,
		Transform: transformers.TransformWithStruct(&pendo.Feature{}),
	}
}

func fetchFeatures(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c, ok := meta.(*client.Client)
	if !ok {
		return fmt.Errorf("invalid client meta: %v", meta)
	}

	features, err := c.Pendo.GetFeatures(ctx)
	if err != nil {
		return fmt.Errorf("failed to get features: %w", err)
	}

	for _, feature := range features {
		res <- feature
	}

	return nil
}
