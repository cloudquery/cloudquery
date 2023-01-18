package dns

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func Dns() *schema.Table {
	return &schema.Table{
		Name:        "tailscale_dns",
		Description: `https://pkg.go.dev/github.com/tailscale/tailscale-client-go/tailscale#DNSPreferences`,
		Resolver:    fetchDns,
		Transform:   transformers.TransformWithStruct(&tailscale.DNSPreferences{}),
		Columns: []schema.Column{
			{
				Name:     "tailnet",
				Type:     schema.TypeString,
				Resolver: client.ResolveTailnet,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "nameservers",
				Type:     schema.TypeStringArray,
				Resolver: fetchDNSNameservers,
			},
			{
				Name:     "search_paths",
				Type:     schema.TypeStringArray,
				Resolver: fetchDNSSearchPaths,
			},
		},
	}
}
