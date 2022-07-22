package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource instance_snapshots --config gen.hcl --output .
func InstanceSnapshots() *schema.Table {
	return &schema.Table{
		Name:          "aws_lightsail_instance_snapshots",
		Description:   "Describes an instance snapshot",
		Resolver:      fetchLightsailInstanceSnapshots,
		Multiplex:     client.ServiceAccountRegionMultiplexer("lightsail"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		IgnoreInTests: true, // can't be deployed using terraform.
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the snapshot (eg, arn:aws:lightsail:us-east-2:123456789101:InstanceSnapshot/d23b5706-3322-4d83-81e5-12345EXAMPLE)",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the snapshot was created (eg, 1479907467024)",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "from_blueprint_id",
				Description: "The blueprint ID from which you created the snapshot (eg, os_debian_8_3)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_bundle_id",
				Description: "The bundle ID from which you created the snapshot (eg, micro_1_0)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_instance_arn",
				Description: "The Amazon Resource Name (ARN) of the instance from which the snapshot was created (eg, arn:aws:lightsail:us-east-2:123456789101:Instance/64b8404c-ccb1-430b-8daf-12345EXAMPLE)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_instance_name",
				Description: "The instance from which the snapshot was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_from_auto_snapshot",
				Description: "A Boolean value indicating whether the snapshot was created from an automatic snapshot",
				Type:        schema.TypeBool,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "name",
				Description: "The name of the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "progress",
				Description: "The progress of the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The type of resource (usually InstanceSnapshot)",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_in_gb",
				Description: "The size in GB of the SSD",
				Type:        schema.TypeInt,
			},
			{
				Name:        "state",
				Description: "The state the snapshot is in",
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
				Resolver:    resolveInstanceSnapshotsTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lightsail_instance_snapshot_from_attached_disks",
				Description: "Describes a block storage disk",
				Resolver:    fetchLightsailInstanceSnapshotFromAttachedDisks,
				Columns: []schema.Column{
					{
						Name:        "instance_snapshot_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instance_snapshots table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the disk",
						Type:        schema.TypeString,
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
						Name:        "gb_in_use",
						Description: "(Deprecated) The number of GB in use by the disk",
						Type:        schema.TypeInt,
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
						Name:        "location_availability_zone",
						Description: "The Availability Zone",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.AvailabilityZone"),
					},
					{
						Name:        "location_region_name",
						Description: "The AWS Region name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.RegionName"),
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
						Resolver:    resolveInstanceSnapshotFromAttachedDisksTags,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_lightsail_instance_snapshot_from_attached_disk_add_ons",
						Description: "Describes an add-on that is enabled for an Amazon Lightsail resource",
						Resolver:    fetchLightsailInstanceSnapshotFromAttachedDiskAddOns,
						Columns: []schema.Column{
							{
								Name:        "instance_snapshot_from_attached_disk_cq_id",
								Description: "Unique CloudQuery ID of aws_lightsail_instance_snapshot_from_attached_disks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The name of the add-on",
								Type:        schema.TypeString,
							},
							{
								Name:        "next_snapshot_time_of_day",
								Description: "The next daily time an automatic snapshot will be created",
								Type:        schema.TypeString,
							},
							{
								Name:        "snapshot_time_of_day",
								Description: "The daily time when an automatic snapshot is created",
								Type:        schema.TypeString,
							},
							{
								Name:        "status",
								Description: "The status of the add-on",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailInstanceSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetInstanceSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetInstanceSnapshots(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.InstanceSnapshots
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func resolveInstanceSnapshotsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.InstanceSnapshot)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchLightsailInstanceSnapshotFromAttachedDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.InstanceSnapshot)
	res <- r.FromAttachedDisks
	return nil
}
func resolveInstanceSnapshotFromAttachedDisksTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Disk)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchLightsailInstanceSnapshotFromAttachedDiskAddOns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Disk)
	res <- r.AddOns
	return nil
}
