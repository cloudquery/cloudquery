package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeFirewalls() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_firewalls",
		Description: "Represents a Firewall Rule resource  Firewall rules allow or deny ingress traffic to, and egress traffic from your instances For more information, read Firewall rules",
		Resolver:    fetchComputeFirewalls,

		Multiplex: client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
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
				Name:          "destination_ranges",
				Description:   "If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges These ranges must be expressed in CIDR format Only IPv4 is supported",
				IgnoreInTests: true,
				Type:          schema.TypeStringArray,
			},
			{
				Name:        "direction",
				Description: "Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS` The default is `INGRESS` For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields",
				Type:        schema.TypeString,
			},
			{
				Name:        "disabled",
				Description: "Denotes whether the firewall rule is disabled When set to true, the firewall rule is not enforced and the network behaves as if it did not exist If this is unspecified, the firewall rule will be enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#firewall for firewall rules",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_config_enable",
				Description: "This field denotes whether to enable logging for a particular firewall rule",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LogConfig.Enable"),
			},
			{
				Name:        "log_config_metadata",
				Description: "This field can only be specified for a particular firewall rule if logging is enabled for that rule This field denotes whether to include or exclude metadata for firewall logs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogConfig.Metadata"),
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "URL of the network resource for this firewall rule If not specified when creating a firewall rule, the default network is used: global/networks/default If you choose to specify this field, you can specify the network as a full or partial URL For example, the following are all valid URLs:  - https://wwwgoogleapis",
				Type:        schema.TypeString,
			},
			{
				Name:        "priority",
				Description: "Priority for this rule This is an integer between `0` and `65535`, both inclusive The default value is `1000` Relative priorities determine which rule takes effect if multiple rules apply Lower values indicate higher priority For example, a rule with priority `0` has higher precedence than a rule with priority `1` DENY rules take precedence over ALLOW rules if they have equal priority Note that VPC networks have implied rules with a priority of `65535` To avoid conflicts with the implied rules, use a priority number less than `65535`",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_ranges",
				Description: "If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges These ranges must be expressed in CIDR format One or both of sourceRanges and sourceTags may be set If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field The connection does not need to match both fields for the rule to apply Only IPv4 is supported",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "source_service_accounts",
				Description:   "If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address sourceRanges can be set at the same time as sourceServiceAccounts If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount The connection does not need to match both fields for the firewall to apply sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags",
				IgnoreInTests: true,
				Type:          schema.TypeStringArray,
			},
			{
				Name:        "source_tags",
				Description: "If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network Because tags are associated with instances, not IP addresses One or both of sourceRanges and sourceTags may be set If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field The connection does not need to match both fields for the firewall to apply",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "target_service_accounts",
				Description:   "A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[] targetServiceAccounts cannot be used at the same time as targetTags or sourceTags If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "target_tags",
				Description: "A list of tags that controls which instances the firewall rule applies to If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags If no targetTags are specified, the firewall rule applies to all instances on the specified network",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_firewall_allowed",
				Resolver: fetchComputeFirewallAllowed,
				Columns: []schema.Column{
					{
						Name:        "firewall_cq_id",
						Description: "Unique ID of gcp_compute_firewalls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "firewall_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "ip_protocol",
						Description: "The IP protocol to which this rule applies The protocol type is required when creating a firewall rule This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp) or the IP protocol number",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPProtocol"),
					},
					{
						Name:        "ports",
						Description: "An optional list of ports to which this rule applies This field is only applicable for the UDP or TCP protocol Each entry must be either an integer or a range If not specified, this rule applies to connections through any port  Example inputs include: [\"22\"], [\"80\",\"443\"], and [\"12345-12349\"]",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "gcp_compute_firewall_denied",
				Resolver: fetchComputeFirewallDenied,
				Columns: []schema.Column{
					{
						Name:        "firewall_cq_id",
						Description: "Unique ID of gcp_compute_firewalls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "firewall_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "ip_protocol",
						Description: "The IP protocol to which this rule applies The protocol type is required when creating a firewall rule This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp) or the IP protocol number",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPProtocol"),
					},
					{
						Name:        "ports",
						Description: "An optional list of ports to which this rule applies This field is only applicable for the UDP or TCP protocol Each entry must be either an integer or a range If not specified, this rule applies to connections through any port  Example inputs include: [\"22\"], [\"80\",\"443\"], and [\"12345-12349\"]",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeFirewalls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Firewalls.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeFirewallAllowed(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Firewall)
	res <- r.Allowed
	return nil
}
func fetchComputeFirewallDenied(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Firewall)
	res <- r.Denied
	return nil
}
