package kms

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func fetchCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	it := c.Services.KmsKeyManagementClient.ListCryptoKeys(ctx, &pb.ListCryptoKeysRequest{
		Parent: parent.Data["name"].(string),
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
	return nil
}
