package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func StackSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_cloudformation_stack_sets",
		Description:         `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSet.html`,
		Resolver:            fetchCloudformationStackSets,
		PreResourceResolver: getStackSet,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cloudformation"),
		Transform:           transformers.TransformWithStruct(&types.StackSet{}),
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
	config := cloudformation.ListStackSetsInput{}
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	paginator := cloudformation.NewListStackSetsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}

func getStackSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	stack := resource.Item.(types.StackSetSummary)
	stackSet, err := meta.(*client.Client).Services().Cloudformation.DescribeStackSet(ctx, &cloudformation.DescribeStackSetInput{
		StackSetName: stack.StackSetName,
	})
	if err != nil {
		return err
	}
	resource.Item = stackSet.StackSet
	return nil
}
