package notificationhubs

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:                 "azure_notificationhubs_namespaces",
		Resolver:             fetchNamespaces,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/notificationhubs/namespaces/list?tabs=HTTP#namespaceresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_notificationhubs_namespaces", client.Namespacemicrosoft_notificationhubs),
		Transform:            transformers.TransformWithStruct(&armnotificationhubs.NamespaceResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnotificationhubs.NewNamespacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
