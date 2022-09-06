// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func AddOnRegionCapabilities() *schema.Table {
	return &schema.Table{
		Name:        "heroku_add_on_region_capabilities",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#add-on-region-capability-attributes",
		Resolver:    fetchAddOnRegionCapabilities,
		Columns: []schema.Column{
			{
				Name:     "addon_service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddonService"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "region",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "supports_private_networking",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsPrivateNetworking"),
			},
		},
	}
}

func fetchAddOnRegionCapabilities(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AddOnRegionCapabilityList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
