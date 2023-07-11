package sqlvirtualmachine

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sqlvirtualmachine_groups",
		Resolver:             fetchGroups,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sqlvm/2022-07-01-preview/sql-virtual-machine-groups/list?tabs=HTTP#sqlvirtualmachinegroup",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_sqlvirtualmachine_groups", client.Namespacemicrosoft_sqlvirtualmachine),
		Transform:            transformers.TransformWithStruct(&armsqlvirtualmachine.Group{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsqlvirtualmachine.NewGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
