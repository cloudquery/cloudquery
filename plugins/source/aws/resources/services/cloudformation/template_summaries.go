package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func templateSummaries() *schema.Table {
	tableName := "aws_cloudformation_template_summaries"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html`,
		Resolver:    fetchTemplateSummary,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform: transformers.TransformWithStruct(
			&cloudformation.GetTemplateSummaryOutput{},
			transformers.WithSkipFields("ResultMetadata"), // This field contains metadata about the API call rather than the template itself, so remove it.
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "stack_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:            "stack_arn",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
		},
	}
}

func fetchTemplateSummary(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation

	stack := parent.Item.(types.Stack)

	summary, err := svc.GetTemplateSummary(ctx, &cloudformation.GetTemplateSummaryInput{
		StackName: stack.StackName,
	}, func(o *cloudformation.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}

	res <- summary
	return nil
}
