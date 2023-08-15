package ses

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveCustomVerificationEmailTemplateArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchSesCustomVerificationEmailTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2

	p := sesv2.NewListCustomVerificationEmailTemplatesPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(o *sesv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.CustomVerificationEmailTemplates
	}

	return nil
}

func getCustomVerificationEmailTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2
	name := resource.Item.(types.CustomVerificationEmailTemplateMetadata).TemplateName

	getOutput, err := svc.GetCustomVerificationEmailTemplate(ctx,
		&sesv2.GetCustomVerificationEmailTemplateInput{TemplateName: name},
		func(o *sesv2.Options) {
			o.Region = cl.Region
		},
	)
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
