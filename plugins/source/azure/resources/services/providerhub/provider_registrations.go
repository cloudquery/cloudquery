package providerhub

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ProviderRegistrations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_providerhub_provider_registrations",
		Resolver:             fetchProviderRegistrations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub@v1.0.0#ProviderRegistration",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_providerhub_provider_registrations", client.Namespacemicrosoft_providerhub),
		Transform:            transformers.TransformWithStruct(&armproviderhub.ProviderRegistration{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchProviderRegistrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armproviderhub.NewProviderRegistrationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
