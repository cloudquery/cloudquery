// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func ComputeNetworks() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_networks",
		Resolver:  fetchComputeNetworks,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name: "i_pv_4_range",
				Type: schema.TypeString,
			},
			{
				Name: "auto_create_subnetworks",
				Type: schema.TypeBool,
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
				Name: "enable_ula_internal_ipv_6",
				Type: schema.TypeBool,
			},
			{
				Name: "firewall_policy",
				Type: schema.TypeString,
			},
			{
				Name: "gateway_i_pv_4",
				Type: schema.TypeString,
			},
			{
				Name: "id",
				Type: schema.TypeInt,
			},
			{
				Name: "internal_ipv_6_range",
				Type: schema.TypeString,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "mtu",
				Type: schema.TypeInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network_firewall_policy_enforcement_order",
				Type: schema.TypeString,
			},
			{
				Name: "peerings",
				Type: schema.TypeJSON,
			},
			{
				Name: "routing_config",
				Type: schema.TypeJSON,
			},
			{
				Name: "self_link_with_id",
				Type: schema.TypeString,
			},
			{
				Name: "subnetworks",
				Type: schema.TypeStringArray,
			},
			{
				Name: "server_response",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchComputeNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Networks.List(c.ProjectId).PageToken(nextPageToken).Do()
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
