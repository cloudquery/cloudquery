package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func FsxBackups() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_backups",
		Description:  "A backup of an Amazon FSx file system.",
		Resolver:     fetchFsxBackups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("fsx"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "id",
				Description: "The ID of the backup.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BackupId"),
			},
			{
				Name:        "creation_time",
				Description: "The time when a particular backup was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the backup.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the file system backup.",
				Type:        schema.TypeString,
			},
			{
				Name:        "directory_information_active_directory_id",
				Description: "The ID of the AWS Managed Microsoft Active Directory instance to which the file system is joined.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryInformation.ActiveDirectoryId"),
			},
			{
				Name:        "directory_information_domain_name",
				Description: "The fully qualified domain name of the self-managed AD directory.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryInformation.DomainName"),
			},
			{
				Name:        "failure_details_message",
				Description: "A message describing the backup creation failure.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDetails.Message"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the backup of the Amazon FSx file system's data at rest.",
				Type:        schema.TypeString,
			},
			{
				Name:        "progress_percent",
				Description: "The current percent of progress of an asynchronous task.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the backup resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceARN"),
			},
			{
				Name:        "tags",
				Description: "Tags associated with a particular file system.",
				Type:        schema.TypeJSON,
				Resolver:    resolveFsxBackupTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchFsxBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config fsx.DescribeBackupsInput
	c := meta.(*client.Client)
	svc := c.Services().FSX
	for {
		response, err := svc.DescribeBackups(ctx, &config, func(options *fsx.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Backups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveFsxBackupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Backup)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}
