package mysql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func databases() *schema.Table {
	return &schema.Table{
		Name:                 "azure_mysql_server_databases",
		Resolver:             fetchDatabases,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/databases/list-by-server?tabs=HTTP#database",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_mysql_server_databases", client.Namespacemicrosoft_dbformysql),
		Transform:            transformers.TransformWithStruct(&armmysql.Database{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}

func fetchDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmysql.NewDatabasesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	server := parent.Item.(*armmysql.Server)
	group, err := client.ParseResourceGroup(*server.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByServerPager(group, *server.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
