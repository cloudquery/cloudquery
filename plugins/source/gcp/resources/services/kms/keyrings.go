package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func KeyRings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_keyrings",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing`,
		Resolver:    fetchKeyrings,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.KeyRing{}, client.Options()...),
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
			CryptoKeys(),
			ImportJobs(),
		},
	}
}
