package kms

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_locations",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   client.TransformWithStruct(&locationpb.Location{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			KeyRings(),
			EkmConnections(),
		},
	}
}
