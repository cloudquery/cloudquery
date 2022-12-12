package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProjectResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &vercel.Project{},
			Service:    "project",
			Multiplex:  "client.TeamMultiplex",
			PKColumns:  []string{"id"},
			Relations:  []string{"ProjectEnvs()"},
		},
		{
			DataStruct: &vercel.ProjectEnv{},
			Service:    "project",
			Multiplex:  "client.TeamMultiplex",
			PKColumns:  []string{"project_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "project_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
			},
		},
	}
}
