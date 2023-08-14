package dns_records

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func DNSRecords() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_dns_records",
		Resolver:  fetchDNSRecords,
		Multiplex: client.ZoneMultiplex,
		Transform: client.TransformWithStruct(&cloudflare.DNSRecord{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:        "meta",
				Type:        types.ExtensionTypes.JSON,
				Description: `Extra Cloudflare-specific information about the record.`,
			},
			{
				Name:        "data",
				Type:        types.ExtensionTypes.JSON,
				Description: `Metadata about the record.`,
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
