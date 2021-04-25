package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeInterconnects() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_interconnects",
		Resolver:     fetchComputeInterconnects,
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
				Name: "admin_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "customer_name",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "google_ip_address",
				Type: schema.TypeString,
			},
			{
				Name: "google_reference_id",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "interconnect_attachments",
				Type: schema.TypeStringArray,
			},
			{
				Name: "interconnect_type",
				Type: schema.TypeString,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "link_type",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "noc_contact_email",
				Type: schema.TypeString,
			},
			{
				Name: "operational_status",
				Type: schema.TypeString,
			},
			{
				Name: "peer_ip_address",
				Type: schema.TypeString,
			},
			{
				Name: "provisioned_link_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "requested_link_count",
				Type: schema.TypeBigInt,
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
				Name:     "gcp_compute_interconnect_circuit_infos",
				Resolver: fetchComputeInterconnectCircuitInfos,
				Columns: []schema.Column{
					{
						Name:     "interconnect_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "customer_demarc_id",
						Type: schema.TypeString,
					},
					{
						Name: "google_circuit_id",
						Type: schema.TypeString,
					},
					{
						Name: "google_demarc_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_compute_interconnect_expected_outages",
				Resolver: fetchComputeInterconnectExpectedOutages,
				Columns: []schema.Column{
					{
						Name:     "interconnect_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "affected_circuits",
						Type: schema.TypeStringArray,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "end_time",
						Type: schema.TypeBigInt,
					},
					{
						Name: "issue_type",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "source",
						Type: schema.TypeString,
					},
					{
						Name: "start_time",
						Type: schema.TypeBigInt,
					},
					{
						Name: "state",
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
func fetchComputeInterconnects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Interconnects.List(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeInterconnectCircuitInfos(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Interconnect)
	res <- r.CircuitInfos
	return nil
}
func fetchComputeInterconnectExpectedOutages(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Interconnect)
	res <- r.ExpectedOutages
	return nil
}
