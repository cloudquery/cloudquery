package rulesets

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ZoneRulesets() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_zone_rulesets",
		Resolver:    fetchZoneRulesets,
		Description: "https://developers.cloudflare.com/api/operations/listZoneRulesets",
		Transform:   client.TransformWithStruct(&cloudflare.Ruleset{}),
		Multiplex:   client.ZoneMultiplex,
		Columns: []schema.Column{
			{
				Name:        "zone_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveZoneID,
				Description: `Zone identifier tag.`,
				PrimaryKey:  true,
			},
			{
				Name:       "ruleset_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchZoneRulesets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId
	params := cloudflare.ListRulesetsParams{}
	resp, err := svc.ClientApi.ListRulesets(ctx, cloudflare.ZoneIdentifier(zoneId), params)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
