package resourcemanager

import (
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/iterator"
)

func fetchProjectTagKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewTagKeysClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	req := &pb.ListTagKeysRequest{
		Parent: "projects/" + c.ProjectId,
	}
	it := fClient.ListTagKeys(ctx, req)
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

func fetchProjectTagValues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewTagValuesClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	var parentName string
	if parent != nil {
		parentName = parent.Item.(*pb.TagKey).Name
	}
	req := &pb.ListTagValuesRequest{
		Parent: parentName,
	}

	it := fClient.ListTagValues(ctx, req)
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

func fetchProjectTagBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewTagBindingsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	req := &pb.ListTagBindingsRequest{
		Parent: "//cloudresourcemanager.googleapis.com/projects/" + c.ProjectId,
	}

	it := fClient.ListTagBindings(ctx, req)
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
