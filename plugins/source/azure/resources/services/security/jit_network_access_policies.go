package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func JitNetworkAccessPolicies() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_jit_network_access_policies",
		Resolver:    fetchJitNetworkAccessPolicies,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/jit-network-access-policies/list?tabs=HTTP#jitnetworkaccesspolicy",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_jit_network_access_policies", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.JitNetworkAccessPolicy{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
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

func fetchJitNetworkAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewJitNetworkAccessPoliciesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
