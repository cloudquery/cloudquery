package kms

import (
	"context"
	"fmt"

	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func fetchKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	locations, err := getAllKmsLocations(ctx, c)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to get kms locations. %w", err))
	}
	for _, l := range locations {
		it := c.Services.KmsKeyManagementClient.ListKeyRings(ctx, &kmspb.ListKeyRingsRequest{Parent: l.Name})
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
	}
	return nil
}

func getAllKmsLocations(ctx context.Context, c *client.Client) ([]*locationpb.Location, error) {
	var locations []*locationpb.Location
	it := c.Services.KmsKeyManagementClient.ListLocations(ctx, &locationpb.ListLocationsRequest{Name: "projects/" + c.ProjectId})
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		locations = append(locations, resp)
	}
	return locations, nil
}
