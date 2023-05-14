package vpcaccess

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vpcaccess_locations",
		Description: `https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/Shared.Types/ListLocationsResponse#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("vpcaccess.googleapis.com"),
		Transform:   client.TransformWithStruct(&locationpb.Location{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
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
