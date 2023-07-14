package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func longTermRetentionPolicies() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_database_long_term_retention_policies",
		Resolver:             fetchLongTermRetentionPolicies,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/long-term-retention-policies/list-by-database?tabs=HTTP#longtermretentionpolicy",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_sql_long_term_retention_policies", client.Namespacemicrosoft_sql),
		Transform:            transformers.TransformWithStruct(&armsql.LongTermRetentionPolicy{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchLongTermRetentionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pp := parent.Parent.Item.(*armsql.Server)
	p := parent.Item.(*armsql.Database)
	cl := meta.(*client.Client)
	svc, err := armsql.NewLongTermRetentionPoliciesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByDatabasePager(group, *pp.Name, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
