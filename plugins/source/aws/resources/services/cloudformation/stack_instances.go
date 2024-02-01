package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func stackInstanceSummaries() *schema.Table {
	table_name := "aws_cloudformation_stack_instance_summaries"
	return &schema.Table{
		Name: table_name,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceSummary.html

**Note**: Sometimes the stack instance ID may be unavailable in the API (i.e., the instance is in a bad state), so it will have value of ` + "`N/A`.",
		Resolver: fetchStackInstanceSummary,
		Transform: transformers.TransformWithStruct(&models.ExpandedStackInstanceSummary{},
			transformers.WithUnwrapStructFields("StackInstanceSummary"),
			transformers.WithSkipFields("CallAs"),
			transformers.WithPrimaryKeyComponents("StackSetId", "StackId"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("StackId"),
			},
			{
				Name:                "stack_set_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("stack_set_arn"),
				PrimaryKeyComponent: true,
			},
		},

		Relations: []*schema.Table{
			stackInstanceResourceDrifts(),
		},
	}
}
func fetchStackInstanceSummary(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(models.ExpandedStackSet)

	input := cloudformation.ListStackInstancesInput{
		StackSetName: stack.StackSetName,
		CallAs:       stack.CallAs,
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudformation).Cloudformation

	paginator := cloudformation.NewListStackInstancesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, summary := range page.Summaries {
			if summary.StackId == nil {
				// can happen if the instance is in a bad state
				summary.StackId = aws.String("N/A")
			}
			res <- models.ExpandedStackInstanceSummary{
				StackInstanceSummary: summary,
				CallAs:               input.CallAs,
			}
		}
	}

	return nil
}
