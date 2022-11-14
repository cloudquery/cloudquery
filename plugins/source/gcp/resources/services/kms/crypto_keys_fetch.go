package kms

import (
	"context"

	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func fetchCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	p := parent.Item.(*kmspb.KeyRing)

	it := c.Services.KmsKeyManagementClient.ListCryptoKeys(ctx, &kmspb.ListCryptoKeysRequest{Parent: p.Name})
	for {
		key, err := it.Next()
		if key != nil {
			res <- key
		}
		if err != nil {
			if errors.Is(err, iterator.Done) {
				return nil
			}
			return errors.WithStack(err)
		}
	}
}

func resolveRotationPeriod(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, resource.Item.(*kmspb.CryptoKey).GetRotationPeriod().AsDuration())
}
