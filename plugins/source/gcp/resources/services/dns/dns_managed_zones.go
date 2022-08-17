package dns

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/dns/v1"
)

func DNSManagedZones() *schema.Table {
	return &schema.Table{
		Name:        "gcp_dns_managed_zones",
		Description: "A zone is a subtree of the DNS namespace under one administrative responsibility A ManagedZone is a resource that represents a DNS zone hosted by the Cloud DNS service",
		Resolver:    fetchDnsManagedZones,
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
				Name:        "creation_time",
				Description: "The time that this resource was created on the server This is in RFC3339 text format Output only",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A mutable string of at most 1024 characters associated with this resource for the user's convenience Has no effect on the managed zone's function",
				Type:        schema.TypeString,
			},
			{
				Name:        "dns_name",
				Description: "The DNS name of this managed zone, for instance \"examplecom\"",
				Type:        schema.TypeString,
			},
			{
				Name:     "dnssec_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnssecConfig.Kind"),
			},
			{
				Name:        "dnssec_config_non_existence",
				Description: "Specifies the mechanism for authenticated denial-of-existence responses Can only be changed while the state is OFF",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DnssecConfig.NonExistence"),
			},
			{
				Name:        "dnssec_config_state",
				Description: "Specifies whether DNSSEC is enabled, and what mode it is in  Possible values:   \"off\" - DNSSEC is disabled; the zone is not signed   \"on\" - DNSSEC is enabled; the zone is signed and fully managed   \"transfer\" - DNSSEC is enabled, but in a \"transfer\" mode",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DnssecConfig.State"),
			},
			{
				Name:     "forwarding_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ForwardingConfig.Kind"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "The resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "User assigned labels for this resource",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "name",
				Description: "User assigned name for this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "name_server_set",
				Description: "specifies the NameServerSet for this ManagedZone",
				Type:        schema.TypeString,
			},
			{
				Name: "name_servers",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "peering_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringConfig.Kind"),
			},
			{
				Name:        "peering_config_target_network_deactivate_time",
				Description: "The time at which the zone was deactivated, in RFC 3339 date-time format An empty string indicates that the peering connection is active The producer network can deactivate a zone The zone is automatically deactivated if the producer network that the zone targeted is deleted Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PeeringConfig.TargetNetwork.DeactivateTime"),
			},
			{
				Name:     "peering_config_target_network_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PeeringConfig.TargetNetwork.Kind"),
			},
			{
				Name:        "peering_config_target_network_network_url",
				Description: "The fully qualified URL of the VPC network to forward queries to This should be formatted like https://wwwgoogleapis",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PeeringConfig.TargetNetwork.NetworkUrl"),
			},
			{
				Name:     "private_visibility_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateVisibilityConfig.Kind"),
			},
			{
				Name:     "reverse_lookup_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReverseLookupConfig.Kind"),
			},
			{
				Name:     "service_directory_config_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceDirectoryConfig.Kind"),
			},
			{
				Name:        "service_directory_config_namespace_deletion_time",
				Description: "The time that the namespace backing this zone was deleted; an empty string if it still exists This is in RFC3339 text format Output only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceDirectoryConfig.Namespace.DeletionTime"),
			},
			{
				Name:     "service_directory_config_namespace_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceDirectoryConfig.Namespace.Kind"),
			},
			{
				Name:        "service_directory_config_namespace_namespace_url",
				Description: "The fully qualified URL of the namespace associated with the zone Format must be https://servicedirectorygoogleapis",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceDirectoryConfig.Namespace.NamespaceUrl"),
			},
			{
				Name:        "visibility",
				Description: "The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_dns_managed_zone_dnssec_config_default_key_specs",
				Description: "Parameters for DnsKey key generation Used for generating initial keys for a new ManagedZone and as default when adding a new DnsKey",
				Resolver:    fetchDnsManagedZoneDnssecConfigDefaultKeySpecs,
				Columns: []schema.Column{
					{
						Name:        "managed_zone_cq_id",
						Description: "Unique ID of gcp_dns_managed_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "managed_zone_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "algorithm",
						Description: "String mnemonic specifying the DNSSEC algorithm of this key",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_length",
						Description: "Length of the keys in bits",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "key_type",
						Description: "Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK) Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets",
						Type:        schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_dns_managed_zone_forwarding_config_target_name_servers",
				Resolver: fetchDnsManagedZoneForwardingConfigTargetNameServers,
				Columns: []schema.Column{
					{
						Name:        "managed_zone_cq_id",
						Description: "Unique ID of gcp_dns_managed_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "managed_zone_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "forwarding_path",
						Description: "Forwarding path for this NameServerTarget If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on IP address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target  Possible values:   \"default\" - Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses forward to the target through the VPC and non-RFC1918 addresses forward to the target through the internet   \"private\" - Cloud DNS always forwards to this target through the VPC",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv4_address",
						Description: "IPv4 address of a target name server",
						Type:        schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_dns_managed_zone_private_visibility_config_networks",
				Resolver: fetchDnsManagedZonePrivateVisibilityConfigNetworks,
				Columns: []schema.Column{
					{
						Name:        "managed_zone_cq_id",
						Description: "Unique ID of gcp_dns_managed_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "managed_zone_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name:        "network_url",
						Description: "The fully qualified URL of the VPC network to bind to Format this URL like https://wwwgoogleapis",
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
func fetchDnsManagedZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Dns.ManagedZones.
			List(c.ProjectId).
			PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.ManagedZones
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchDnsManagedZoneDnssecConfigDefaultKeySpecs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*dns.ManagedZone)
	if p.DnssecConfig == nil {
		return nil
	}

	res <- p.DnssecConfig.DefaultKeySpecs
	return nil
}
func fetchDnsManagedZoneForwardingConfigTargetNameServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*dns.ManagedZone)
	if p.ForwardingConfig == nil {
		return nil
	}

	res <- p.ForwardingConfig.TargetNameServers
	return nil
}
func fetchDnsManagedZonePrivateVisibilityConfigNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*dns.ManagedZone)
	if p.PrivateVisibilityConfig == nil {
		return nil
	}

	res <- p.PrivateVisibilityConfig.Networks
	return nil
}
