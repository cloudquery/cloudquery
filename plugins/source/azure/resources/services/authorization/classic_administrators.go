package authorization

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ClassicAdministrators() *schema.Table {
	return &schema.Table{
		Name:                 "azure_authorization_classic_administrators",
		Resolver:             fetchClassicAdministrators,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/authorization/classic-administrators/list?tabs=HTTP#classicadministrator",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_authorization_classic_administrators", client.Namespacemicrosoft_authorization),
		Transform:            transformers.TransformWithStruct(&armauthorization.ClassicAdministrator{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchClassicAdministrators(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armauthorization.NewClassicAdministratorsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
