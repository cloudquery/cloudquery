package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeAddresses() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_addresses",
		Description: "Addresses for GFE-based external HTTP(S) load balancers.",
		Resolver:    fetchComputeAddresses,
		Multiplex:   client.ProjectMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "address",
				Description: "The static IP address represented by this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "address_type",
				Description: "The type of address to reserve, either INTERNAL or EXTERNAL If unspecified, defaults to EXTERNAL",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this field when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "ip_version",
				Description: "The IP version that will be used by this address Valid options are IPV4 or IPV6 This can only be specified for a global address",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#address for addresses",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "The URL of the network in which to reserve the address This field can only be used with INTERNAL type with the VPC_PEERING purpose",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_tier",
				Description: "This signifies the networking tier used for configuring this address",
				Type:        schema.TypeString,
			},
			{
				Name:        "prefix_length",
				Description: "The prefix length if the resource represents an IP range",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "purpose",
				Description: "The purpose of this resource, which can be one of the following values: - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources - `DNS_RESOLVER` for a DNS resolver address in a subnetwork - `VPC_PEERING` for addresses that are reserved for VPC peer networks - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration These addresses are regional resources",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "The URL of the region where a regional address resides For regional addresses, you must specify the region as a path parameter in the HTTP request URL This field is not applicable to global addresses",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the address, which can be one of RESERVING, RESERVED, or IN_USE An address that is RESERVING is currently in the process of being reserved A RESERVED address is currently reserved and available to use An IN_USE address is currently being used by another resource and is not available",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnetwork",
				Description: "The URL of the subnetwork in which to reserve the address If an IP address is specified, it must be within the subnetwork's IP range This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose",
				Type:        schema.TypeString,
			},
			{
				Name:        "users",
				Description: "The URLs of the resources that are using this address",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Addresses.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var addresses []*compute.Address
		for _, items := range output.Items {
			addresses = append(addresses, items.Addresses...)
		}
		res <- addresses

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
