package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func CryptoKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_crypto_keys",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKey`,
		Resolver:    fetchCryptoKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.CryptoKey{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
			{
				Name:     "rotation_period",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: resolveRotationPeriod,
			},
		},
		Relations: []*schema.Table{
			CryptoKeyVersions(),
		},
	}
}
