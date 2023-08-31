package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func stackInstanceSummaries() *schema.Table {
	table_name := "aws_cloudformation_stack_instance_summaries"
	return &schema.Table{
		Name:        table_name,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceSummary.html`,
		Resolver:    fetchStackInstanceSummary,
		Transform:   transformers.TransformWithStruct(&models.ExpandedStackInstanceSummary{}, transformers.WithUnwrapStructFields("StackInstanceSummary"), transformers.WithSkipFields("CallAs"), transformers.WithPrimaryKeys("StackSetId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("StackId"),
			},
			{
				Name:       "stack_set_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("stack_set_arn"),
				PrimaryKey: true,
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
		for i := range page.Summaries {
			res <- models.ExpandedStackInstanceSummary{
				StackInstanceSummary: page.Summaries[i],
				CallAs:               input.CallAs,
			}
		}
	}

	return nil
}
