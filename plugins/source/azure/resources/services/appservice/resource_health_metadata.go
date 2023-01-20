package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceHealthMetadata() *schema.Table {
	return &schema.Table{
		Name:        "azure_appservice_resource_health_metadata",
		Resolver:    fetchResourceHealthMetadata,
		Description: "https://learn.microsoft.com/en-us/rest/api/appservice/resource-health-metadata/list?tabs=HTTP#resourcehealthmetadata",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_resource_health_metadata", client.Namespacemicrosoft_web),
		Transform:   transformers.TransformWithStruct(&armappservice.ResourceHealthMetadata{}),
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

func fetchResourceHealthMetadata(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappservice.NewResourceHealthMetadataClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
