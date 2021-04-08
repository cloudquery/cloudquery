package resources

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
		Name:         "aws_efs_filesystems",
		Resolver:     fetchEfsFilesystems,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "creation_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "creation_token",
				Type: schema.TypeString,
			},
			{
				Name: "file_system_id",
				Type: schema.TypeString,
			},
			{
				Name: "life_cycle_state",
				Type: schema.TypeString,
			},
			{
				Name: "number_of_mount_targets",
				Type: schema.TypeInt,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name: "performance_mode",
				Type: schema.TypeString,
			},
			{
				Name:     "size_in_bytes_value",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SizeInBytes.Value"),
			},
			{
				Name:     "size_in_bytes_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SizeInBytes.Timestamp"),
			},
			{
				Name:     "size_in_bytes_value_in_i_a",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SizeInBytes.ValueInIA"),
			},
			{
				Name:     "size_in_bytes_value_in_standard",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SizeInBytes.ValueInStandard"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEfsFilesystemTags,
			},
			{
				Name: "availability_zone_id",
				Type: schema.TypeString,
			},
			{
				Name: "availability_zone_name",
				Type: schema.TypeString,
			},
			{
				Name: "encrypted",
				Type: schema.TypeBool,
			},
			{
				Name: "file_system_arn",
				Type: schema.TypeString,
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "provisioned_throughput_in_mibps",
				Type: schema.TypeFloat,
			},
			{
				Name: "throughput_mode",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEfsFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func resolveEfsFilesystemTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.FileSystemDescription)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
