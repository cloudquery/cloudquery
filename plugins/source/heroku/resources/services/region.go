package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func Regions() *schema.Table {
	return &schema.Table{
		Name:        "heroku_regions",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#region`,
		Resolver:    fetchRegions,
		Transform:   transformers.TransformWithStruct(&heroku.Region{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchRegions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange) // nolint:revive,staticcheck
		v, err := c.Heroku.RegionList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
