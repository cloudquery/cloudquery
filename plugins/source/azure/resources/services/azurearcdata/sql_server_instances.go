package azurearcdata

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SqlServerInstances() *schema.Table {
	return &schema.Table{
		Name:                 "azure_azurearcdata_sql_server_instances",
		Resolver:             fetchSqlServerInstances,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata@v0.5.0#SQLServerInstance",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_azurearcdata_sql_server_instances", client.Namespacemicrosoft_azurearcdata),
		Transform:            transformers.TransformWithStruct(&armazurearcdata.SQLServerInstance{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSqlServerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armazurearcdata.NewSQLServerInstancesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
