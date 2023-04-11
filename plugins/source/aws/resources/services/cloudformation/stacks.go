package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Stacks() *schema.Table {
	tableName := "aws_cloudformation_stacks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Stack.html`,
		Resolver:    fetchCloudformationStacks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform:   transformers.TransformWithStruct(&types.Stack{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackId"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			stackResources(),
		},
	}
}

func fetchCloudformationStacks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config cloudformation.DescribeStacksInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	paginator := cloudformation.NewDescribeStacksPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Stacks
	}
	return nil
}
func fetchCloudformationStackResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.ListStackResourcesInput{
		StackName: stack.StackName,
	}
	svc := meta.(*client.Client).Services().Cloudformation
	paginator := cloudformation.NewListStackResourcesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.StackResourceSummaries
	}
	return nil
}
