package authorization

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ProviderOperationsMetadata() *schema.Table {
	return &schema.Table{
		Name:      "azure_authorization_provider_operations_metadata",
		Resolver:  fetchProviderOperationsMetadata,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_authorization_provider_operations_metadata", client.Namespacemicrosoft_authorization),
		Transform: transformers.TransformWithStruct(&armauthorization.ProviderOperationsMetadata{}),
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

func fetchProviderOperationsMetadata(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armauthorization.NewProviderOperationsMetadataClient(cl.Creds, cl.Options)
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
