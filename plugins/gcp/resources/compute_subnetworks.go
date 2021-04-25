package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeSubnetworks() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_subnetworks",
		Resolver:     fetchComputeSubnetworks,
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
				Name: "enable_flow_logs",
				Type: schema.TypeBool,
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "gateway_address",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "ip_cidr_range",
				Type: schema.TypeString,
			},
			{
				Name: "ipv6_cidr_range",
				Type: schema.TypeString,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name:     "log_config_aggregation_interval",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogConfig.AggregationInterval"),
			},
			{
				Name:     "log_config_enable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("LogConfig.Enable"),
			},
			{
				Name:     "log_config_filter_expr",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogConfig.FilterExpr"),
			},
			{
				Name:     "log_config_flow_sampling",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("LogConfig.FlowSampling"),
			},
			{
				Name:     "log_config_metadata",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogConfig.Metadata"),
			},
			{
				Name:     "log_config_metadata_fields",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("LogConfig.MetadataFields"),
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
				Name: "private_ip_google_access",
				Type: schema.TypeBool,
			},
			{
				Name: "private_ipv6_google_access",
				Type: schema.TypeString,
			},
			{
				Name: "purpose",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "role",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_subnetwork_secondary_ip_ranges",
				Resolver: fetchComputeSubnetworkSecondaryIpRanges,
				Columns: []schema.Column{
					{
						Name:     "subnetwork_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "ip_cidr_range",
						Type: schema.TypeString,
					},
					{
						Name: "range_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeSubnetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Subnetworks.AggregatedList(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var subnetworks []*compute.Subnetwork
		for _, scopedNetworkList := range output.Items {
			subnetworks = append(subnetworks, scopedNetworkList.Subnetworks...)
		}
		res <- subnetworks

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeSubnetworkSecondaryIpRanges(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Subnetwork)
	res <- r.SecondaryIpRanges
	return nil
}
