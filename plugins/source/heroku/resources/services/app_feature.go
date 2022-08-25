// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func AppFeatures() *schema.Table {
	return &schema.Table{
		Name:      "heroku_app_features",
		Resolver:  fetchAppFeatures,
		Multiplex: client.NoMultiplex,
		Columns: []schema.Column{
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "doc_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocURL"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "feedback_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FeedbackEmail"),
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
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchAppFeatures(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	items, err := c.Heroku.AppList(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, it := range items {
		v, err := c.Heroku.AppFeatureList(ctx, it.ID, nil)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
