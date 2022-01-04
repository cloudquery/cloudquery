package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_load_balancers",
		Description:  "LoadBalancer represents a DigitalOcean load balancer configuration. Tags can only be provided upon the creation of a Load Balancer.",
		Resolver:     fetchLoadBalancers,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique ID that can be used to identify and reference a load balancer.",
				Type:        schema.TypeUUID,
				Resolver:    schema.UUIDResolver("ID"),
			},
			{
				Name:        "name",
				Description: "A human-readable name for a load balancer instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ip",
				Description: "An attribute containing the public-facing IP address of the load balancer.",
				Type:        schema.TypeInet,
				Resolver:    client.IPAddressResolver("IP"),
			},
			{
				Name:        "size",
				Description: "The size of the load balancer. The available sizes are lb-small, lb-medium, or lb-large. You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SizeSlug"),
			},
			{
				Name:        "algorithm",
				Description: "The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either `round_robin` or `least_connections`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created",
				Description: "A time value given in ISO8601 combined date and time format that represents when the load balancer was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_check_protocol",
				Description: "The protocol used for health checks sent to the backend Droplets. The possible values are `http`, `https`, or `tcp`.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheck.Protocol"),
			},
			{
				Name:        "health_check_port",
				Description: "An integer representing the port on the backend Droplets on which the health check will attempt a connection.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("HealthCheck.Port"),
			},
			{
				Name:        "health_check_path",
				Description: "The path on the backend Droplets to which the load balancer instance will send a request.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheck.Path"),
			},
			{
				Name:        "health_check_check_interval_seconds",
				Description: "The number of seconds between between two consecutive health checks.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("HealthCheck.CheckIntervalSeconds"),
			},
			{
				Name:        "health_check_response_timeout_seconds",
				Description: "The number of seconds the load balancer instance will wait for a response until marking a health check as failed.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("HealthCheck.ResponseTimeoutSeconds"),
			},
			{
				Name:        "health_check_healthy_threshold",
				Description: "The number of times a health check must pass for a backend Droplet to be marked \"healthy\" and be re-added to the pool.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("HealthCheck.HealthyThreshold"),
			},
			{
				Name:        "health_check_unhealthy_threshold",
				Description: "The number of times a health check must fail for a backend Droplet to be marked \"unhealthy\" and be removed from the pool.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("HealthCheck.UnhealthyThreshold"),
			},
			{
				Name:        "sticky_sessions_type",
				Description: "An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StickySessions.Type"),
			},
			{
				Name:        "sticky_sessions_cookie_name",
				Description: "The name of the cookie sent to the client. This attribute is only returned when using `cookies` for the sticky sessions type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StickySessions.CookieName"),
			},
			{
				Name:        "sticky_sessions_cookie_ttl_seconds",
				Description: "The number of seconds until the cookie set by the load balancer expires. This attribute is only returned when using `cookies` for the sticky sessions type.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("StickySessions.CookieTtlSeconds"),
			},
			{
				Name:        "region_slug",
				Description: "A human-readable string that is used as a unique identifier for each region.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Region.Slug"),
			},
			{
				Name:        "region_name",
				Description: "The display name of the region.  This will be a full name that is used in the control panel and other interfaces.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Region.Name"),
			},
			{
				Name:        "region_sizes",
				Description: "This attribute is set to an array which contains the identifying slugs for the sizes available in this region.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Region.Sizes"),
			},
			{
				Name:        "region_available",
				Description: "This is a boolean value that represents whether new Droplets can be created in this region.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Region.Available"),
			},
			{
				Name:        "region_features",
				Description: "This attribute is set to an array which contains features available in this region",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Region.Features"),
			},
			{
				Name:        "droplet_ids",
				Description: "An array containing the IDs of the Droplets assigned to the load balancer.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("DropletIDs"),
			},
			{
				Name:        "tag",
				Description: "The name of a Droplet tag corresponding to Droplets assigned to the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:          "tags",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "redirect_http_to_https",
				Description: "A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enable_proxy_protocol",
				Description: "A boolean value indicating whether PROXY Protocol is in use.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enable_backend_keepalive",
				Description: "A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "vpc_uuid",
				Description: "A string specifying the UUID of the VPC to which the load balancer is assigned.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VPCUUID"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_load_balancer_forwarding_rules",
				Description: "ForwardingRule represents load balancer forwarding rules.",
				Resolver:    fetchLoadBalancerForwardingRules,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entry_protocol",
						Description: "The protocol used for traffic to the load balancer. The possible values are: `http`, `https`, `http2`, or `tcp`.\n",
						Type:        schema.TypeString,
					},
					{
						Name:        "entry_port",
						Description: "An integer representing the port on which the load balancer instance will listen.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "target_protocol",
						Description: "The protocol used for traffic from the load balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, or `tcp`.\n",
						Type:        schema.TypeString,
					},
					{
						Name:        "target_port",
						Description: "An integer representing the port on the backend Droplets to which the load balancer will send traffic.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "certificate_id",
						Description: "The ID of the TLS certificate used for SSL termination if enabled.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CertificateID"),
					},
					{
						Name:        "tls_passthrough",
						Description: "A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets.",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "digitalocean_load_balancer_droplets",
				Description: "Droplets that are co-located on the same physical hardware",
				Resolver:    fetchLoadBalancerDroplets,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "droplet_id",
						Description: "Unique identifier of Droplet assigned to the load balancer.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "load_balancer_id",
						Description: "The unique identifier for the load balancer.",
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

func fetchLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		lb, resp, err := svc.DoClient.LoadBalancers.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- lb
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
func fetchLoadBalancerForwardingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lb := parent.Item.(godo.LoadBalancer)
	res <- lb.ForwardingRules
	return nil
}
func fetchLoadBalancerDroplets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lb := parent.Item.(godo.LoadBalancer)
	if lb.DropletIDs == nil {
		return nil
	}
	vd := make([]lbDroplets, len(lb.DropletIDs))
	for i, dropletId := range lb.DropletIDs {
		vd[i] = lbDroplets{
			DropletId:      dropletId,
			LoadBalancerId: lb.ID,
		}
	}
	res <- vd
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type lbDroplets struct {
	DropletId      int
	LoadBalancerId string
}
