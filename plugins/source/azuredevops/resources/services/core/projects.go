package core

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "azuredevops_core_projects",
		Resolver:  fetchProjects,
		Transform: transformers.TransformWithStruct(&core.TeamProjectReference{}, transformers.WithSkipFields("Id")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("Id"),
			},
		},
	}
}
