package ips

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/schema"
)

type IpWrapper struct {
	Ip   string
	Type string
}

func IPs() *schema.Table {
	return &schema.Table{
		Name:     "cloudflare_ips",
		Resolver: fetchIPs,
		Columns: []schema.Column{
			{
				Name:        "ip",
				Description: "Cloudflare ip cidr address.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Ip type, ipv4, ipv6, ipv4_china, ipv6_china.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIPs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	resp, err := cloudflare.IPs()
	if err != nil {
		return err
	}

	for _, ip := range resp.IPv4CIDRs {
		res <- IpWrapper{Ip: ip, Type: "ipv4"}
	}

	for _, ip := range resp.IPv6CIDRs {
		res <- IpWrapper{Ip: ip, Type: "ipv6"}
	}

	for _, ip := range resp.ChinaIPv4CIDRs {
		res <- IpWrapper{Ip: ip, Type: "ipv4_china"}
	}

	for _, ip := range resp.ChinaIPv6CIDRs {
		res <- IpWrapper{Ip: ip, Type: "ipv6_china"}
	}

	return nil
}
