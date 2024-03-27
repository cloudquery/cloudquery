package rulesets

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func AccountRulesets() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_account_rulesets",
		Resolver:  fetchAccountRulesets,
		Transform: client.TransformWithStruct(&cloudflare.Ruleset{}),
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "ruleset_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ID"),
			},
		},
	}
}

func fetchAccountRulesets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	accountID := svc.AccountId
	params := cloudflare.ListRulesetsParams{}
	resp, err := svc.ClientApi.ListRulesets(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
