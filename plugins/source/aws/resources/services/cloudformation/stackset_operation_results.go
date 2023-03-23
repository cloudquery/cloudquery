package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func stackSetOperationResults() *schema.Table {
	table_name := "aws_cloudformation_stack_set_operation_results"
	return &schema.Table{
		Name: table_name,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html.
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:  fetchCloudformationStackSetOperationResults,
		Multiplex: client.ServiceAccountRegionMultiplexer(table_name, "cloudformation"),
		Transform: transformers.TransformWithStruct(&types.StackSetOperationResultSummary{}),
		Columns: []schema.Column{
			{
				Name:     "request_account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "request_region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "operation_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("operation_id"),
			},
			{
				Name:     "stack_set_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("stack_set_arn"),
			},
		},
	}
}
func fetchCloudformationStackSetOperationResults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stackSet := parent.Parent.Item.(models.ExpandedStackSet)
	operation := parent.Item.(models.ExpandedStackSetOperation)
	input := cloudformation.ListStackSetOperationResultsInput{
		OperationId:  operation.OperationId,
		StackSetName: stackSet.StackSetName,
		CallAs:       stackSet.CallAs,
	}

	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	paginator := cloudformation.NewListStackSetOperationResultsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}
