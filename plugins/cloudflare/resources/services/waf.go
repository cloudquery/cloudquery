package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource waf --config waf.hcl --output .
func Wafs() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_waf",
		Description: "WAFPackage represents a WAF package configuration.",
		Resolver:    fetchWafs,
		Multiplex:   client.ZoneMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "id",
				Description: "The unique identifier of a WAF package.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the WAF package.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A summary of the purpose/function of the WAF package.",
				Type:        schema.TypeString,
			},
			{
				Name:        "zone_id",
				Description: "Zone identifier tag.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ZoneID"),
			},
			{
				Name:        "detection_mode",
				Description: "When a WAF package uses anomaly detection, each rule is given a score when triggered. If the total score of all triggered rules exceeds the sensitivity defined on the WAF package, the action defined on the package will be taken.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sensitivity",
				Description: "The sensitivity of the WAF package",
				Type:        schema.TypeString,
			},
			{
				Name:        "action_mode",
				Description: "The default action performed by the rules in the WAF package.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "cloudflare_waf_rule_groups",
				Description: "WAFGroup represents a WAF rule group.",
				Resolver:    fetchWafRuleGroups,
				Columns: []schema.Column{
					{
						Name:        "waf_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_waf table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAccountId,
					},
					{
						Name:        "zone_id",
						Description: "Zone identifier tag.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveZoneId,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the rule group.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the rule group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "An informative summary of what the rule group does.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rules_count",
						Description: "The number of rules in the current rule group.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "modified_rules_count",
						Description: "The number of rules within the group that have been modified from their default configuration.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "package_id",
						Description: "The unique identifier of a WAF package.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PackageID"),
					},
					{
						Name:        "mode",
						Description: "The state of the rules contained in the rule group. When on, the rules in the group are configurable/usable.",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_modes",
						Description: "The available states for the rule group.",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "cloudflare_waf_rules",
				Description: "WAFRule represents a WAF rule.",
				Resolver:    fetchWafRules,
				Columns: []schema.Column{
					{
						Name:        "waf_cq_id",
						Description: "Unique CloudQuery ID of cloudflare_waf table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAccountId,
					},
					{
						Name:        "zone_id",
						Description: "Zone identifier tag.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveZoneId,
					},
					{
						Name:        "group",
						Description: "The rule group to which the current WAF rule belongs.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the WAF rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "description",
						Description: "The public description of the WAF rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "priority",
						Description: "The order in which the individual WAF rule is executed within its rule group.",
						Type:        schema.TypeString,
					},
					{
						Name:     "package_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PackageID"),
					},
					{
						Name:        "mode",
						Description: "The action that the current WAF rule will perform when triggered. Applies to traditional (deny) WAF rules.",
						Type:        schema.TypeString,
					},
					{
						Name:        "default_mode",
						Description: "The default action/mode of a rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_modes",
						Description: "The list of possible actions of the WAF rule when it is triggered.",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWAFPackages(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func fetchWafRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId
	pack := parent.Item.(cloudflare.WAFPackage)

	resp, err := svc.ClientApi.ListWAFGroups(ctx, zoneId, pack.ID)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func fetchWafRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId
	pack := parent.Item.(cloudflare.WAFPackage)

	resp, err := svc.ClientApi.ListWAFRules(ctx, zoneId, pack.ID)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
