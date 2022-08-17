package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeSubnetworks() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_subnetworks",
		Description: "Represents a Subnetwork resource  A subnetwork (also known as a subnet) is a logical partition of a Virtual Private Cloud network with one primary IP range and zero or more secondary IP ranges",
		Resolver:    fetchComputeSubnetworks,
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
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource This field can be set only at resource creation time",
				Type:        schema.TypeString,
			},
			{
				Name:        "enable_flow_logs",
				Description: "Whether to enable flow logging for this subnetwork If this field is not explicitly set, it will not appear in get listings If not set the default behavior is to disable flow logging This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER",
				Type:        schema.TypeBool,
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource A hash of the contents stored in this object This field is used in optimistic locking This field will be ignored when inserting a Subnetwork An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve a Subnetwork",
				Type:        schema.TypeString,
			},
			{
				Name:        "gateway_address",
				Description: "The gateway address for default routes to reach destination addresses outside this subnetwork",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "ip_cidr_range",
				Description: "The range of internal addresses that are owned by this subnetwork Provide this property when you create the subnetwork For example, 10000/8 or 1006400/10 Ranges must be unique and non-overlapping within a network Only IPv4 is supported This field is set at resource creation time The range can be any range listed in the Valid ranges list The range can be expanded after creation using expandIpCidrRange",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_cidr_range",
				Description: "The range of internal IPv6 addresses that are owned by this subnetwork",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#subnetwork for Subnetwork resources",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_config_aggregation_interval",
				Description: "Can only be specified if VPC flow logging for this subnetwork is enabled Toggles the aggregation interval for collecting flow logs Increasing the interval time will reduce the amount of generated flow logs for long lasting connections Default is an interval of 5 seconds per connection",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogConfig.AggregationInterval"),
			},
			{
				Name:        "log_config_enable",
				Description: "Whether to enable flow logging for this subnetwork If this field is not explicitly set, it will not appear in get listings If not set the default behavior is to disable flow logging",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LogConfig.Enable"),
			},
			{
				Name:        "log_config_filter_expr",
				Description: "Can only be specified if VPC flow logs for this subnetwork is enabled Export filter used to define which VPC flow logs should be logged",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogConfig.FilterExpr"),
			},
			{
				Name:        "log_config_flow_sampling",
				Description: "Can only be specified if VPC flow logging for this subnetwork is enabled The value of the field must be in [0, 1] Set the sampling rate of VPC flow logs within the subnetwork where 10 means all collected logs are reported and 00 means no logs are reported Default is 05, which means half of all collected logs are reported",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("LogConfig.FlowSampling"),
			},
			{
				Name:        "log_config_metadata",
				Description: "Can only be specified if VPC flow logs for this subnetwork is enabled Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs Default is EXCLUDE_ALL_METADATA",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogConfig.Metadata"),
			},
			{
				Name:          "log_config_metadata_fields",
				Description:   "Can only be specified if VPC flow logs for this subnetwork is enabled and \"metadata\" was set to CUSTOM_METADATA",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("LogConfig.MetadataFields"),
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork Only networks that are in the distributed mode can have subnetworks This field can be set only at resource creation time",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ip_google_access",
				Description: "Whether the VMs in this subnet can access Google services without assigned external IP addresses This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess",
				Type:        schema.TypeBool,
			},
			{
				Name:        "private_ipv6_google_access",
				Description: "The private IPv6 google access type for the VMs in this subnet This is an expanded field of enablePrivateV6Access If both fields are set, privateIpv6GoogleAccess will take priority  This field can be both set at resource creation time and updated using patch",
				Type:        schema.TypeString,
			},
			{
				Name:        "purpose",
				Description: "The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing If unspecified, the purpose defaults to PRIVATE_RFC_1918 The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCE",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the Subnetwork resides This field can be set only at resource creation time",
				Type:        schema.TypeString,
			},
			{
				Name:        "role",
				Description: "The role of subnetwork Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER The value can be set to ACTIVE or BACKUP An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining This field can be updated with a patch request",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_subnetwork_secondary_ip_ranges",
				Description: "Represents a secondary IP range of a subnetwork",
				Resolver:    fetchComputeSubnetworkSecondaryIpRanges,
				Columns: []schema.Column{
					{
						Name:        "subnetwork_cq_id",
						Description: "Unique ID of gcp_compute_subnetworks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "subnetwork_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "ip_cidr_range",
						Description: "The range of IP addresses belonging to this subnetwork secondary range Provide this property when you create the subnetwork Ranges must be unique and non-overlapping with all primary and secondary IP ranges within a network Only IPv4 is supported The range can be any range listed in the Valid ranges list",
						Type:        schema.TypeString,
					},
					{
						Name:        "range_name",
						Description: "The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance The name must be 1-63 characters long, and comply with RFC1035 The name must be unique within the subnetwork",
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
func fetchComputeSubnetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Subnetworks.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
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
func fetchComputeSubnetworkSecondaryIpRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Subnetwork)
	res <- r.SecondaryIpRanges
	return nil
}
