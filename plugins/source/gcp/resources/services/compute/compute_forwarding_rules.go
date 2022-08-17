package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeForwardingRules() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_forwarding_rules",
		Description: "Represents a Forwarding Rule resource  Forwarding rule resources in GCP can be either regional or global.",
		Resolver:    fetchComputeForwardingRules,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "ip_address",
				Description: "IP address that this forwarding rule serves When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IPAddress"),
			},
			{
				Name:        "ip_protocol",
				Description: "The IP protocol to which this rule applies  For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP  The valid IP protocols are different for different load balancing products: - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IPProtocol"),
			},
			{
				Name:        "all_ports",
				Description: "This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance This field cannot be used with port or portRange fields  When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule",
				Type:        schema.TypeBool,
			},
			{
				Name:        "allow_global_access",
				Description: "This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance If the field is set to TRUE, clients can access ILB from all regions Otherwise only allows access from clients in the same region as the internal load balancer",
				Type:        schema.TypeBool,
			},
			{
				Name:        "backend_service",
				Description: "Identifies the backend service to which the forwarding rule sends traffic Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource A hash of the contents stored in this object",
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
				Description: "The IP Version that will be used by this forwarding rule Valid options are IPV4 or IPV6 This can only be specified for an external global forwarding rule",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_mirroring_collector",
				Description: "Indicates whether or not this load balancer can be used as a collector for packet mirroring To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL",
				Type:        schema.TypeBool,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#forwardingRule for Forwarding Rule resources",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_fingerprint",
				Description: "A fingerprint for the labels being applied to this resource",
				Type:        schema.TypeString,
			},
			{
				Name:          "labels",
				Description:   "Labels for this resource",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "load_balancing_scheme",
				Description: "Specifies the forwarding rule type  - EXTERNAL is used for: - Classic Cloud VPN gateways - Protocol forwarding to VMs from an external IP address - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing - INTERNAL is used for: - Protocol forwarding to VMs from an internal IP address - Internal TCP/UDP Load Balancing - INTERNAL_MANAGED is used for: - Internal HTTP(S) Load Balancing - INTERNAL_SELF_MANAGED is used for: - Traffic Director  For more information about forwarding rules, refer to Forwarding rule concepts",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "This field is not used for external load balancing  For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule If this field is not specified, the default network will be used  For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_tier",
				Description: "This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD  For regional ForwardingRule, the valid values are PREMIUM and STANDARD For GlobalForwardingRule, the valid value is PREMIUM  If this field is not specified, it is assumed to be PREMIUM If IPAddress is specified, this value must be equal to the networkTier of the Address",
				Type:        schema.TypeString,
			},
			{
				Name:        "port_range",
				Description: "This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP  Packets addressed to ports in the specified range will be forwarded to target or  backend_service You can only use one of ports, port_range, or allPorts The three are mutually exclusive Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges",
				Type:        schema.TypeString,
			},
			{
				Name:          "ports",
				Description:   "The ports field is only supported when the forwarding rule references a backend_service directly Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing Only packets addressed to the specified list of ports are forwarded to backends  You can only use one of ports and port_range, or allPorts The three are mutually exclusive  You can specify a list of up to five ports, which can be non-contiguous  For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports  For more information, see Port specifications (/load-balancing/docs/forwarding-rule-concepts#port_specifications)",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "psc_connection_id",
				Description: "The PSC connection id of the PSC Forwarding Rule",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "region",
				Description: "URL of the region where the regional forwarding rule resides",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_label",
				Description: "An optional prefix to the service name for this Forwarding Rule",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_name",
				Description: "The internal fully qualified service name for this Forwarding Rule  This field is only used for internal load balancing",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnetwork",
				Description: "This field is only used for internal load balancing",
				Type:        schema.TypeString,
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
func fetchComputeForwardingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.ForwardingRules.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
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
