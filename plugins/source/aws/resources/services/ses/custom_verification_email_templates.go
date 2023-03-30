package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CustomVerificationEmailTemplates() *schema.Table {
	tableName := "aws_ses_custom_verification_email_templates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetCustomVerificationEmailTemplate.html`,
		Resolver:            fetchSesCustomVerificationEmailTemplates,
		PreResourceResolver: getCustomVerificationEmailTemplate,
		Transform: client.TransformWithStruct(&sesv2.GetCustomVerificationEmailTemplateOutput{},
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
