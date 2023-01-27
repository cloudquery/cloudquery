package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_applications",
		Resolver:    fetchApplications,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity@v0.9.0#Application",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_applications", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.Application{}),
		Columns:     schema.ColumnList{client.SubscriptionID, client.IDColumn},
	}
}

func fetchApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewApplicationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
