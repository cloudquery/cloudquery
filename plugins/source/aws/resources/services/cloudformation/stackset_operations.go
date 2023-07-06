package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("OperationId"),
			},
			{
				Name:       "stack_set_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("stack_set_arn"),
				PrimaryKey: true,
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

	cl := meta.(*client.Client)
	svc := cl.Services().Cloudformation
	paginator := cloudformation.NewListStackSetOperationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
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
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudformation
	stack := resource.Parent.Item.(models.ExpandedStackSet)
	operation := resource.Item.(models.ExpandedStackSetOperationSummary)

	input := cloudformation.DescribeStackSetOperationInput{
		StackSetName: stack.StackSetName,
		OperationId:  operation.OperationId,
		CallAs:       stack.CallAs,
	}

	stackSetOperation, err := svc.DescribeStackSetOperation(ctx, &input, func(options *cloudformation.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = models.ExpandedStackSetOperation{
		StackSetOperation: *stackSetOperation.StackSetOperation,
		CallAs:            input.CallAs,
	}
	return nil
}
