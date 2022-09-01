// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func PipelineDeployments() *schema.Table {
	return &schema.Table{
		Name:        "heroku_pipeline_deployments",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#pipeline-deployment-attributes",
		Resolver:    fetchPipelineDeployments,
		Columns: []schema.Column{
			{
				Name:     "addon_plan_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AddonPlanNames"),
			},
			{
				Name:     "app",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("App"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "current",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Current"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "output_stream_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutputStreamURL"),
			},
			{
				Name:     "slug",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Slug"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}

func fetchPipelineDeployments(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.Pipeline, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.PipelineList(ctxWithRange, nextRange)
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
			v, err := c.Heroku.PipelineDeploymentList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
