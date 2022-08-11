package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource security_configurations --config security_configurations.hcl --output .
func SecurityConfigurations() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_security_configurations",
		Description:  "Specifies a security configuration",
		Resolver:     fetchGlueSecurityConfigurations,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
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
				Name:        "created_time_stamp",
				Description: "The time at which this security configuration was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "cloud_watch_encryption_mode",
				Description: "The encryption mode to use for CloudWatch data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.CloudWatchEncryption.CloudWatchEncryptionMode"),
			},
			{
				Name:        "cloud_watch_encryption_kms_key_arn",
				Description: "The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.CloudWatchEncryption.KmsKeyArn"),
			},
			{
				Name:        "job_bookmarks_encryption_mode",
				Description: "The encryption mode to use for job bookmarks data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.JobBookmarksEncryption.JobBookmarksEncryptionMode"),
			},
			{
				Name:        "job_bookmarks_encryption_kms_key_arn",
				Description: "The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.JobBookmarksEncryption.KmsKeyArn"),
			},
			{
				Name:        "name",
				Description: "The name of the security configuration",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_security_configuration_s3_encryption",
				Description: "Specifies how Amazon Simple Storage Service (Amazon S3) data should be encrypted",
				Resolver:    schema.PathTableResolver("EncryptionConfiguration.S3Encryption"),
				Columns: []schema.Column{
					{
						Name:        "security_configuration_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_security_configurations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data",
						Type:        schema.TypeString,
					},
					{
						Name:        "s3_encryption_mode",
						Description: "The encryption mode to use for Amazon S3 data",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueSecurityConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetSecurityConfigurationsInput{}
	for {
		result, err := svc.GetSecurityConfigurations(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.SecurityConfigurations
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
