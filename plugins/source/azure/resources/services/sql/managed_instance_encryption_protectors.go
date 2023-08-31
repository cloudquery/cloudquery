package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func managedInstanceEncryptionProtectors() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_managed_instance_encryption_protectors",
		Resolver:             fetchManagedInstanceEncryptionProtectors,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instance-encryption-protectors/list-by-instance?tabs=HTTP#managedinstanceencryptionprotector",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_sql_managed_instance_encryption_protectors", client.Namespacemicrosoft_sql),
		Transform:            transformers.TransformWithStruct(&armsql.ManagedInstanceEncryptionProtector{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchManagedInstanceEncryptionProtectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armsql.ManagedInstance)
	cl := meta.(*client.Client)
	svc, err := armsql.NewManagedInstanceEncryptionProtectorsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByInstancePager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
