package resourcemanager

import (
	"context"

	"google.golang.org/api/iterator"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
			if err := do(resp.Name); err != nil {
				return err
			}
		}
		return nil
	}

	return do(p.Name)
}
