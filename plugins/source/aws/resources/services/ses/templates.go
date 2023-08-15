package ses

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Templates() *schema.Table {
	tableName := "aws_ses_templates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailTemplate.html`,
		Resolver:            fetchSesTemplates,
		PreResourceResolver: getTemplate,
		Transform:           transformers.TransformWithStruct(&models.Template{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveTemplateArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchSesTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2

	p := sesv2.NewListEmailTemplatesPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(o *sesv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.TemplatesMetadata
	}

	return nil
}

func getTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2
	templateMeta := resource.Item.(types.EmailTemplateMetadata)

	getOutput, err := svc.GetEmailTemplate(ctx,
		&sesv2.GetEmailTemplateInput{TemplateName: templateMeta.TemplateName},
		func(o *sesv2.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	resource.SetItem(&models.Template{
		TemplateName:     getOutput.TemplateName,
		Text:             getOutput.TemplateContent.Text,
		Html:             getOutput.TemplateContent.Html,
		Subject:          getOutput.TemplateContent.Subject,
		CreatedTimestamp: templateMeta.CreatedTimestamp,
	})

	return nil
}

func resolveTemplateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"template", *resource.Item.(*models.Template).TemplateName}, nil
	})(ctx, meta, resource, c)
}
