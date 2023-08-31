package azurearcdata

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SqlManagedInstances() *schema.Table {
	return &schema.Table{
		Name:                 "azure_azurearcdata_sql_managed_instances",
		Resolver:             fetchSqlManagedInstances,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata@v0.5.0#SQLManagedInstance",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_azurearcdata_sql_managed_instances", client.Namespacemicrosoft_azurearcdata),
		Transform:            transformers.TransformWithStruct(&armazurearcdata.SQLManagedInstance{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSqlManagedInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armazurearcdata.NewSQLManagedInstancesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
