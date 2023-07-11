package logic

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:                 "azure_logic_workflows",
		Resolver:             fetchWorkflows,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/logic/workflows/list-by-subscription?tabs=HTTP#workflow",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_logic_workflows", client.Namespacemicrosoft_logic),
		Transform:            transformers.TransformWithStruct(&armlogic.Workflow{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armlogic.NewWorkflowsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
