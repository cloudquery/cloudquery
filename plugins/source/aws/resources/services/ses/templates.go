package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Templates() *schema.Table {
	tableName := "aws_ses_templates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailTemplate.html`,
		Resolver:            fetchSesTemplates,
		PreResourceResolver: getTemplate,
		Transform:           client.TransformWithStruct(&models.Template{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "email"),
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
