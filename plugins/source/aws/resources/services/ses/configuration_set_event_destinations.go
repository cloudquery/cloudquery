package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigurationSetEventDestinations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ses_configuration_set_event_destinations",
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html`,
		Resolver:    fetchSesConfigurationSetEventDestinations,
		Transform:   transformers.TransformWithStruct(&types.EventDestination{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "configuration_set_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
