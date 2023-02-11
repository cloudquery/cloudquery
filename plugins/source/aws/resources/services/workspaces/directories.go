package workspaces

import (
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Directories() *schema.Table {
	return &schema.Table{
		Name:        "aws_workspaces_directories",
		Description: `https://docs.aws.amazon.com/workspaces/latest/api/API_WorkspaceDirectory.html`,
		Resolver:    fetchWorkspacesDirectories,
		Transform:   transformers.TransformWithStruct(&types.WorkspaceDirectory{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("workspaces"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveDirectoryArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
