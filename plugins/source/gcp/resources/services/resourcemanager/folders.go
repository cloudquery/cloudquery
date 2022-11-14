// Code generated by codegen; DO NOT EDIT.

package resourcemanager

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Folders() *schema.Table {
	return &schema.Table{
		Name:      "gcp_resourcemanager_folders",
		Resolver:  fetchFolders,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "delete_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("DeleteTime"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
		},
	}
}

func fetchFolders(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListFoldersRequest{}
	it := c.Services.ResourcemanagerFoldersClient.ListFolders(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
