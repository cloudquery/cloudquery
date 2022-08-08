package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource dns_records --config dns_records.hcl --output .
func DNSRecords() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_dns_records",
		Description: "DNSRecord represents a DNS record in a zone.",
		Multiplex:   client.ZoneMultiplex,
		Resolver:    fetchDnsRecords,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "created_on",
				Description: "When the record was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified_on",
				Description: "When the record was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "type",
				Description: "Record type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "DNS record name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "content",
				Description: "A valid IPv4 address.",
				Type:        schema.TypeString,
			},
			{
				Name:        "meta",
				Description: "Extra Cloudflare-specific information about the record.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "data",
				Description: "Metadata about the record.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "DNS record identifier tag.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "zone_id",
				Description: "Zone identifier tag.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ZoneID"),
			},
			{
				Name:        "zone_name",
				Description: "The domain of the record.",
				Type:        schema.TypeString,
			},
			{
				Name:        "priority",
				Description: "The priority of the record.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "ttl",
				Description: "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("TTL"),
			},
			{
				Name:        "proxied",
				Description: "Whether the record is receiving the performance and security benefits of Cloudflare.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "proxiable",
				Description: "Whether the record can be proxied by Cloudflare or not.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "locked",
				Description: "Whether this record can be modified/deleted (true means it's managed by Cloudflare).",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDnsRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	records, err := svc.ClientApi.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- records
	return nil
}
