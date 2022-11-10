// Auto generated code - DO NOT EDIT.

package servicebus

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func authorizationRules() *schema.Table {
	return &schema.Table{
		Name:        "azure_servicebus_authorization_rules",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus#SBAuthorizationRule`,
		Resolver:    fetchServicebusAuthorizationRules,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "servicebus_topic_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "rights",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Rights"),
			},
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			accessKeys(),
		},
	}
}

func fetchServicebusAuthorizationRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicebus.AuthorizationRules

	namespace := parent.Parent.Item.(servicebus.SBNamespace)
	topic := parent.Item.(servicebus.SBTopic)
	resourceDetails, err := client.ParseResourceID(*topic.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListAuthorizationRules(ctx, resourceDetails.ResourceGroup, *namespace.Name, *topic.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
