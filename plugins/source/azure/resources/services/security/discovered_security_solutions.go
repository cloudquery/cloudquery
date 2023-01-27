package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DiscoveredSecuritySolutions() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_discovered_security_solutions",
		Resolver:    fetchDiscoveredSecuritySolutions,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/discovered-security-solutions/list?tabs=HTTP#discoveredsecuritysolution",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_discovered_security_solutions", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.DiscoveredSecuritySolution{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			client.IDColumn,
		},
	}
}

func fetchDiscoveredSecuritySolutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewDiscoveredSecuritySolutionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
