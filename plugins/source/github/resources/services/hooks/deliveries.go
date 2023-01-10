package hooks

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Deliveries() *schema.Table {
	return &schema.Table{
		Name:      "github_hook_deliveries",
		Resolver:  fetchDeliveries,
		Transform: transformers.TransformWithStruct(&github.HookDelivery{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "hook_id",
				Type:        schema.TypeInt,
				Resolver:    client.ResolveParentColumn("ID"),
				Description: `Hook ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "request",
				Type:     schema.TypeString,
				Resolver: resolveRequest,
			},
			{
				Name:     "response",
				Type:     schema.TypeString,
				Resolver: resolveResponse,
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
