package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConfigurationSetEventDestinations() *schema.Table {
	tableName := "aws_ses_configuration_set_event_destinations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html`,
		Resolver:    fetchSesConfigurationSetEventDestinations,
		Transform:   client.TransformWithStruct(&types.EventDestination{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "email"),
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
