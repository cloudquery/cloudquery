package kms

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKmsKeyGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	k := parent.Item.(*types.KeyMetadata)
	config := kms.ListGrantsInput{
		KeyId: k.KeyId,
		Limit: aws.Int32(100),
	}

	c := meta.(*client.Client)
	svc := c.Services().Kms
	p := kms.NewListGrantsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Grants
	}
	return nil
}

func resolveKeyGrantsKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key := resource.Item.(types.GrantListEntry)
	if strings.HasPrefix(aws.ToString(key.KeyId), "arn:") {
		return resource.Set(c.Name, key.KeyId)
	}

	cl := meta.(*client.Client)
	return resource.Set(c.Name, cl.ARN("kms", "key", aws.ToString(key.KeyId)))
}
