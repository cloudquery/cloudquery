package postgresqlflexibleservers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func serverConfigurations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_postgresqlflexibleservers_server_configurations",
		Resolver:             fetchServerConfigurations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description: "https://learn.microsoft.com/en-us/rest/api/postgresql/flexibleserver/configurations/list-by-server?view=rest-postgresql-flexibleserver-2022-12-01&tabs=HTTP",
		Transform: 		  transformers.TransformWithStruct(&armpostgresqlflexibleservers.Configuration{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}

func fetchServerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armpostgresqlflexibleservers.Server)
	cl := meta.(*client.Client)
	svc, err := armpostgresqlflexibleservers.NewConfigurationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}

	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByServerPager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
