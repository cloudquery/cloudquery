package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKmsKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Kms

	config := kms.ListKeysInput{Limit: aws.Int32(1000)}
	p := kms.NewListKeysPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Keys
	}
	return nil
}

func getKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Kms
	item := resource.Item.(types.KeyListEntry)

	d, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: item.KeyId})
	if err != nil {
		return err
	}
	resource.Item = d.KeyMetadata
	return nil
}

func resolveKeysReplicaKeys(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key := resource.Item.(*types.KeyMetadata)
	if key.MultiRegionConfiguration == nil {
		return nil
	}
	return resource.Set(c.Name, key.MultiRegionConfiguration.ReplicaKeys)
}

func resolveKeysTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	key := resource.Item.(*types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	params := kms.ListResourceTagsInput{KeyId: key.KeyId}
	tags := make(map[string]string)
	for {
		result, err := svc.ListResourceTags(ctx, &params)
		if err != nil {
			return err
		}
		for _, v := range result.Tags {
			tags[aws.ToString(v.TagKey)] = aws.ToString(v.TagValue)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.Marker = result.NextMarker
	}
	return resource.Set(c.Name, tags)
}

func resolveKeysRotationEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	key := resource.Item.(*types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.KeyRotationEnabled)
}
