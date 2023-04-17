package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func StackSets() *schema.Table {
	table_name := "aws_cloudformation_stack_sets"
	return &schema.Table{
		Name:                table_name,
		Description:         `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSet.html`,
		Resolver:            fetchCloudformationStackSets,
		PreResourceResolver: getStackSet,
		Multiplex:           client.ServiceAccountRegionMultiplexer(table_name, "cloudformation"),
		Transform:           transformers.TransformWithStruct(&models.ExpandedStackSet{}, transformers.WithUnwrapStructFields("StackSet"), transformers.WithSkipFields("CallAs")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackSetId"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackSetARN"),
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
			stackSetOperations(),
		},
	}
}

func fetchCloudformationStackSets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	var err error
	var page *cloudformation.ListStackSetsOutput
	// There is no way of determining if an account is a delegated admin or not. So just need to test it out and fail over to the other one
	for _, callAs := range []types.CallAs{types.CallAsDelegatedAdmin, types.CallAsSelf} {
		config := cloudformation.ListStackSetsInput{
			CallAs: callAs,
		}
		paginator := cloudformation.NewListStackSetsPaginator(svc, &config)
		for paginator.HasMorePages() {
			page, err = paginator.NextPage(ctx)
			if err != nil {
				c.Logger().Info().Err(err).Msgf("failed to list stack sets with callAs: %s", string(callAs))
				break
			}
			for _, summary := range page.Summaries {
				res <- models.ExpandedSummary{
					StackSetSummary: summary,
					CallAs:          callAs,
				}
			}
		}
		if err == nil {
			return nil
		}
	}
	return err
}

func getStackSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	stack := resource.Item.(models.ExpandedSummary)
	var err error
	var stackSet *cloudformation.DescribeStackSetOutput

	input := &cloudformation.DescribeStackSetInput{
		StackSetName: stack.StackSetName,
		CallAs:       stack.CallAs,
	}

	stackSet, err = meta.(*client.Client).Services().Cloudformation.DescribeStackSet(ctx, input)
	if err != nil {
		return err
	}
	resource.Item = models.ExpandedStackSet{
		StackSet: *stackSet.StackSet,
		CallAs:   stack.CallAs,
	}
	return nil
}
