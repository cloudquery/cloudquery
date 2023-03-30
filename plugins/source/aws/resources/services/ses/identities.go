package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Identities() *schema.Table {
	tableName := "aws_ses_identities"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailIdentity.html`,
		Resolver:            fetchSesIdentities,
		PreResourceResolver: getIdentity,
		Transform:           client.TransformWithStruct(&models.EmailIdentity{}, transformers.WithUnwrapStructFields("GetEmailIdentityOutput"), transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveIdentityArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
