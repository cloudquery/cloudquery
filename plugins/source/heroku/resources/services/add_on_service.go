// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func AddOnServices() *schema.Table {
	return &schema.Table{
		Name:        "heroku_add_on_services",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#add-on-service`,
		Resolver:    fetchAddOnServices,
		Columns: []schema.Column{
			{
				Name:     "cli_plugin_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CliPluginName"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "human_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HumanName"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "supports_multiple_installations",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsMultipleInstallations"),
			},
			{
				Name:     "supports_sharing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportsSharing"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchAddOnServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AddOnServiceList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
