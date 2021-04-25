package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeFirewalls() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_firewalls",
		Resolver:     fetchComputeFirewalls,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "destination_ranges",
				Type: schema.TypeStringArray,
			},
			{
				Name: "direction",
				Type: schema.TypeString,
			},
			{
				Name: "disabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name:     "log_config_enable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("LogConfig.Enable"),
			},
			{
				Name:     "log_config_metadata",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogConfig.Metadata"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network",
				Type: schema.TypeString,
			},
			{
				Name: "priority",
				Type: schema.TypeBigInt,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "source_ranges",
				Type: schema.TypeStringArray,
			},
			{
				Name: "source_service_accounts",
				Type: schema.TypeStringArray,
			},
			{
				Name: "source_tags",
				Type: schema.TypeStringArray,
			},
			{
				Name: "target_service_accounts",
				Type: schema.TypeStringArray,
			},
			{
				Name: "target_tags",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_firewall_allowed",
				Resolver: fetchComputeFirewallAllowed,
				Columns: []schema.Column{
					{
						Name:     "firewall_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "ip_protocol",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("IPProtocol"),
					},
					{
						Name: "ports",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "gcp_compute_firewall_denied",
				Resolver: fetchComputeFirewallDenied,
				Columns: []schema.Column{
					{
						Name:     "firewall_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "ip_protocol",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("IPProtocol"),
					},
					{
						Name: "ports",
						Type: schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeFirewalls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, _ chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Firewalls.List(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeFirewallAllowed(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Firewall)
	res <- r.Allowed
	return nil
}
func fetchComputeFirewallDenied(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Firewall)
	res <- r.Denied
	return nil
}
