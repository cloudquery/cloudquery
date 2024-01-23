package kms

import (
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
			client.ProjectIDColumn(true),
		},
		Relations: []*schema.Table{
			KeyRings(),
			EkmConnections(),
		},
	}
}
