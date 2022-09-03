package kms

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudkms/v1"
)

type KeyRing struct {
	*cloudkms.KeyRing
	Location string
}

func fetchKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	locations, err := getAllKmsLocations(ctx, c)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to get kms locations. %w", err))
	}
	nextPageToken := ""
	for _, l := range locations {
		call := c.Services.Kms.Projects.Locations.KeyRings.List(l.Name)
		for {
			call.PageToken(nextPageToken)
			resp, err := call.Do()
			if err != nil {
				return errors.WithStack(err)
			}

			rings := make([]*KeyRing, len(resp.KeyRings))
			for i, k := range resp.KeyRings {
				rings[i] = &KeyRing{
					KeyRing:  k,
					Location: l.LocationId,
				}
			}
			res <- rings

			if resp.NextPageToken == "" {
				break
			}
			nextPageToken = resp.NextPageToken
		}
	}
	return nil
}

func getAllKmsLocations(ctx context.Context, c *client.Client) ([]*cloudkms.Location, error) {
	var locations []*cloudkms.Location
	call := c.Services.Kms.Projects.Locations.List("projects/" + c.ProjectId)
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
