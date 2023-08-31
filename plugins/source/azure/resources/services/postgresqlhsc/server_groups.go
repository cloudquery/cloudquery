package postgresqlhsc

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ServerGroups() *schema.Table {
	return &schema.Table{
		Name:                 "azure_postgresqlhsc_server_groups",
		Resolver:             fetchServerGroups,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc@v0.5.0#ServerGroup",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_postgresqlhsc_server_groups", client.Namespacemicrosoft_dbforpostgresql),
		Transform:            transformers.TransformWithStruct(&armpostgresqlhsc.ServerGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchServerGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armpostgresqlhsc.NewServerGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
