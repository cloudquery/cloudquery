package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Topology() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_topology",
		Resolver:    fetchTopology,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/topology/list?tabs=HTTP#topologyresource",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_topology", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.TopologyResource{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			client.IDColumn,
		},
	}
}

func fetchTopology(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewTopologyClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
