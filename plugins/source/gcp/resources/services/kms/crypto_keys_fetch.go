package kms

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func fetchCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	call := c.Services.KmsoldService.Projects.Locations.KeyRings.CryptoKeys.List(parent.Data["Name"].(string)).Context(ctx)
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
