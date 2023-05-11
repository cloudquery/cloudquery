package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func stackTemplates() *schema.Table {
	tableName := "aws_cloudformation_stack_templates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplate.html`,
		Resolver:    fetchCloudformationStackTemplates,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform:   transformers.TransformWithStruct(&cloudformation.GetTemplateOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "stack_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "template_body",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TemplateBody"),
			},
		},
	}
}

func fetchCloudformationStackTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.GetTemplateInput{
		StackName: stack.StackName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	resp, err := svc.GetTemplate(ctx, &config, func(options *cloudformation.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- resp
	return nil
}
