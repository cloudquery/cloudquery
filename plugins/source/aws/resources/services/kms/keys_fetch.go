package kms

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKmsKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input kms.ListKeysInput
	c := meta.(*client.Client)
	svc := c.Services().KMS
	for {
		response, err := svc.ListKeys(ctx, &input)
		if err != nil {
			return err
		}
		for _, item := range response.Keys {
			d, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: item.KeyId}, func(options *kms.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			if d.KeyMetadata != nil {
				res <- *d.KeyMetadata
			}
		}
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func resolveKeysReplicaKeys(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key := resource.Item.(types.KeyMetadata)
	if key.MultiRegionConfiguration == nil {
		return nil
	}
	b, err := json.Marshal(key.MultiRegionConfiguration.ReplicaKeys)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveKeysTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().KMS
	key := resource.Item.(types.KeyMetadata)
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
	svc := cl.Services().KMS
	key := resource.Item.(types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.KeyRotationEnabled)
}
