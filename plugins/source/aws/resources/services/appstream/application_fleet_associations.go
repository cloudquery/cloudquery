// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApplicationFleetAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_application_fleet_associations",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html`,
		Resolver:    fetchAppstreamApplicationFleetAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
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
