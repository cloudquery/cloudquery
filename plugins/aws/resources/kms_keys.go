package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func KmsKeys() *schema.Table {
	return &schema.Table{
		Name:                 "aws_kms_keys",
		Description:          "Contains information about each entry in the key list.",
		Resolver:             fetchKmsKeys,
		Multiplex:            client.ServiceAccountRegionMultiplexer("kms"),
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveKmsKey,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "rotation_enabled",
				Description: "specifies whether key rotation is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "cloud_hsm_cluster_id",
				Description: "The cluster ID of the AWS CloudHSM cluster that contains the key material for the CMK",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_date",
				Description: "The date and time when the CMK was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "custom_key_store_id",
				Description: "A unique identifier for the custom key store.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_master_key_spec",
				Description: "Describes the type of key material in the CMK.",
				Type:        schema.TypeString,
			},
			{
				Name:        "deletion_date",
				Description: "he date and time after which AWS KMS deletes the CMK. This value is present only when KeyState is PendingDeletion.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the CMK.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "Specifies whether the CMK is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "encryption_algorithms",
				Description: "The encryption algorithms that the CMK supports.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "expiration_model",
				Description: "Specifies whether the CMK's key material expires.",
				Type:        schema.TypeString,
			},
			{
				Name:        "manager",
				Description: "The manager of the CMK.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_state",
				Description: "The current status of the CMK.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_usage",
				Description: "The cryptographic operations for which you can use the CMK.",
				Type:        schema.TypeString,
			},
			{
				Name:        "origin",
				Description: "The source of the CMK's key material.",
				Type:        schema.TypeString,
			},
			{
				Name:        "signing_algorithms",
				Description: "The signing algorithms that the CMK supports.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "valid_to",
				Description: "The time at which the imported key material expires.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "ARN of the key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KeyArn"),
			},
			{
				Name:        "id",
				Description: "Unique identifier of the key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KeyId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchKmsKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input kms.ListKeysInput
	c := meta.(*client.Client)
	svc := c.Services().KMS
	for {
		response, err := svc.ListKeys(ctx, &input, func(options *kms.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Keys
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func resolveKmsKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r, ok := resource.Item.(types.KeyListEntry)
	if !ok {
		return fmt.Errorf("expected types.KeyListEntry but got %T", resource.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().KMS
	output, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: r.KeyId}, func(options *kms.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	if output.KeyMetadata != nil {
		if err := resource.Set("cloud_hsm_cluster_id", output.KeyMetadata.CloudHsmClusterId); err != nil {
			return err
		}
		if err := resource.Set("creation_date", output.KeyMetadata.CreationDate); err != nil {
			return err
		}
		if err := resource.Set("custom_key_store_id", output.KeyMetadata.CustomKeyStoreId); err != nil {
			return err
		}
		if err := resource.Set("customer_master_key_spec", output.KeyMetadata.CustomerMasterKeySpec); err != nil {
			return err
		}
		if err := resource.Set("deletion_date", output.KeyMetadata.DeletionDate); err != nil {
			return err
		}
		if err := resource.Set("description", output.KeyMetadata.Description); err != nil {
			return err
		}
		if err := resource.Set("enabled", output.KeyMetadata.Enabled); err != nil {
			return err
		}
		if err := resource.Set("expiration_model", output.KeyMetadata.ExpirationModel); err != nil {
			return err
		}
		if err := resource.Set("manager", output.KeyMetadata.KeyManager); err != nil {
			return err
		}
		if err := resource.Set("key_state", output.KeyMetadata.KeyState); err != nil {
			return err
		}
		if err := resource.Set("key_usage", output.KeyMetadata.KeyUsage); err != nil {
			return err
		}
		if err := resource.Set("origin", output.KeyMetadata.Origin); err != nil {
			return err
		}
		if err := resource.Set("valid_to", output.KeyMetadata.ValidTo); err != nil {
			return err
		}
		var encryptionAlgorithms []string
		for _, algorithm := range output.KeyMetadata.EncryptionAlgorithms {
			encryptionAlgorithms = append(encryptionAlgorithms, string(algorithm))
		}
		if err := resource.Set("encryption_algorithms", encryptionAlgorithms); err != nil {
			return err
		}

		var signingAlgorithms []string
		for _, algorithm := range output.KeyMetadata.SigningAlgorithms {
			signingAlgorithms = append(signingAlgorithms, string(algorithm))
		}
		if err := resource.Set("signing_algorithms", signingAlgorithms); err != nil {
			return err
		}
	}

	if string(output.KeyMetadata.Origin) != "EXTERNAL" && string(output.KeyMetadata.KeyManager) != "AWS" {
		output, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: r.KeyId}, func(options *kms.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		if err := resource.Set("rotation_enabled", output.KeyRotationEnabled); err != nil {
			return err
		}

		tagsResponse, err := svc.ListResourceTags(ctx, &kms.ListResourceTagsInput{
			KeyId: r.KeyId,
		}, func(options *kms.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		tags := make(map[string]interface{})
		for _, t := range tagsResponse.Tags {
			tags[*t.TagKey] = *t.TagValue
		}
		return resource.Set("tags", tags)
	}
	return nil
}
