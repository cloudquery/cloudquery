// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func ComputeProjects() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_projects",
		Resolver:  fetchComputeProjects,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name: "common_instance_metadata",
				Type: schema.TypeJSON,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "default_network_tier",
				Type: schema.TypeString,
			},
			{
				Name: "default_service_account",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enabled_features",
				Type: schema.TypeStringArray,
			},
			{
				Name: "id",
				Type: schema.TypeInt,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "quotas",
				Type: schema.TypeJSON,
			},
			{
				Name: "usage_export_location",
				Type: schema.TypeJSON,
			},
			{
				Name: "xpn_project_status",
				Type: schema.TypeString,
			},
			{
				Name: "server_response",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchComputeProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services.Compute.Projects.Get(c.ProjectId).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output
	return nil
}
