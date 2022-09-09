package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Disks() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_disks",
		Description: "Describes a block storage disk",
		Resolver:    fetchLightsailDisks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the disk",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "attached_to",
				Description: "The resources to which the disk is attached",
				Type:        schema.TypeString,
			},
			{
				Name:        "attachment_state",
				Description: "(Deprecated) The attachment state of the disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The date when the disk was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "gb_in_use",
				Description:   "(Deprecated) The number of GB in use by the disk",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:        "iops",
				Description: "The input/output operations per second (IOPS) of the disk",
				Type:        schema.TypeInt,
			},
			{
				Name:        "is_attached",
				Description: "A Boolean value indicating whether the disk is attached",
				Type:        schema.TypeBool,
			},
			{
				Name:        "is_system_disk",
				Description: "A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it)",
				Type:        schema.TypeBool,
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:        "name",
				Description: "The unique name of the disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "path",
				Description: "The disk path",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type (eg, Disk)",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_in_gb",
				Description: "The size of the disk in GB",
				Type:        schema.TypeInt,
			},
			{
				Name:        "state",
				Description: "Describes the status of the disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name: "add_ons",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_lightsail_disk_snapshot",
				Description:   "Describes a block storage disk snapshot",
				Resolver:      fetchLightsailDiskSnapshots,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "disk_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_disks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the disk snapshot",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The date when the disk snapshot was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "from_disk_arn",
						Description: "The Amazon Resource Name (ARN) of the source disk from which the disk snapshot was created",
						Type:        schema.TypeString,
					},
					{
						Name:        "from_disk_name",
						Description: "The unique name of the source disk from which the disk snapshot was created",
						Type:        schema.TypeString,
					},
					{
						Name:        "from_instance_arn",
						Description: "The Amazon Resource Name (ARN) of the source instance from which the disk (system volume) snapshot was created",
						Type:        schema.TypeString,
					},
					{
						Name:        "from_instance_name",
						Description: "The unique name of the source instance from which the disk (system volume) snapshot was created",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_from_auto_snapshot",
						Description: "A Boolean value indicating whether the snapshot was created from an automatic snapshot",
						Type:        schema.TypeBool,
					},
					{
						Name:     "location",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Location"),
					},
					{
						Name:        "name",
						Description: "The name of the disk snapshot (eg, my-disk-snapshot)",
						Type:        schema.TypeString,
					},
					{
						Name:        "progress",
						Description: "The progress of the snapshot",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The Lightsail resource type (eg, DiskSnapshot)",
						Type:        schema.TypeString,
					},
					{
						Name:        "size_in_gb",
						Description: "The size of the disk in GB",
						Type:        schema.TypeInt,
					},
					{
						Name:        "state",
						Description: "The status of the disk snapshot operation",
						Type:        schema.TypeString,
					},
					{
						Name:        "support_code",
						Description: "The support code",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The tag keys and optional values for the resource",
						Type:        schema.TypeJSON,
						Resolver:    client.ResolveTags,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetDisksInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetDisks(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Disks
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDiskSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetDiskSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetDiskSnapshots(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.DiskSnapshots
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
