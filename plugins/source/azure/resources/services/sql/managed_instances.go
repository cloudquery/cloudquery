package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ManagedInstances() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_managed_instances",
		Resolver:             fetchManagedInstances,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instances/list?tabs=HTTP#managedinstance",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_sql_managed_instances", client.Namespacemicrosoft_sql),
		Transform:            transformers.TransformWithStruct(&armsql.ManagedInstance{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			managedInstanceEncryptionProtectors(),
			managedInstanceVulnerabilityAssessments(),
		},
	}
}

func fetchManagedInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsql.NewManagedInstancesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
