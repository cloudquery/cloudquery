package client

import (
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.projects))
	for i, projectId := range client.projects {
		l[i] = client.withProject(projectId)
	}
	return l
}

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	dupes := map[string]struct{}{}

	l := make([]schema.ClientMeta, 0, len(client.projects))
	for _, projectId := range client.projects {
		resp, err := client.Services.ResourcemanagerProjectsClient.GetProject(client.ctx, &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + projectId,
		})
		if err != nil {
			client.Logger().Warn().Err(err).Str("project_id", projectId).Msg("OrgMultiplex: Failed to get project info")
			continue
		}

		// Each parent gets a single client

		if _, ok := dupes[resp.Parent]; ok {
			client.Logger().Debug().Str("project_id", projectId).Str("parent", resp.Parent).Msg("OrgMultiplex: Skipping project")
			continue
		}

		client.Logger().Debug().Str("project_id", projectId).Str("parent", resp.Parent).Msg("OrgMultiplex: Found project")
		dupes[resp.Parent] = struct{}{}
		l = append(l, client.withProject(projectId))
	}
	return l
}
