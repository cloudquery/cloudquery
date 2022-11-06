// Code generated by codegen; DO NOT EDIT.

package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConfigurationSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_configuration_sets",
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html`,
		Resolver:            fetchSesConfigurationSets,
		PreResourceResolver: getConfigurationSet,
		Multiplex:           client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigurationSetName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "delivery_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryOptions"),
			},
			{
				Name:     "reputation_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReputationOptions"),
			},
			{
				Name:     "sending_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SendingOptions"),
			},
			{
				Name:     "suppression_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SuppressionOptions"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "tracking_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrackingOptions"),
			},
			{
				Name:     "vdm_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VdmOptions"),
			},
		},

		Relations: []*schema.Table{
			ConfigurationSetEventDestinations(),
		},
	}
}
