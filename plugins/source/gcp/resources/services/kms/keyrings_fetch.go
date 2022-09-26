package kms

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func fetchKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	locations, err := getAllKmsLocations(c)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to get kms locations. %w", err))
	}
	for _, l := range locations {
		it := c.Services.KmsKeyManagementClient.ListKeyRings(ctx, &pb.ListKeyRingsRequest{
			Parent: l.Name,
		})
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

func getAllKmsLocations(c *client.Client) ([]*cloudkms.Location, error) {
	var locations []*cloudkms.Location
	call := c.Services.KmsoldService.Projects.Locations.List("projects/" + c.ProjectId)
	nextPageToken := ""
	for {
		resp, err := call.PageToken(nextPageToken).Do()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		locations = append(locations, resp.Locations...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return locations, nil
}
