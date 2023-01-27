package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GovernanceRule() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_governance_rule",
		Resolver:    fetchGovernanceRule,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity@v0.9.0#GovernanceRule",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_governance_rule", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.GovernanceRule{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			client.IDColumn,
		},
	}
}

func fetchGovernanceRule(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewGovernanceRuleClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
