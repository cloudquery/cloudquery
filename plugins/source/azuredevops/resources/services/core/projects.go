package core

import (
	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "azuredevops_core_projects",
		Resolver:  fetchProjects,
		Transform: client.TransformWithStruct(&core.TeamProjectReference{}),
	}
}
