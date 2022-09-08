// Auto generated code - DO NOT EDIT.

package eventhub

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/eventhub/mgmt/eventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func networkRuleSets() *schema.Table {
	return &schema.Table{
		Name:     "azure_eventhub_network_rule_sets",
		Resolver: fetchEventHubNetworkRuleSets,
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
		return errors.WithStack(err)
	}
	response, err := svc.GetNetworkRuleSet(ctx, resource.ResourceGroup, *namespace.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
	return nil
}
