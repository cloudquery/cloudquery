// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ForwardingRules() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_forwarding_rules",
		Resolver:  fetchForwardingRules,
		Multiplex: client.ProjectMultiplex,
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
				Name:     "all_ports",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllPorts"),
			},
			{
				Name:     "allow_global_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowGlobalAccess"),
			},
			{
				Name:     "backend_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackendService"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "ip_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpVersion"),
			},
			{
				Name:     "is_mirroring_collector",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMirroringCollector"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "label_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelFingerprint"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "load_balancing_scheme",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoadBalancingScheme"),
			},
			{
				Name:     "metadata_filters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataFilters"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "network_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkTier"),
			},
			{
				Name:     "no_automate_dns_zone",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("NoAutomateDnsZone"),
			},
			{
				Name:     "port_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PortRange"),
			},
			{
				Name:     "ports",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Ports"),
			},
			{
				Name:     "psc_connection_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PscConnectionId"),
			},
			{
				Name:     "psc_connection_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PscConnectionStatus"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_directory_registrations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceDirectoryRegistrations"),
			},
			{
				Name:     "service_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceLabel"),
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
			},
			{
				Name:     "subnetwork",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subnetwork"),
			},
			{
				Name:     "target",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Target"),
			},
		},
	}
}

func fetchForwardingRules(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListForwardingRulesRequest{}
	it := c.Services.ComputeForwardingRulesClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.ForwardingRules

	}
	return nil
}
