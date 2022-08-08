package kms

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=keys.hcl -domain=kms -resource=keys
func Keys() *schema.Table {
	return &schema.Table{
		Name:         "aws_kms_keys",
		Description:  "Contains metadata about a KMS key",
		Resolver:     fetchKmsKeys,
		Multiplex:    client.ServiceAccountRegionMultiplexer("kms"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "rotation_enabled",
				Description: "Specifies whether key rotation is enabled.",
				Type:        schema.TypeBool,
				Resolver:    resolveKeysRotationEnabled,
			},
			{
				Name:        "tags",
				Description: "Key tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveKeysTags,
			},
			{
				Name:        "id",
				Description: "The globally unique identifier for the KMS key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KeyId"),
			},
			{
				Name:        "aws_account_id",
				Description: "The twelve-digit account ID of the Amazon Web Services account that owns the KMS key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AWSAccountId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the KMS key",
				Type:        schema.TypeString,
			},
			{
				Name:          "cloud_hsm_cluster_id",
				Description:   "The cluster ID of the CloudHSM cluster that contains the key material for the KMS key",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "creation_date",
				Description: "The date and time when the KMS key was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "custom_key_store_id",
				Description:   "A unique identifier for the custom key store that contains the KMS key",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "deletion_date",
				Description:   "The date and time after which KMS deletes this KMS key",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "description",
				Description: "The description of the KMS key.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "Specifies whether the KMS key is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "encryption_algorithms",
				Description: "The encryption algorithms that the KMS key supports",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "expiration_model",
				Description: "Specifies whether the KMS key's key material expires",
				Type:        schema.TypeString,
			},
			{
				Name:        "manager",
				Description: "The manager of the KMS key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KeyManager"),
			},
			{
				Name:        "key_spec",
				Description: "Describes the type of key material in the KMS key.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_state",
				Description: "The current status of the KMS key",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_usage",
				Description: "The cryptographic operations for which you can use the KMS key.",
				Type:        schema.TypeString,
			},
			{
				Name:          "mac_algorithms",
				Description:   "The message authentication code (MAC) algorithm that the HMAC KMS key supports. This value is present only when the KeyUsage of the KMS key is GENERATE_VERIFY_MAC.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "multi_region",
				Description: "Indicates whether the KMS key is a multi-Region (True) or regional (False) key. This value is True for multi-Region primary and replica keys and False for regional KMS keys",
				Type:        schema.TypeBool,
			},
			{
				Name:        "multi_region_key_type",
				Description: "Indicates whether the KMS key is a PRIMARY or REPLICA key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MultiRegionConfiguration.MultiRegionKeyType"),
			},
			{
				Name:          "primary_key_arn",
				Description:   "Displays the key ARN of a primary or replica key of a multi-Region key.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("MultiRegionConfiguration.PrimaryKey.Arn"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_key_region",
				Description:   "Displays the Amazon Web Services Region of a primary or replica key in a multi-Region key.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("MultiRegionConfiguration.PrimaryKey.Region"),
				IgnoreInTests: true,
			},
			{
				Name:          "replica_keys",
				Description:   "displays the key ARNs and Regions of all replica keys",
				Type:          schema.TypeJSON,
				Resolver:      resolveKeysReplicaKeys,
				IgnoreInTests: true,
			},
			{
				Name:        "origin",
				Description: "The source of the key material for the KMS key",
				Type:        schema.TypeString,
			},
			{
				Name:          "pending_deletion_window_in_days",
				Description:   "The waiting period before the primary key in a multi-Region key is deleted",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "signing_algorithms",
				Description:   "The signing algorithms that the KMS key supports",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "valid_to",
				Description:   "The time at which the imported key material expires",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchKmsKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input kms.ListKeysInput
	c := meta.(*client.Client)
	svc := c.Services().KMS
	for {
		response, err := svc.ListKeys(ctx, &input, func(options *kms.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range response.Keys {
			d, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: item.KeyId}, func(options *kms.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
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
		result, err := svc.ListResourceTags(ctx, &params, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, v := range result.Tags {
			tags[aws.ToString(v.TagKey)] = aws.ToString(v.TagValue)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.Marker = result.NextMarker
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
func resolveKeysRotationEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().KMS
	key := resource.Item.(types.KeyMetadata)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId}, func(options *kms.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.KeyRotationEnabled))
}
