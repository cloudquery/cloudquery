// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func BuildpackInstallations() *schema.Table {
	return &schema.Table{
		Name:        "heroku_buildpack_installations",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#buildpack-installation-attributes",
		Resolver:    fetchBuildpackInstallations,
		Multiplex:   client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "buildpack",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Buildpack"),
			},
			{
				Name:     "ordinal",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Ordinal"),
			},
		},
	}
}

func fetchBuildpackInstallations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.App, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AppList(ctxWithRange, nextRange)
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
			ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
			v, err := c.Heroku.BuildpackInstallationList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
