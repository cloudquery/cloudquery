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
		resource.Set("cloud_hsm_cluster_id", output.KeyMetadata.CloudHsmClusterId)
		resource.Set("creation_date", output.KeyMetadata.CreationDate)
		resource.Set("custom_key_store_id", output.KeyMetadata.CustomKeyStoreId)
		resource.Set("customer_master_key_spec", output.KeyMetadata.CustomerMasterKeySpec)
		resource.Set("deletion_date", output.KeyMetadata.DeletionDate)
		resource.Set("description", output.KeyMetadata.Description)
		resource.Set("enabled", output.KeyMetadata.Enabled)
		resource.Set("expiration_model", output.KeyMetadata.ExpirationModel)
		resource.Set("manager", output.KeyMetadata.KeyManager)
		resource.Set("key_state", output.KeyMetadata.KeyState)
		resource.Set("key_usage", output.KeyMetadata.KeyUsage)
		resource.Set("origin", output.KeyMetadata.Origin)
		resource.Set("valid_to", output.KeyMetadata.ValidTo)

		var encryptionAlgorithms []string
		for _, algorithm := range output.KeyMetadata.EncryptionAlgorithms {
			encryptionAlgorithms = append(encryptionAlgorithms, string(algorithm))
		}
		resource.Set("encryption_algorithms", encryptionAlgorithms)

		var signingAlgorithms []string
		for _, algorithm := range output.KeyMetadata.SigningAlgorithms {
			signingAlgorithms = append(signingAlgorithms, string(algorithm))
		}
		resource.Set("signing_algorithms", signingAlgorithms)
	}

	if string(output.KeyMetadata.Origin) != "EXTERNAL" {
		output, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: r.KeyId}, func(options *kms.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		resource.Set("rotation_enabled", output.KeyRotationEnabled)
	}
	return nil
}
