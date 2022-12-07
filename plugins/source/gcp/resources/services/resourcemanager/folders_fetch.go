package resourcemanager

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	o, err := c.Services.ResourcemanagerProjectsClient.GetProject(ctx, &pb.GetProjectRequest{Name: "projects/" + c.ProjectId})
	if err != nil {
		return errors.WithStack(err)
	}

	req := &pb.ListFoldersRequest{
		Parent: o.Parent,
	}
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

func resolveOrganizationId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	item := r.Item.(*pb.Folder)
	if !strings.HasPrefix(item.Parent, "organizations/") {
		return nil
	}

	return r.Set(c.Name, strings.TrimPrefix(item.Parent, "organizations/"))
}
