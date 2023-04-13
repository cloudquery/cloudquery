package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func CustomVerificationEmailTemplates() *schema.Table {
	tableName := "aws_ses_custom_verification_email_templates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetCustomVerificationEmailTemplate.html`,
		Resolver:            fetchSesCustomVerificationEmailTemplates,
		PreResourceResolver: getCustomVerificationEmailTemplate,
		Transform: transformers.TransformWithStruct(
			&sesv2.GetCustomVerificationEmailTemplateOutput{},
			transformers.WithSkipFields("ResultMetadata"),
			transformers.WithNameTransformer(client.CreateTrimPrefixTransformer("template_")),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCustomVerificationEmailTemplateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchSesCustomVerificationEmailTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	p := sesv2.NewListCustomVerificationEmailTemplatesPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.CustomVerificationEmailTemplates
	}

	return nil
}

func getCustomVerificationEmailTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2
	name := resource.Item.(types.CustomVerificationEmailTemplateMetadata).TemplateName

	getOutput, err := svc.GetCustomVerificationEmailTemplate(ctx, &sesv2.GetCustomVerificationEmailTemplateInput{TemplateName: name})
	if err != nil {
		return err
	}

	resource.SetItem(getOutput)

	return nil
}

func resolveCustomVerificationEmailTemplateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"custom-verification-email-template", *resource.Item.(*sesv2.GetCustomVerificationEmailTemplateOutput).TemplateName}, nil
	})(ctx, meta, resource, c)
}
