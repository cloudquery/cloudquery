package cloudformation

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func stackSetOperationResults() *schema.Table {
	return &schema.Table{
		Name: "aws_cloudformation_stack_set_operation_results",
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html.
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:  fetchCloudformationStackSetOperationResults,
		Multiplex: client.ServiceAccountRegionMultiplexer("cloudformation"),
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
