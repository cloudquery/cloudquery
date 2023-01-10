package kms

import (
	"context"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func fetchKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	kmsClient, err := kms.NewKeyManagementClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	locations, err := getAllKmsLocations(ctx, c, kmsClient)
	if err != nil {
		return err
	}
	for _, l := range locations {
		it := kmsClient.ListKeyRings(ctx, &kmspb.ListKeyRingsRequest{Parent: l.Name})
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
	}
	return nil
}

func getAllKmsLocations(ctx context.Context, c *client.Client, kmsClient *kms.KeyManagementClient) ([]*locationpb.Location, error) {
	var locations []*locationpb.Location
	it := kmsClient.ListLocations(ctx, &locationpb.ListLocationsRequest{Name: "projects/" + c.ProjectId})
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		locations = append(locations, resp)
	}
	return locations, nil
}
