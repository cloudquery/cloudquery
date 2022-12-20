package resourcemanager

import (
	"context"
	"strings"

	"google.golang.org/api/iterator"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	req := &pb.ListFoldersRequest{
		Parent: "organizations/" + c.OrgId,
	}
	it := fClient.ListFolders(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
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
