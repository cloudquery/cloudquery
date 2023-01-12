package vpcaccess

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vpcaccess_locations",
		Description: `https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/Shared.Types/ListLocationsResponse#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("vpcaccess.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&locationpb.Location{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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

		Relations: []*schema.Table{
			Connectors(),
		},
	}
}
