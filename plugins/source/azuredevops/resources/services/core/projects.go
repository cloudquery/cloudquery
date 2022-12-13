package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:     "azuredevops_core_projects",
		Resolver: fetchProjects,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("Id"),
			},
		},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	coreClient, err := core.NewClient(ctx, cl.Connection)
	if err != nil {
		return err
	}

	input := core.GetProjectsArgs{StateFilter: &core.ProjectStateValues.All}
	for {
		projects, err := coreClient.GetProjects(ctx, input)
		if err != nil {
			return err
		}

		res <- projects.Value

		if len(projects.ContinuationToken) == 0 {
			break
		}

		input.ContinuationToken = &projects.ContinuationToken
	}

	return nil
}
