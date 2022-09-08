// Auto generated code - DO NOT EDIT.

package servicebus

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/servicebus/mgmt/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func accessKeys() *schema.Table {
	return &schema.Table{
		Name:     "azure_servicebus_access_keys",
		Resolver: fetchServicebusAccessKeys,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "primary_connection_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrimaryConnectionString"),
			},
			{
				Name:     "secondary_connection_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryConnectionString"),
			},
			{
				Name:     "alias_primary_connection_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AliasPrimaryConnectionString"),
			},
			{
				Name:     "alias_secondary_connection_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AliasSecondaryConnectionString"),
			},
			{
				Name:     "primary_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrimaryKey"),
			},
			{
				Name:     "secondary_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryKey"),
			},
			{
				Name:     "key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyName"),
			},
		},
	}
}

func fetchServicebusAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicebus.AccessKeys

	namespace := parent.Parent.Parent.Item.(servicebus.SBNamespace)
	topic := parent.Parent.Item.(servicebus.SBTopic)
	rule := parent.Item.(servicebus.SBAuthorizationRule)
	resourceDetails, err := client.ParseResourceID(*rule.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.ListKeys(ctx, resourceDetails.ResourceGroup, *namespace.Name, *topic.Name, *rule.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
