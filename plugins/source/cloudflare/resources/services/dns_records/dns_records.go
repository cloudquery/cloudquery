package dns_records

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DNSRecords() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_dns_records",
		Resolver:  fetchDNSRecords,
		Multiplex: client.ZoneMultiplex,
		Transform: transformers.TransformWithStruct(&cloudflare.DNSRecord{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:        "meta",
				Type:        schema.TypeJSON,
				Description: `Extra Cloudflare-specific information about the record.`,
			},
			{
				Name:        "data",
				Type:        schema.TypeJSON,
				Description: `Metadata about the record.`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
