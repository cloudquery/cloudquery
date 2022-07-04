package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource templates --config gen.hcl --output .
func Templates() *schema.Table {
	return &schema.Table{
		Name:         "aws_ses_templates",
		Description:  "Amazon Simple Email Service (SES) is a cost-effective, flexible, and scalable email service that enables developers to send mail from within any application.",
		Resolver:     fetchSesTemplates,
		Multiplex:    client.ServiceAccountRegionMultiplexer("email"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver:    ResolveSesTemplateArn,
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "created_timestamp",
				Description: "The time and date the template was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the template",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Template.TemplateName"),
			},
			{
				Name:        "html_part",
				Description: "The HTML body of the email.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Template.HtmlPart"),
			},
			{
				Name:        "subject_part",
				Description: "The subject line of the email.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Template.SubjectPart"),
			},
			{
				Name:        "text_part",
				Description: "The email body that will be visible to recipients whose email clients do not display HTML.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Template.TextPart"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSesTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SES

	listInput := new(ses.ListTemplatesInput)
	for {
		output, err := svc.ListTemplates(ctx, listInput, func(o *ses.Options) { o.Region = c.Region })
		if err != nil {
			return diag.WrapError(err)
		}

		for _, templateMeta := range output.TemplatesMetadata {
			getInput := &ses.GetTemplateInput{TemplateName: templateMeta.Name}
			template, err := svc.GetTemplate(ctx, getInput, func(o *ses.Options) { o.Region = c.Region })
			if err != nil {
				return diag.WrapError(err)
			}
			res <- &Template{
				CreatedTimestamp: templateMeta.CreatedTimestamp,
				Template:         template.Template,
			}
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		listInput.NextToken = output.NextToken
	}

	return nil
}
func ResolveSesTemplateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"template", *resource.Item.(*Template).TemplateName}, nil
	})(ctx, meta, resource, c)
}
