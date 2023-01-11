package actions

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:      "github_workflows",
		Resolver:  fetchWorkflows,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&Workflow{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "org",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrg,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "contents",
				Type:     schema.TypeString,
				Resolver: resolveContents,
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
