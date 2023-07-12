package mysqlflexibleservers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_mysqlflexibleservers_servers",
		Resolver:             fetchServers,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/mysql/flexibleserver/servers/list?tabs=HTTP#server",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_mysqlflexibleservers_servers", client.Namespacemicrosoft_dbformysql),
		Transform:            transformers.TransformWithStruct(&armmysqlflexibleservers.Server{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmysqlflexibleservers.NewServersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
