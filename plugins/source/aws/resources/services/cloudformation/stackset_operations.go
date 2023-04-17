package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func stackSetOperations() *schema.Table {
	table_name := "aws_cloudformation_stack_set_operations"
	return &schema.Table{
		Name:                table_name,
		Description:         `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperation.html`,
		Resolver:            fetchCloudformationStackSetOperations,
		PreResourceResolver: getStackSetOperation,
		Multiplex:           client.ServiceAccountRegionMultiplexer(table_name, "cloudformation"),
		Transform:           transformers.TransformWithStruct(&models.ExpandedStackSetOperation{}, transformers.WithUnwrapStructFields("StackSetOperation"), transformers.WithSkipFields("CallAs"), transformers.WithPrimaryKeys("OperationId", "CreationTimestamp")),
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
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			stackSetOperationResults(),
		},
	}
}
func fetchCloudformationStackSetOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(models.ExpandedStackSet)

	input := cloudformation.ListStackSetOperationsInput{
		StackSetName: stack.StackSetName,
		CallAs:       stack.CallAs,
	}

	svc := meta.(*client.Client).Services().Cloudformation
	paginator := cloudformation.NewListStackSetOperationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, operation := range page.Summaries {
			res <- models.ExpandedStackSetOperationSummary{
				StackSetOperationSummary: operation,
				CallAs:                   input.CallAs,
			}
		}
	}
	return nil
}

func getStackSetOperation(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	stack := resource.Parent.Item.(models.ExpandedStackSet)
	operation := resource.Item.(models.ExpandedStackSetOperationSummary)

	input := cloudformation.DescribeStackSetOperationInput{
		StackSetName: stack.StackSetName,
		OperationId:  operation.OperationId,
		CallAs:       stack.CallAs,
	}

	stackSetOperation, err := meta.(*client.Client).Services().Cloudformation.DescribeStackSetOperation(ctx, &input)
	if err != nil {
		return err
	}
	resource.Item = models.ExpandedStackSetOperation{
		StackSetOperation: *stackSetOperation.StackSetOperation,
		CallAs:            input.CallAs,
	}
	return nil
}
