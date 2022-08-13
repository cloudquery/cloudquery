package dns

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/dns/v1"
)

func DNSPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_dns_policies",
		Description: "A policy is a collection of DNS rules applied to one or more Virtual Private Cloud resources",
		Resolver:    fetchDnsPolicies,
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
				Name:        "alternative_name_server_config_kind",
				Description: "alternative name server type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AlternativeNameServerConfig.Kind"),
			},
			{
				Name:        "description",
				Description: "A mutable string of at most 1024 characters associated with this resource for the user's convenience Has no effect on the policy's function",
				Type:        schema.TypeString,
			},
			{
				Name:        "enable_inbound_forwarding",
				Description: "Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enable_logging",
				Description: "Controls whether logging is enabled for the networks bound to this policy Defaults to no logging if not set",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "Unique identifier for the resource; defined by the server (output only)",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "The resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "User-assigned name for this policy",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_dns_policy_alternative_name_servers",
				Resolver: fetchDnsPolicyAlternativeNameServers,
				Columns: []schema.Column{
					{
						Name:        "policy_cq_id",
						Description: "Unique ID of gcp_dns_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "policy_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "forwarding_path",
						Description: "Forwarding path for this TargetNameServer If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target  Possible values:   \"default\" - Cloud DNS makes forwarding decision based on IP address ranges; that is, RFC1918 addresses forward to the target through the VPC and non-RFC1918 addresses forward to the target through the internet   \"private\" - Cloud DNS always forwards to this target through the VPC",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv4_address",
						Description: "IPv4 address to forward to",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "The resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_dns_policy_networks",
				Resolver: fetchDnsPolicyNetworks,
				Columns: []schema.Column{
					{
						Name:        "policy_cq_id",
						Description: "Unique ID of gcp_dns_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "policy_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "kind",
						Description: "The resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_url",
						Description: "The fully qualified URL of the VPC network to bind to This should be formatted like https://wwwgoogleapis",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDnsPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Dns.Policies.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Policies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchDnsPolicyAlternativeNameServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*dns.Policy)
	if p.AlternativeNameServerConfig == nil {
		return nil
	}

	res <- p.AlternativeNameServerConfig.TargetNameServers
	return nil
}
func fetchDnsPolicyNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*dns.Policy)
	res <- p.Networks
	return nil
}
