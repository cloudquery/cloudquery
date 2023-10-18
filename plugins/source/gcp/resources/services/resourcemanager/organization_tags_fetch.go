package resourcemanager

import (
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/iterator"
)

func fetchOrganizationTagKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	fClient, err := resourcemanager.NewTagKeysClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	req := &pb.ListTagKeysRequest{
		Parent: "organizations/" + c.OrgId,
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
		fmt.Printf("key resp: %#v\n", resp)

		res <- resp
	}
	return nil
}

func fetchOgranizationTagValues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
