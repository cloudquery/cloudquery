package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func CryptoKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_crypto_keys",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKey`,
		Resolver:    fetchCryptoKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.CryptoKey{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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
				Name:     "rotation_period",
				Type:     schema.TypeInt,
				Resolver: resolveRotationPeriod,
			},
		},
		Relations: []*schema.Table{
			CryptoKeyVersions(),
		},
	}
}
