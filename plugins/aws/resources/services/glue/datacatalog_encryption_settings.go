package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource datacatalog_encryption_settings --config datacatalog_encryption_settings.hcl --output .
func DatacatalogEncryptionSettings() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_datacatalog_encryption_settings",
		Description:  "Contains configuration information for maintaining Data Catalog security",
		Resolver:     fetchGlueDatacatalogEncryptionSettings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id"}},
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
				Name:        "return_connection_password_encrypted",
				Description: "When the ReturnConnectionPasswordEncrypted flag is set to \"true\", passwords remain encrypted in the responses of GetConnection and GetConnections",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ConnectionPasswordEncryption.ReturnConnectionPasswordEncrypted"),
			},
			{
				Name:        "aws_kms_key_id",
				Description: "An KMS key that is used to encrypt the connection password",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConnectionPasswordEncryption.AwsKmsKeyId"),
			},
			{
				Name:        "encryption_at_rest_catalog_encryption_mode",
				Description: "The encryption-at-rest mode for encrypting Data Catalog data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionAtRest.CatalogEncryptionMode"),
			},
			{
				Name:        "encryption_at_rest_sse_aws_kms_key_id",
				Description: "The ID of the KMS key to use for encryption at rest",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionAtRest.SseAwsKmsKeyId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueDatacatalogEncryptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	res <- result.DataCatalogEncryptionSettings
	return nil
}
