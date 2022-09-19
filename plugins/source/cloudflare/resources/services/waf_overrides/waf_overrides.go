// Code generated by codegen using template resource_manual.go.tpl; DO NOT EDIT.

package waf_overrides

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFOverrides() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_waf_overrides",
		Resolver:  fetchWAFOverrides,
		Multiplex: client.ZoneMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:        "zone_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneID,
				Description: `Zone identifier tag.`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "urls",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("URLs"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Groups"),
			},
			{
				Name:     "rewrite_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RewriteAction"),
			},
			{
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rules"),
			},
			{
				Name:     "paused",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Paused"),
			},
		},
	}
}
