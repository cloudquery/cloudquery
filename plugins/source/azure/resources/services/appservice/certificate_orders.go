package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CertificateOrders() *schema.Table {
	return &schema.Table{
		Name:        "azure_appservice_certificate_orders",
		Resolver:    fetchCertificateOrders,
		Description: "https://learn.microsoft.com/en-us/rest/api/appservice/app-service-certificate-orders/list?tabs=HTTP#appservicecertificateorder",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_certificate_orders", client.Namespacemicrosoft_certificateregistration),
		Transform:   transformers.TransformWithStruct(&armappservice.CertificateOrder{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchCertificateOrders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappservice.NewCertificateOrdersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
