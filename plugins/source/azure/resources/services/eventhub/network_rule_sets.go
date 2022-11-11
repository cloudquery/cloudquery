// Auto generated code - DO NOT EDIT.

package eventhub

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
)

func networkRuleSets() *schema.Table {
	return &schema.Table{
		Name:        "azure_eventhub_network_rule_sets",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub#NetworkRuleSet`,
		Resolver:    fetchEventHubNetworkRuleSets,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "eventhub_namespace_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "trusted_service_access_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TrustedServiceAccessEnabled"),
			},
			{
				Name:     "default_action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultAction"),
			},
			{
				Name:     "virtual_network_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualNetworkRules"),
			},
			{
				Name:     "ip_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPRules"),
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
	}
}

func fetchEventHubNetworkRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().EventHub.NetworkRuleSets

	namespace := parent.Item.(eventhub.EHNamespace)
	resource, err := client.ParseResourceID(*namespace.ID)
	if err != nil {
		return err
	}
	response, err := svc.GetNetworkRuleSet(ctx, resource.ResourceGroup, *namespace.Name)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
