package workspaces

import (
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Directories() *schema.Table {
	tableName := "aws_workspaces_directories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/workspaces/latest/api/API_WorkspaceDirectory.html`,
		Resolver:    fetchWorkspacesDirectories,
		Transform:   client.TransformWithStruct(&types.WorkspaceDirectory{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "workspaces"),
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
