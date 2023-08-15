package frontdoor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ManagedRuleSets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_frontdoor_managed_rule_sets",
		Resolver:             fetchManagedRuleSets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/frontdoor/webapplicationfirewall/managed-rule-sets/list#managedrulesetdefinition",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_frontdoor_managed_rule_sets", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armfrontdoor.ManagedRuleSetDefinition{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchManagedRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armfrontdoor.NewManagedRuleSetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
