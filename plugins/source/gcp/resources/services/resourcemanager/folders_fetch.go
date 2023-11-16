package resourcemanager

import (
	"context"
	"errors"

	"google.golang.org/api/iterator"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchFolders(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return err
		}

		res <- resp
	}
	return nil
}

func fetchSubFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*pb.Folder)

	fClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	var do func(string) error
	do = func(parentName string) error {
		req := &pb.ListFoldersRequest{
			Parent: parentName,
		}
		it := fClient.ListFolders(ctx, req, c.CallOptions...)
		for {
			resp, err := it.Next()
			if err != nil {
				if errors.Is(err, iterator.Done) {
					break
				}
				return err
			}

			res <- resp
			if err := do(resp.Name); err != nil {
				return err
			}
		}
		return nil
	}

	return do(p.Name)
}
