// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func Pipelines() *schema.Table {
	return &schema.Table{
		Name:      "heroku_pipelines",
		Resolver:  fetchPipelines,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Owner"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchPipelines(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	v, err := c.Heroku.PipelineList(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- v
	return nil
}
