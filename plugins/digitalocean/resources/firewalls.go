package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Firewalls() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_firewalls",
		Description:  "Firewall represents a DigitalOcean Firewall configuration.",
		Resolver:     fetchFirewalls,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique ID that can be used to identify and reference a firewall.",
				Type:        schema.TypeUUID,
				Resolver:    schema.UUIDResolver("ID"),
			},
			{
				Name:        "name",
				Description: "A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "A status string indicating the current state of the firewall. This can be \"waiting\", \"succeeded\", or \"failed\".",
				Type:        schema.TypeString,
			},
			{
				Name:        "droplet_ids",
				Description: "An array containing the IDs of the Droplets assigned to the firewall.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("DropletIDs"),
			},
			{
				Name: "tags",
				Type: schema.TypeStringArray,
			},
			{
				Name:        "created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the firewall was created.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_firewall_inbound_rules",
				Description: "InboundRule represents a DigitalOcean Firewall inbound rule.",
				Resolver:    fetchFirewallInboundRules,
				Columns: []schema.Column{
					{
						Name:        "firewall_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_firewalls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "protocol",
						Type: schema.TypeString,
					},
					{
						Name: "port_range",
						Type: schema.TypeString,
					},
					{
						Name:     "sources_addresses",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Sources.Addresses"),
					},
					{
						Name:          "sources_tags",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Sources.Tags"),
						IgnoreInTests: true,
					},
					{
						Name:          "sources_droplet_ids",
						Type:          schema.TypeIntArray,
						Resolver:      schema.PathResolver("Sources.DropletIDs"),
						IgnoreInTests: true,
					},
					{
						Name:          "sources_load_balancer_uid_s",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Sources.LoadBalancerUIDs"),
						IgnoreInTests: true,
					},
					{
						Name:          "sources_kubernetes_ids",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Sources.KubernetesIDs"),
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:        "digitalocean_firewall_outbound_rules",
				Description: "OutboundRule represents a DigitalOcean Firewall outbound rule.",
				Resolver:    fetchFirewallOutboundRules,
				Columns: []schema.Column{
					{
						Name:        "firewall_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_firewalls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "protocol",
						Type: schema.TypeString,
					},
					{
						Name: "port_range",
						Type: schema.TypeString,
					},
					{
						Name:     "destinations_addresses",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Destinations.Addresses"),
					},
					{
						Name:          "destinations_tags",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Destinations.Tags"),
						IgnoreInTests: true,
					},
					{
						Name:          "destinations_droplet_ids",
						Type:          schema.TypeIntArray,
						Resolver:      schema.PathResolver("Destinations.DropletIDs"),
						IgnoreInTests: true,
					},
					{
						Name:          "destinations_load_balancer_uid_s",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Destinations.LoadBalancerUIDs"),
						IgnoreInTests: true,
					},
					{
						Name:          "destinations_kubernetes_ids",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("Destinations.KubernetesIDs"),
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "digitalocean_firewall_pending_changes",
				Description:   "PendingChange represents a DigitalOcean Firewall status details.",
				Resolver:      fetchFirewallPendingChanges,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "firewall_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_firewalls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "droplet_id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("DropletID"),
					},
					{
						Name: "removing",
						Type: schema.TypeBool,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:        "digitalocean_firewall_droplets",
				Description: "IDs of the Droplets assigned to the firewall",
				Resolver:    fetchFirewallDroplets,
				Columns: []schema.Column{
					{
						Name:        "firewall_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_firewalls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "droplet_id",
						Description: "Unique identifier of Droplet assigned to the firewall.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "firewall_id",
						Description: "The unique identifier for the firewall.",
						Type:        schema.TypeUUID,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFirewalls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		firewalls, resp, err := svc.DoClient.Firewalls.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- firewalls
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
func fetchFirewallInboundRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fw := parent.Item.(godo.Firewall)
	res <- fw.InboundRules
	return nil
}

func fetchFirewallOutboundRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fw := parent.Item.(godo.Firewall)
	res <- fw.OutboundRules
	return nil
}
func fetchFirewallPendingChanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fw := parent.Item.(godo.Firewall)
	res <- fw.PendingChanges
	return nil
}
func fetchFirewallDroplets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fw := parent.Item.(godo.Firewall)
	if fw.DropletIDs == nil {
		return nil
	}
	vd := make([]firewallDroplets, len(fw.DropletIDs))
	for i, dropletId := range fw.DropletIDs {
		vd[i] = firewallDroplets{
			DropletId:  dropletId,
			FirewallId: fw.ID,
		}
	}
	res <- vd
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type firewallDroplets struct {
	DropletId  int
	FirewallId string
}
