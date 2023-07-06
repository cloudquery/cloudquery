package workspaces

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Directories() *schema.Table {
	tableName := "aws_workspaces_directories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/workspaces/latest/api/API_WorkspaceDirectory.html`,
		Resolver:    fetchWorkspacesDirectories,
		Transform:   transformers.TransformWithStruct(&types.WorkspaceDirectory{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "workspaces"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveDirectoryArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchWorkspacesDirectories(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Workspaces
	input := workspaces.DescribeWorkspaceDirectoriesInput{}
	paginator := workspaces.NewDescribeWorkspaceDirectoriesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *workspaces.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Directories
	}
	return nil
}

func resolveDirectoryArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.WorkspaceDirectory)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "workspaces",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "diretory/" + *item.DirectoryId,
	}
	return resource.Set(c.Name, a.String())
}
