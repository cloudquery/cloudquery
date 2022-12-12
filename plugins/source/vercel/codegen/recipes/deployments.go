package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DeploymentResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &vercel.Deployment{},
			Service:    "deployment",
			PKColumns:  []string{"uid"},
			Relations:  []string{"DeploymentChecks()"},
		},
		{
			DataStruct: &vercel.DeploymentCheck{},
			Service:    "deployment",
			PKColumns:  []string{"deployment_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "deployment_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("uid")`,
				},
			},
		},
	}
}
