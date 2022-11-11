package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSesTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	listInput := new(sesv2.ListEmailTemplatesInput)
	for {
		output, err := svc.ListEmailTemplates(ctx, listInput)
		if err != nil {
			return err
		}

		res <- output.TemplatesMetadata

		if aws.ToString(output.NextToken) == "" {
			break
		}
		listInput.NextToken = output.NextToken
	}

	return nil
}

func getTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2
	templateMeta := resource.Item.(types.EmailTemplateMetadata)

	getOutput, err := svc.GetEmailTemplate(ctx, &sesv2.GetEmailTemplateInput{TemplateName: templateMeta.TemplateName})
	if err != nil {
		return err
	}

	resource.Item = &models.Template{
		TemplateName:     getOutput.TemplateName,
		Text:             getOutput.TemplateContent.Text,
		Html:             getOutput.TemplateContent.Html,
		Subject:          getOutput.TemplateContent.Subject,
		CreatedTimestamp: templateMeta.CreatedTimestamp,
	}
	return nil
}

func resolveSesTemplateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"template", *resource.Item.(*models.Template).TemplateName}, nil
	})(ctx, meta, resource, c)
}
