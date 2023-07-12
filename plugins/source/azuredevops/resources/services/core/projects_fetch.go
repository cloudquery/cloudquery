package core

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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
