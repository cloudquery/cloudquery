// Auto generated code - DO NOT EDIT.

package frontdoor

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func Doors() *schema.Table {
	return &schema.Table{
		Name:      "azure_frontdoor_doors",
		Resolver:  fetchFrontDoorDoors,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "resource_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceState"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "cname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Cname"),
			},
			{
				Name:     "frontdoor_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FrontdoorID"),
			},
			{
				Name:     "rules_engines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RulesEngines"),
			},
			{
				Name:     "friendly_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FriendlyName"),
			},
			{
				Name:     "routing_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RoutingRules"),
			},
			{
				Name:     "load_balancing_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoadBalancingSettings"),
			},
			{
				Name:     "health_probe_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthProbeSettings"),
			},
			{
				Name:     "backend_pools",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackendPools"),
			},
			{
				Name:     "frontend_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FrontendEndpoints"),
			},
			{
				Name:     "backend_pools_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackendPoolsSettings"),
			},
			{
				Name:     "enabled_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnabledState"),
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
	}
}

func fetchFrontDoorDoors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().FrontDoor.Doors

	response, err := svc.List(ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
