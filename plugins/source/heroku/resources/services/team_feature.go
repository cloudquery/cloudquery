package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func TeamFeatures() *schema.Table {
	return &schema.Table{
		Name:        "heroku_team_features",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#team-feature`,
		Resolver:    fetchTeamFeatures,
		Transform:   transformers.TransformWithStruct(&heroku.TeamFeature{}),
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

func fetchTeamFeatures(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.Team, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange) // nolint:revive,staticcheck
		v, err := c.Heroku.TeamList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		items = append(items, v...)
	}

	for _, it := range items {
		nextRange = &heroku.ListRange{
			Field: "id",
			Max:   1000,
		}
		// Roundtripper middleware in client/pagination.go
		// sets the nextRange value after each request
		for nextRange.Max != 0 {
			ctxWithRange := context.WithValue(ctx, "nextRange", nextRange) // nolint:revive,staticcheck
			v, err := c.Heroku.TeamFeatureList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
