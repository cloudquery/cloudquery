package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func stackInstanceResourceDrifts() *schema.Table {
	table_name := "aws_cloudformation_stack_instance_resource_drifts"
	return &schema.Table{
		Name: table_name,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceResourceDriftsSummary.html.
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:  fetchStackInstanceResourceDrifts,
		Transform: transformers.TransformWithStruct(&types.StackInstanceResourceDriftsSummary{}, transformers.WithPrimaryKeyComponents("StackId", "LogicalResourceId", "PhysicalResourceId")),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(false),
			client.RequestRegionColumn(false),
			{
				Name:                "stack_set_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("stack_set_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "operation_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("last_operation_id"),
				PrimaryKeyComponent: true,
			},
		},
	}
}
func fetchStackInstanceResourceDrifts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance := parent.Item.(models.ExpandedStackInstanceSummary)
	if instance.DriftStatus == types.StackDriftStatusNotChecked {
		return nil
	}

	stackSet := parent.Parent.Item.(models.ExpandedStackSet)
	config := cloudformation.ListStackInstanceResourceDriftsInput{
		OperationId:          instance.LastOperationId,
		StackInstanceAccount: instance.Account,
		StackInstanceRegion:  instance.Region,
		StackSetName:         stackSet.StackSetName,
		CallAs:               stackSet.CallAs,
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudformation).Cloudformation

	// No paginator available
	for {
		output, err := svc.ListStackInstanceResourceDrifts(ctx, &config, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Summaries
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
