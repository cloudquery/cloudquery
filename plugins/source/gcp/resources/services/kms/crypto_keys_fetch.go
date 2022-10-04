package kms

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func fetchCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	p := parent.Item.(*kms.KeyRing)

	nextPageToken := ""
	call := c.Services.KmsoldService.Projects.Locations.KeyRings.CryptoKeys.List(p.Name).Context(ctx)
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- resp.CryptoKeys

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}

	return nil
}
