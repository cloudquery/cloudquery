package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkspacesResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "workspaces",
			Struct: &types.Workspace{},
			Multiplex: `client.ServiceAccountRegionMultiplexer("workspaces")`,
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveWorkspaceArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
		},
		{
			SubService: "directories",
			Struct: &types.WorkspaceDirectory{},
			Multiplex: `client.ServiceAccountRegionMultiplexer("workspaces")`,
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveDirectoryArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
		},
	}

	for _, r := range resources {
		r.Service = "workspaces"
	}
	return resources
}