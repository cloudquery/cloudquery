package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_security_configurations",
		Description: "Specifies a security configuration",
		Resolver:    fetchGlueSecurityConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:            "name",
				Description:     "The name of the security configuration",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "encryption_configuration",
				Description: "Specifies how data should be encrypted",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("EncryptionConfiguration"),
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
			return err
		}
		res <- result.SecurityConfigurations
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
