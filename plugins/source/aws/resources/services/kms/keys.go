package kms

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Keys() *schema.Table {
	tableName := "aws_kms_keys"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html`,
		Resolver:            fetchKmsKeys,
		PreResourceResolver: getKey,
		Transform:           transformers.TransformWithStruct(&types.KeyMetadata{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "rotation_enabled",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: resolveKeysRotationEnabled,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveKeysTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
			{
				Name:     "replica_keys",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveKeysReplicaKeys,
			},
		},

		Relations: []*schema.Table{
			keyGrants(),
			keyPolicies(),
		},
	}
}

func fetchKmsKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kms

	config := kms.ListKeysInput{Limit: aws.Int32(1000)}
	p := kms.NewListKeysPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Keys
	}
	return nil
}

func getKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	item := resource.Item.(types.KeyListEntry)

	d, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: item.KeyId}, func(options *kms.Options) {
		options.Region = cl.Region
	})
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
	key := resource.Item.(*types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	params := kms.ListResourceTagsInput{KeyId: key.KeyId}
	paginator := kms.NewListResourceTagsPaginator(svc, &params)
	tags := make(map[string]string)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		// Cannot use client.TagToMap because key/val names are different
		for _, v := range page.Tags {
			tags[aws.ToString(v.TagKey)] = aws.ToString(v.TagValue)
		}
	}
	return resource.Set(c.Name, tags)
}

func resolveKeysRotationEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key := resource.Item.(*types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId}, func(options *kms.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.KeyRotationEnabled)
}
