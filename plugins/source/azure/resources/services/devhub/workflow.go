package devhub

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Workflow() *schema.Table {
	return &schema.Table{
		Name:                 "azure_devhub_workflow",
		Resolver:             fetchWorkflow,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub@v0.2.0#Workflow",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_devhub_workflow", client.Namespacemicrosoft_devhub),
		Transform:            transformers.TransformWithStruct(&armdevhub.Workflow{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWorkflow(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdevhub.NewWorkflowClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
