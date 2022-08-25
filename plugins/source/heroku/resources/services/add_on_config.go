// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func AddOnConfigs() *schema.Table {
	return &schema.Table{
		Name:      "heroku_add_on_configs",
		Resolver:  fetchAddOnConfigs,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Value"),
			},
		},
	}
}

func fetchAddOnConfigs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	items, err := c.Heroku.AddOnList(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, it := range items {
		v, err := c.Heroku.AddOnConfigList(ctx, it.ID, nil)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
