package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWorkspacesWorkspaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Workspaces
	input := workspaces.DescribeWorkspacesInput{}
	for {
		output, err := svc.DescribeWorkspaces(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Workspaces
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func resolveWorkspaceArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Workspace)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "workspaces",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "workspaces/" + *item.WorkspaceId,
	}
	return resource.Set(c.Name, a.String())
}
