package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EfsFilesystems() *schema.Table {
	return &schema.Table{
		Name:          "aws_efs_filesystems",
		Description:   "A description of the file system.",
		Resolver:      fetchEfsFilesystems,
		Multiplex:     client.ServiceAccountRegionMultiplexer("elasticfilesystem"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
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
				Name:     "backup_policy_status",
				Type:     schema.TypeString,
				Resolver: ResolveEfsFilesystemBackupPolicyStatus,
			},
			{
				Name:        "creation_time",
				Description: "The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z). ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "creation_token",
				Description: "The opaque string specified in the request. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the file system, assigned by Amazon EFS. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemId"),
			},
			{
				Name:        "life_cycle_state",
				Description: "The lifecycle phase of the file system. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "number_of_mount_targets",
				Description: "The current number of mount targets that the file system has",
				Type:        schema.TypeInt,
			},
			{
				Name:        "owner_id",
				Description: "The AWS account that created the file system",
				Type:        schema.TypeString,
			},
			{
				Name:        "performance_mode",
				Description: "The performance mode of the file system. ",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_in_bytes_value",
				Description: "The latest known metered size (in bytes) of data stored in the file system. ",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SizeInBytes.Value"),
			},
			{
				Name:        "size_in_bytes_timestamp",
				Description: "The time at which the size of data, returned in the Value field, was determined. The value is the integer number of seconds since 1970-01-01T00:00:00Z.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("SizeInBytes.Timestamp"),
			},
			{
				Name:        "size_in_bytes_value_in_ia",
				Description: "The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SizeInBytes.ValueInIA"),
			},
			{
				Name:        "size_in_bytes_value_in_standard",
				Description: "The latest known metered size (in bytes) of data stored in the Standard storage class.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SizeInBytes.ValueInStandard"),
			},
			{
				Name:        "tags",
				Description: "The tags associated with the file system, presented as an array of Tag objects. ",
				Type:        schema.TypeJSON,
				Resolver:    resolveEfsFilesystemsTags,
			},
			{
				Name:        "availability_zone_id",
				Description: "The unique and consistent identifier of the Availability Zone in which the file system's One Zone storage classes exist",
				Type:        schema.TypeString,
			},
			{
				Name:        "availability_zone_name",
				Description: "Describes the AWS Availability Zone in which the file system is located, and is valid only for file systems using One Zone storage classes",
				Type:        schema.TypeString,
			},
			{
				Name:        "encrypted",
				Description: "A Boolean value that, if true, indicates that the file system is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the EFS file system, in the format arn:aws:elasticfilesystem:region:account-id:file-system/file-system-id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemArn"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the encrypted file system.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "You can add tags to a file system, including a Name tag",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioned_throughput_in_mibps",
				Description: "The amount of provisioned throughput, measured in MiB/s, for the file system. Valid for file systems using ThroughputMode set to provisioned.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "throughput_mode",
				Description: "Displays the file system's throughput mode",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEfsFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config efs.DescribeFileSystemsInput
	c := meta.(*client.Client)
	svc := c.Services().EFS
	for {
		response, err := svc.DescribeFileSystems(ctx, &config, func(options *efs.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.FileSystems
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func ResolveEfsFilesystemBackupPolicyStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	var config efs.DescribeBackupPolicyInput
	client := meta.(*client.Client)
	svc := client.Services().EFS
	response, err := svc.DescribeBackupPolicy(ctx, &config, func(options *efs.Options) {
		options.Region = client.Region
	})
	if err != nil {
		return err
	}

	if response.BackupPolicy == nil {
		return nil
	}

	return resource.Set(c.Name, response.BackupPolicy.Status)
}
func resolveEfsFilesystemsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.FileSystemDescription)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
