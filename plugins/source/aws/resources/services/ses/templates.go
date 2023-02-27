package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Templates() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_templates",
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailTemplate.html`,
		Resolver:            fetchSesTemplates,
		PreResourceResolver: getTemplate,
		Transform:           transformers.TransformWithStruct(&models.Template{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveTemplateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
