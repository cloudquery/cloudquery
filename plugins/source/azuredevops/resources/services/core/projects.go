package core

import (
	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "azuredevops_core_projects",
		Resolver:  fetchProjects,
		Transform: transformers.TransformWithStruct(&core.TeamProjectReference{}, client.Options()...),
	}
}
