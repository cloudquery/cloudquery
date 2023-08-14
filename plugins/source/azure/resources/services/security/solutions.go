package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Solutions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_solutions",
		Resolver:             fetchSolutions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity@v0.9.0#Solution",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_solutions", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.Solution{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSolutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewSolutionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
