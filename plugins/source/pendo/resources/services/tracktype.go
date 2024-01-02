package services

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	"github.com/cloudquery/pendo/client"
	"github.com/cloudquery/pendo/pendo"
)

func TrackTypesTable() *schema.Table {
	return &schema.Table{
		Name:      "tracktypes",
		Resolver:  fetchTrackTypes,
		Transform: transformers.TransformWithStruct(&pendo.TrackType{}),
	}
}

func fetchTrackTypes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c, ok := meta.(*client.Client)
	if !ok {
		return fmt.Errorf("invalid client meta: %v", meta)
	}

	trackTypes, err := c.Pendo.GetTrackTypes(ctx)
	if err != nil {
		return fmt.Errorf("failed to get trackTypes: %w", err)
	}

	for _, trackType := range trackTypes {
		res <- trackType
	}

	return nil
}
