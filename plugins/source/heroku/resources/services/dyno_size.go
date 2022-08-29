// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func DynoSizes() *schema.Table {
	return &schema.Table{
		Name:        "heroku_dyno_sizes",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#dyno-size-attributes",
		Resolver:    fetchDynoSizes,
		Multiplex:   client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "compute",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Compute"),
			},
			{
				Name:     "cost",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cost"),
			},
			{
				Name:     "dedicated",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Dedicated"),
			},
			{
				Name:     "dyno_units",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DynoUnits"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "memory",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Memory"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "private_space_only",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PrivateSpaceOnly"),
			},
		},
	}
}

func fetchDynoSizes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.DynoSizeList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
