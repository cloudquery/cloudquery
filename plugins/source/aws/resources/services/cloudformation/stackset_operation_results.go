package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func stackSetOperationResults() *schema.Table {
	table_name := "aws_cloudformation_stack_set_operation_results"
	return &schema.Table{
		Name: table_name,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html

The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:  fetchCloudformationStackSetOperationResults,
		Transform: transformers.TransformWithStruct(&types.StackSetOperationResultSummary{}, transformers.WithPrimaryKeyComponents("Account", "Region")),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "stack_set_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("stack_set_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "operation_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("operation_id"),
				PrimaryKeyComponent: true,
			},
		},
	}
}
func fetchCloudformationStackSetOperationResults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stackSet := parent.Parent.Item.(models.ExpandedStackSet)
	operation := parent.Item.(models.ExpandedStackSetOperation)
	input := cloudformation.ListStackSetOperationResultsInput{
		OperationId:  operation.OperationId,
		StackSetName: stackSet.StackSetId,
		CallAs:       stackSet.CallAs,
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudformation).Cloudformation
	paginator := cloudformation.NewListStackSetOperationResultsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}
