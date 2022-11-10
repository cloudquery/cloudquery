// Auto generated code - DO NOT EDIT.

package container

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:        "azure_container_registries",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry#Registry`,
		Resolver:    fetchContainerRegistries,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "login_server",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoginServer"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "admin_user_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AdminUserEnabled"),
			},
			{
				Name:     "storage_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StorageAccount"),
			},
			{
				Name:     "network_rule_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkRuleSet"),
			},
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policies"),
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
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			replications(),
		},
	}
}

func fetchContainerRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Container.Registries

	response, err := svc.List(ctx)

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
