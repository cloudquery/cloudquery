// Code generated by codegen; DO NOT EDIT.

package frontdoor

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Doors() *schema.Table {
	return &schema.Table{
		Name:        "azure_front_doors",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor#FrontDoor`,
		Resolver:    fetchDoors,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "backend_pools",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.BackendPools"),
			},
			{
				Name:     "backend_pools_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.BackendPoolsSettings"),
			},
			{
				Name:     "enabled_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.EnabledState"),
			},
			{
				Name:     "friendly_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.FriendlyName"),
			},
			{
				Name:     "frontend_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.FrontendEndpoints"),
			},
			{
				Name:     "health_probe_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.HealthProbeSettings"),
			},
			{
				Name:     "load_balancing_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.LoadBalancingSettings"),
			},
			{
				Name:     "routing_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.RoutingRules"),
			},
			{
				Name:     "cname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Cname"),
			},
			{
				Name:     "frontdoor_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.FrontdoorID"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:     "resource_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ResourceState"),
			},
			{
				Name:     "rules_engines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.RulesEngines"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
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
