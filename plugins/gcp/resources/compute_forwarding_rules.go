package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeForwardingRules() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_forwarding_rules",
		Resolver:     fetchComputeForwardingRules,
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
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPAddress"),
			},
			{
				Name:     "ip_protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPProtocol"),
			},
			{
				Name: "all_ports",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_global_access",
				Type: schema.TypeBool,
			},
			{
				Name: "backend_service",
				Type: schema.TypeString,
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
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "ip_version",
				Type: schema.TypeString,
			},
			{
				Name: "is_mirroring_collector",
				Type: schema.TypeBool,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "label_fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "load_balancing_scheme",
				Type: schema.TypeString,
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
				Name: "network_tier",
				Type: schema.TypeString,
			},
			{
				Name: "port_range",
				Type: schema.TypeString,
			},
			{
				Name: "ports",
				Type: schema.TypeStringArray,
			},
			{
				Name: "psc_connection_id",
				Type: schema.TypeBigInt,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "service_label",
				Type: schema.TypeString,
			},
			{
				Name: "service_name",
				Type: schema.TypeString,
			},
			{
				Name: "subnetwork",
				Type: schema.TypeString,
			},
			{
				Name: "target",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeForwardingRules(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.ForwardingRules.AggregatedList(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var forwardingRules []*compute.ForwardingRule
		for _, item := range output.Items {
			forwardingRules = append(forwardingRules, item.ForwardingRules...)
		}
		res <- forwardingRules
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
