package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ManagedRuleSets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cdn_managed_rule_sets",
		Resolver:             fetchManagedRuleSets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn@v1.0.0#ManagedRuleSetDefinition",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cdn_managed_rule_sets", client.Namespacemicrosoft_cdn),
		Transform:            transformers.TransformWithStruct(&armcdn.ManagedRuleSetDefinition{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchManagedRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcdn.NewManagedRuleSetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
