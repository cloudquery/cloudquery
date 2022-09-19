package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSesTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SES

	listInput := new(sesv2.ListEmailTemplatesInput)
	for {
		output, err := svc.ListEmailTemplates(ctx, listInput, func(o *sesv2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}

		for _, templateMeta := range output.TemplatesMetadata {
			getInput := &sesv2.GetEmailTemplateInput{TemplateName: templateMeta.TemplateName}
			getOutput, err := svc.GetEmailTemplate(ctx, getInput, func(o *sesv2.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}

			res <- &Template{
				TemplateName:     getOutput.TemplateName,
				Text:             getOutput.TemplateContent.Text,
				Html:             getOutput.TemplateContent.Html,
				Subject:          getOutput.TemplateContent.Subject,
				CreatedTimestamp: templateMeta.CreatedTimestamp,
			}
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		listInput.NextToken = output.NextToken
	}

	return nil
}
func resolveSesTemplateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"template", *resource.Item.(*Template).TemplateName}, nil
	})(ctx, meta, resource, c)
}
