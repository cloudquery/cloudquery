package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func KmsKeys() *schema.Table {
	return &schema.Table{
		Name:                 "aws_kms_keys",
		Resolver:             fetchKmsKeys,
		Multiplex:            client.AccountRegionMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveKmsKey,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "rotation_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "cloud_hsm_cluster_id",
				Type: schema.TypeString,
			},
			{
				Name: "creation_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "custom_key_store_id",
				Type: schema.TypeString,
			},
			{
				Name: "customer_master_key_spec",
				Type: schema.TypeString,
			},
			{
				Name: "deletion_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "encryption_algorithms",
				Type: schema.TypeStringArray,
			},
			{
				Name: "expiration_model",
				Type: schema.TypeString,
			},
			{
				Name: "manager",
				Type: schema.TypeString,
			},
			{
				Name: "key_state",
				Type: schema.TypeString,
			},
			{
				Name: "key_usage",
				Type: schema.TypeString,
			},
			{
				Name: "origin",
				Type: schema.TypeString,
			},
			{
				Name: "signing_algorithms",
				Type: schema.TypeStringArray,
			},
			{
				Name: "valid_to",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "key_arn",
				Type: schema.TypeString,
			},
			{
				Name: "key_id",
				Type: schema.TypeString,
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
	r := resource.Item.(types.KeyListEntry)
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

	if string(output.KeyMetadata.Origin) != "EXTERNAL" {
		output, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: r.KeyId}, func(options *kms.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		if err := resource.Set("rotation_enabled", output.KeyRotationEnabled); err != nil {
			return err
		}
	}
	return nil
}
