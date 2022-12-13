package resources

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func recipes() []*Resource {
	return []*Resource{
		{
			Name:    "tailscale_acls",
			Service: "acl",
			Struct:  new(tailscale.ACL),
		},
		{
			Name:       "tailscale_dns",
			Service:    "dns",
			SubService: "dns",
			Struct:     new(tailscale.DNSPreferences),
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "nameservers",
					Type:     schema.TypeStringArray,
					Resolver: "fetchDNSNameservers",
				},
				{
					Name:     "search_paths",
					Type:     schema.TypeStringArray,
					Resolver: "fetchDNSSearchPaths",
				},
			},
		},
		{
			Name:    "tailscale_devices",
			Service: "device",
			Struct:  new(tailscale.Device),
			Children: []*Resource{
				{
					SubService: "routes",
					Struct:     new(tailscale.DeviceRoutes),
					ExtraColumns: codegen.ColumnDefinitions{
						{
							Name:     "device_id",
							Type:     schema.TypeString,
							Resolver: `schema.ParentColumnResolver("id")`,
							Options:  schema.ColumnCreationOptions{PrimaryKey: true},
						},
					},
				},
			},
		},
		{
			Name:        "tailscale_keys",
			Service:     "key",
			Struct:      new(tailscale.Key),
			PreResolver: "getKey",
		},
	}
}
