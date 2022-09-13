package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Zones() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_zones",
		Description: "Zone describes a Cloudflare zone.",
		Resolver:    fetchZones,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
			},
			{
				Name:        "host",
				Description: "Zone host details.",
				Type:        schema.TypeJSON,
			},
			{
				Name:            "id",
				Description:     "The unique universal identifier for a Cloudflare zone.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "Cloudflare zone name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dev_mode",
				Description: "DevMode contains the time in seconds until development expires (if positive) or since it expired (if negative)",
				Type:        schema.TypeInt,
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
				Name:        "owner",
				Description: "Zone owner details.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "permissions",
				Description: "Zone permissions.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "plan",
				Description: "Cloudflare zone plan details.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "plan_pending",
				Description: "Cloudflare pending zone plan details.",
				Type:        schema.TypeJSON,
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
				Name:        "meta",
				Type:        schema.TypeJSON,
				Description: "Zone meta information.",
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
