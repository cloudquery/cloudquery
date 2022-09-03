// Code generated by codegen; DO NOT EDIT.

package cloudresourcemanager

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "gcp_cloudresourcemanager_projects",
		Resolver:  fetchProjects,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: resolveProjectPolicy,
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "delete_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeleteTime"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "parent",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Parent"),
			},
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProjectId"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdateTime"),
			},
		},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services.Cloudresourcemanager.Projects.Get("projects/" + c.ProjectId).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output

	return nil
}
