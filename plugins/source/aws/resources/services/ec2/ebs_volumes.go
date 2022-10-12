// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EbsVolumes() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_ebs_volumes",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Volume.html",
		Resolver:    fetchEc2EbsVolumes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveEbsVolumeArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attachments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attachments"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Encrypted"),
			},
			{
				Name:     "fast_restored",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FastRestored"),
			},
			{
				Name:     "iops",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Iops"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "multi_attach_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAttachEnabled"),
			},
			{
				Name:     "outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutpostArn"),
			},
			{
				Name:     "size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "snapshot_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotId"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "throughput",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Throughput"),
			},
			{
				Name:     "volume_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VolumeId"),
			},
			{
				Name:     "volume_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VolumeType"),
			},
		},
	}
}
