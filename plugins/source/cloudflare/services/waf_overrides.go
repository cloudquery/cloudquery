package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WafOverrides() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_waf_overrides",
		Description: "WAFOverride represents a WAF override.",
		Resolver:    fetchWafOverrides,
		Multiplex:   client.ZoneMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "zone_id",
				Description: "The Zone ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneId,
			},
			{
				Name:            "id",
				Description:     "The unique identifier of the WAF override.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "description",
				Description: "An informative summary of the current URI-based WAF override.",
				Type:        schema.TypeString,
			},
			{
				Name:     "urls",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("URLs"),
			},
			{
				Name:        "priority",
				Description: "The relative priority of the current URI-based WAF override when multiple overrides match a single URL. A lower number indicates higher priority. Higher priority overrides may overwrite values set by lower priority overrides.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "groups",
				Description: "An object that allows you to enable or disable WAF rule groups for the current WAF override. Each key of this object must be the ID of a WAF rule group, and each value must be a valid WAF action (usually default or disable). When creating a new URI-based WAF override, you must provide a groups object or a rules object.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "rewrite_action",
				Description: "Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "rules",
				Description: "The default action performed by the rules in the WAF package.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "paused",
				Description: "When true, indicates that the WAF package is currently paused.",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafOverrides(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWAFOverrides(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
