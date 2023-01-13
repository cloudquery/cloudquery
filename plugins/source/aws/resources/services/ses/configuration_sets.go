package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigurationSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_configuration_sets",
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html`,
		Resolver:            fetchSesConfigurationSets,
		PreResourceResolver: getConfigurationSet,
		Transform: transformers.TransformWithStruct(
			&sesv2.GetConfigurationSetOutput{},
			transformers.WithSkipFields("ResultMetadata"),
			transformers.WithNameTransformer(client.CreateTrimPrefixTransformer("configuration_set_")),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer("email"),
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
				Resolver: resolveConfigurationSetArn,
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

		Relations: []*schema.Table{
			ConfigurationSetEventDestinations(),
		},
	}
}
