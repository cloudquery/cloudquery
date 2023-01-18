package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Identities() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_identities",
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailIdentity.html`,
		Resolver:            fetchSesIdentities,
		PreResourceResolver: getIdentity,
		Transform:           transformers.TransformWithStruct(&models.EmailIdentity{}, transformers.WithUnwrapStructFields("GetEmailIdentityOutput"), transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
