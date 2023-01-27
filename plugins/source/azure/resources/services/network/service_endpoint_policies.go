package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ServiceEndpointPolicies() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_service_endpoint_policies",
		Resolver:    fetchServiceEndpointPolicies,
		Description: "https://learn.microsoft.com/en-us/rest/api/expressroute/service-endpoint-policies/list?tabs=HTTP#serviceendpointpolicy",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_service_endpoint_policies", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.ServiceEndpointPolicy{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchServiceEndpointPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewServiceEndpointPoliciesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
