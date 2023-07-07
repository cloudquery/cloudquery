package cloudformation

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Stacks() *schema.Table {
	tableName := "aws_cloudformation_stacks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Stack.html`,
		Resolver:    fetchCloudformationStacks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform: transformers.TransformWithStruct(
			&types.Stack{},
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"ar_ns": "arns"})),
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("StackId"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			stackResources(),
			templateSummaries(),
			stackTemplates(),
		},
	}
}

func fetchCloudformationStacks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config cloudformation.DescribeStacksInput
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudformation
	paginator := cloudformation.NewDescribeStacksPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Stacks
	}
	return nil
}
