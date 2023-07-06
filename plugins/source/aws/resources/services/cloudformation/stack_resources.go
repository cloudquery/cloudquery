package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func stackResources() *schema.Table {
	tableName := "aws_cloudformation_stack_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html`,
		Resolver:    fetchCloudformationStackResources,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform:   transformers.TransformWithStruct(&types.StackResourceSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "stack_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchCloudformationStackResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.ListStackResourcesInput{
		StackName: stack.StackName,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudformation
	paginator := cloudformation.NewListStackResourcesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.StackResourceSummaries
	}
	return nil
}
