package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func stackSetOperations() *schema.Table {
	return &schema.Table{
		Name:                "aws_cloudformation_stack_set_operations",
		Description:         `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperation.html`,
		Resolver:            fetchCloudformationStackSetOperations,
		PreResourceResolver: getStackSetOperation,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cloudformation"),
		Transform:           transformers.TransformWithStruct(&types.StackSetOperation{}, transformers.WithPrimaryKeys("StackSetId", "OperationId", "CreationTimestamp")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationId"),
			},
			{
				Name:     "stack_set_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("stack_set_arn"),
			},
		},

		Relations: []*schema.Table{
			stackSetOperationResults(),
		},
	}
}
func fetchCloudformationStackSetOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(*types.StackSet)
	config := cloudformation.ListStackSetOperationsInput{
		StackSetName: stack.StackSetName,
	}
	svc := meta.(*client.Client).Services().Cloudformation
	paginator := cloudformation.NewListStackSetOperationsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}

func getStackSetOperation(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	stack := resource.Parent.Item.(*types.StackSet)
	operation := resource.Item.(types.StackSetOperationSummary)
	stackSetOperation, err := meta.(*client.Client).Services().Cloudformation.DescribeStackSetOperation(ctx, &cloudformation.DescribeStackSetOperationInput{
		StackSetName: stack.StackSetName,
		OperationId:  operation.OperationId,
	})
	if err != nil {
		return err
	}
	resource.Item = stackSetOperation.StackSetOperation
	return nil
}
