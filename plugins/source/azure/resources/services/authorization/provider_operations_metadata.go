package authorization

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ProviderOperationsMetadata() *schema.Table {
	return &schema.Table{
		Name:                 "azure_authorization_provider_operations_metadata",
		Resolver:             fetchProviderOperationsMetadata,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/authorization/provider-operations-metadata/list?tabs=HTTP#provideroperationsmetadata",
		Transform:            transformers.TransformWithStruct(&armauthorization.ProviderOperationsMetadata{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchProviderOperationsMetadata(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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
