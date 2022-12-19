// Code generated by codegen; DO NOT EDIT.

package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConfigurationSetEventDestinations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ses_configuration_set_event_destinations",
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html`,
		Resolver:    fetchSesConfigurationSetEventDestinations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("email"),
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
				Name:     "configuration_set_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "matching_event_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MatchingEventTypes"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cloud_watch_destination",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudWatchDestination"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "kinesis_firehose_destination",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KinesisFirehoseDestination"),
			},
			{
				Name:     "pinpoint_destination",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PinpointDestination"),
			},
			{
				Name:     "sns_destination",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SnsDestination"),
			},
		},
	}
}
