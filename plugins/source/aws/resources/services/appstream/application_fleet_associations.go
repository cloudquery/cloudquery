package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApplicationFleetAssociations() *schema.Table {
	tableName := "aws_appstream_application_fleet_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html`,
		Resolver:    fetchAppstreamApplicationFleetAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   client.TransformWithStruct(&types.ApplicationFleetAssociation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "application_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "fleet_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FleetName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
