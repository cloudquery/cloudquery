package certificate_packs

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func CertificatePacks() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_certificate_packs",
		Resolver:  fetchCertificatePacks,
		Multiplex: client.ZoneMultiplex,
		Transform: client.TransformWithStruct(&cloudflare.CertificatePack{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:        "zone_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveZoneID,
				Description: `Zone identifier tag.`,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
