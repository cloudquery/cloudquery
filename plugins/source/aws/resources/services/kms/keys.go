package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Type:     schema.TypeBool,
				Resolver: resolveKeysRotationEnabled,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKeysTags,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "replica_keys",
				Type:     schema.TypeJSON,
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
	key := resource.Item.(*types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	svc := meta.(*client.Client).Services().Kms
	params := kms.ListResourceTagsInput{KeyId: key.KeyId}
	paginator := kms.NewListResourceTagsPaginator(svc, &params)
	tags := make(map[string]string)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
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
	svc := meta.(*client.Client).Services().Kms
	result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.KeyRotationEnabled)
}
