package operationalinsights

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Workspaces() *schema.Table {
	return &schema.Table{
		Name:                 "azure_operationalinsights_workspaces",
		Resolver:             fetchWorkspaces,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights@v1.0.0#Workspace",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_operationalinsights_workspaces", client.Namespacemicrosoft_operationalinsights),
		Transform:            transformers.TransformWithStruct(&armoperationalinsights.Workspace{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWorkspaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armoperationalinsights.NewWorkspacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
