package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeInterconnects() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_interconnects",
		Description: "Represents an Interconnect resource  An Interconnect resource is a dedicated connection between the GCP network and your on-premises network",
		Resolver:    fetchComputeInterconnects,
		Multiplex:   client.ProjectMultiplex,

		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "admin_enabled",
				Description: "Administrative status of the interconnect When this is set to true, the Interconnect is functional and can carry traffic When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it By default, the status is set to true",
				Type:        schema.TypeBool,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_name",
				Description: "Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "google_ip_address",
				Description: "IP address configured on the Google side of the Interconnect link This can be used only for ping tests",
				Type:        schema.TypeString,
			},
			{
				Name:        "google_reference_id",
				Description: "Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "interconnect_attachments",
				Description: "A list of the URLs of all InterconnectAttachments configured to use this Interconnect",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "interconnect_type",
				Description: "Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner - DEDICATED: A dedicated physical interconnection with the customer Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#interconnect for interconnects",
				Type:        schema.TypeString,
			},
			{
				Name:        "link_type",
				Description: "Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "URL of the InterconnectLocation object that represents where this connection is to be provisioned",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "noc_contact_email",
				Description: "Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications",
				Type:        schema.TypeString,
			},
			{
				Name:        "operational_status",
				Description: "The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use Attachments may be provisioned on this Interconnect - OS_UNPROVISIONED: An Interconnect that has not completed turnup No attachments may be provisioned on this Interconnect - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance No attachments may be provisioned or updated on this Interconnect",
				Type:        schema.TypeString,
			},
			{
				Name:        "peer_ip_address",
				Description: "IP address configured on the customer side of the Interconnect link The customer should configure this IP address during turnup when prompted by Google NOC This can be used only for ping tests",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioned_link_count",
				Description: "Number of links actually provisioned in this interconnect",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "requested_link_count",
				Description: "Target number of physical links in the link bundle, as requested by the customer",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use Attachments may be provisioned on this Interconnect - UNPROVISIONED: The Interconnect has not completed turnup No attachments may be provisioned on this Interconnect - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance No attachments may be provisioned or updated on this Interconnect",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_interconnect_circuit_infos",
				Description: "Describes a single physical circuit between the Customer and Google CircuitInfo objects are created by Google, so all fields are output only",
				Resolver:    fetchComputeInterconnectCircuitInfos,
				Columns: []schema.Column{
					{
						Name:        "interconnect_cq_id",
						Description: "Unique ID of gcp_compute_interconnects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "interconnect_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "customer_demarc_id",
						Description: "Customer-side demarc ID for this circuit",
						Type:        schema.TypeString,
					},
					{
						Name:        "google_circuit_id",
						Description: "Google-assigned unique ID for this circuit Assigned at circuit turn-up",
						Type:        schema.TypeString,
					},
					{
						Name:        "google_demarc_id",
						Description: "Google-side demarc ID for this circuit Assigned at circuit turn-up and provided by Google to the customer in the LOA",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_compute_interconnect_expected_outages",
				Description: "Description of a planned outage on this Interconnect",
				Resolver:    fetchComputeInterconnectExpectedOutages,
				Columns: []schema.Column{
					{
						Name:        "interconnect_cq_id",
						Description: "Unique ID of gcp_compute_interconnects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "interconnect_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "affected_circuits",
						Description: "If issue_type is IT_PARTIAL_OUTAGE, a list of the Google-side circuit IDs that will be affected",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "description",
						Description: "A description about the purpose of the outage",
						Type:        schema.TypeString,
					},
					{
						Name:        "end_time",
						Description: "Scheduled end time for the outage (milliseconds since Unix epoch)",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "issue_type",
						Description: "Form this outage is expected to take, which can take one of the following values: - OUTAGE: The Interconnect may be completely out of service for some or all of the specified window - PARTIAL_OUTAGE: Some circuits comprising the Interconnect as a whole should remain up, but with reduced bandwidth Note that the versions of this enum prefixed with \"IT_\" have been deprecated in favor of the unprefixed values",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Unique identifier for this outage notification",
						Type:        schema.TypeString,
					},
					{
						Name:        "source",
						Description: "The party that generated this notification, which can take the following value: - GOOGLE: this notification as generated by Google Note that the value of NSRC_GOOGLE has been deprecated in favor of GOOGLE",
						Type:        schema.TypeString,
					},
					{
						Name:        "start_time",
						Description: "Scheduled start time for the outage (milliseconds since Unix epoch)",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "state",
						Description: "State of this notification, which can take one of the following values: - ACTIVE: This outage notification is active The event could be in the past, present, or future See start_time and end_time for scheduling - CANCELLED: The outage associated with this notification was cancelled before the outage was due to start Note that the versions of this enum prefixed with \"NS_\" have been deprecated in favor of the unprefixed values",
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
func fetchComputeInterconnects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Interconnects.List(c.ProjectId).PageToken(nextPageToken).Do()
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
func fetchComputeInterconnectCircuitInfos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Interconnect)
	res <- r.CircuitInfos
	return nil
}
func fetchComputeInterconnectExpectedOutages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Interconnect)
	res <- r.ExpectedOutages
	return nil
}
