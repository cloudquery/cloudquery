package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource zones --config zones.hcl --output .
func Zones() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_zones",
		Description: "Zone describes a Cloudflare zone.",
		Resolver:    fetchZones,
		Multiplex:   client.AccountMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "host_name",
				Description: "Zone host name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Host.Name"),
			},
			{
				Name:        "host_website",
				Description: "Zone host website.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Host.Website"),
			},
			{
				Name:        "id",
				Description: "The unique universal identifier for a Cloudflare zone.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Cloudflare zone name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dev_mode",
				Description: "DevMode contains the time in seconds until development expires (if positive) or since it expired (if negative)",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "original_ns",
				Description: "Creation timestamp of the account.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("OriginalNS"),
			},
			{
				Name:        "original_registrar",
				Description: "Cloudflare zone original name servers.",
				Type:        schema.TypeString,
			},
			{
				Name:        "original_dns_host",
				Description: "Cloudflare zone original registrar.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OriginalDNSHost"),
			},
			{
				Name:        "created_on",
				Description: "Cloudflare zone original dns host.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified_on",
				Description: "Zone created on date and time.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name_servers",
				Description: "Zone last modified date and time.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "owner_id",
				Description: "Zone owner id.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.ID"),
			},
			{
				Name:        "owner_email",
				Description: "Zone owner email address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.Email"),
			},
			{
				Name:        "owner_name",
				Description: "Zone owner name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.Name"),
			},
			{
				Name:        "owner_type",
				Description: "Zone owner type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.OwnerType"),
			},
			{
				Name:        "permissions",
				Description: "Zone permissions.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "plan_id",
				Description: "The unique universal identifier for a Cloudflare zone plan.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.ZonePlanCommon.ID"),
			},
			{
				Name:        "plan_name",
				Description: "Cloudflare zone plan name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.ZonePlanCommon.Name"),
			},
			{
				Name:        "plan_price",
				Description: "Cloudflare zone plan price.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Plan.ZonePlanCommon.Price"),
			},
			{
				Name:        "plan_currency",
				Description: "Cloudflare zone plan currency.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.ZonePlanCommon.Currency"),
			},
			{
				Name:        "plan_frequency",
				Description: "Cloudflare zone plan frequency.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.ZonePlanCommon.Frequency"),
			},
			{
				Name:        "plan_legacy_id",
				Description: "True if zone plan is subscribed.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.LegacyID"),
			},
			{
				Name:        "plan_is_subscribed",
				Description: "True if zone plan can subscribe.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Plan.IsSubscribed"),
			},
			{
				Name:        "plan_can_subscribe",
				Description: "Cloudflare zone plan legacy id.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Plan.CanSubscribe"),
			},
			{
				Name:        "plan_legacy_discount",
				Description: "True if zone plan has legacy discount.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Plan.LegacyDiscount"),
			},
			{
				Name:        "plan_externally_managed",
				Description: "True if zone plan is externally managed.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Plan.ExternallyManaged"),
			},
			{
				Name:        "plan_pending_id",
				Description: "The unique universal identifier for a Cloudflare zone plan.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PlanPending.ZonePlanCommon.ID"),
			},
			{
				Name:        "plan_pending_name",
				Description: "Cloudflare zone plan name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PlanPending.ZonePlanCommon.Name"),
			},
			{
				Name:        "plan_pending_price",
				Description: "Cloudflare zone plan price.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("PlanPending.ZonePlanCommon.Price"),
			},
			{
				Name:        "plan_pending_currency",
				Description: "Cloudflare zone plan currency.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PlanPending.ZonePlanCommon.Currency"),
			},
			{
				Name:        "plan_pending_frequency",
				Description: "Cloudflare zone plan frequency.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PlanPending.ZonePlanCommon.Frequency"),
			},
			{
				Name:        "plan_pending_legacy_id",
				Description: "True if zone plan is subscribed.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PlanPending.LegacyID"),
			},
			{
				Name:        "plan_pending_is_subscribed",
				Description: "True if zone plan can subscribe.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PlanPending.IsSubscribed"),
			},
			{
				Name:        "plan_pending_can_subscribe",
				Description: "Cloudflare zone plan legacy id.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PlanPending.CanSubscribe"),
			},
			{
				Name:        "plan_pending_legacy_discount",
				Description: "True if zone plan has legacy discount.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PlanPending.LegacyDiscount"),
			},
			{
				Name:        "plan_pending_externally_managed",
				Description: "True if zone plan is externally managed.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PlanPending.ExternallyManaged"),
			},
			{
				Name:        "status",
				Description: "Zone status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "paused",
				Description: "True if zone is paused.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "type",
				Description: "Zone type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vanity_ns",
				Description: "Zone vanity name servers.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VanityNS"),
			},
			{
				Name:        "betas",
				Description: "Zone betas.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "deactivation_reason",
				Description: "Zone deactivation reason.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeactReason"),
			},
			{
				Name:        "page_rule_quota",
				Description: "custom_certificate_quota is broken - sometimes it's a string, sometimes a number! CustCertQuota     int    `json:\"custom_certificate_quota\"`",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Meta.PageRuleQuota"),
			},
			{
				Name:     "wildcard_proxiable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Meta.WildcardProxiable"),
			},
			{
				Name:     "phishing_detected",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Meta.PhishingDetected"),
			},
			{
				Name:        "verification_key",
				Description: "Zone verification key.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	opts := cloudflare.WithZoneFilters("", svc.AccountId, "")

	resp, err := svc.ClientApi.ListZonesContext(ctx, opts)
	if err != nil {
		return err
	}
	res <- resp.Result

	return nil
}
