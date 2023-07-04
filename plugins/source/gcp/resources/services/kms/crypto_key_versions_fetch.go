package kms

import (
	"context"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchCryptoKeyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*kmspb.CryptoKey)
	kmsClient, err := kms.NewKeyManagementClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := kmsClient.ListCryptoKeyVersions(ctx, &kmspb.ListCryptoKeyVersionsRequest{Parent: p.Name})
	for {
		key, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		res <- key
	}
	return nil
}
