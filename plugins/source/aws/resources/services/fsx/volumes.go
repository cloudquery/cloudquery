package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource volumes --config volumes.hcl --output .
func Volumes() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_volumes",
		Description:  "Describes an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volume",
		Resolver:     fetchFsxVolumes,
		Multiplex:    client.ServiceAccountRegionMultiplexer("fsx"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "creation_time",
				Description: "The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "file_system_id",
				Description: "The globally unique ID of the file system, assigned by Amazon FSx",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the volume",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle_transition_reason_message",
				Description: "A detailed error message",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LifecycleTransitionReason.Message"),
			},
			{
				Name:        "name",
				Description: "The name of the volume",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for a given resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceARN"),
			},
			{
				Name:        "tags",
				Description: "A list of Tag values, with a maximum of 50 elements",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "id",
				Description: "The system-generated, unique ID of the volume",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VolumeId"),
			},
			{
				Name:        "volume_type",
				Description: "The type of the volume",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_fsx_volume_ontap_configuration",
				Description: "The configuration of an Amazon FSx for NetApp ONTAP volume",
				Resolver:    schema.PathTableResolver("OntapConfiguration"),
				Columns: []schema.Column{
					{
						Name:        "volume_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_volumes table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "flex_cache_endpoint_type",
						Description: "Specifies the FlexCache endpoint type of the volume",
						Type:        schema.TypeString,
					},
					{
						Name:        "junction_path",
						Description: "Specifies the directory that network-attached storage (NAS) clients use to mount the volume, along with the storage virtual machine (SVM) Domain Name System (DNS) name or IP address",
						Type:        schema.TypeString,
					},
					{
						Name:        "volume_type",
						Description: "Specifies the type of volume",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OntapVolumeType"),
					},
					{
						Name:        "security_style",
						Description: "The security style for the volume, which can be UNIX, NTFS, or MIXED",
						Type:        schema.TypeString,
					},
					{
						Name:        "size_in_megabytes",
						Description: "The configured size of the volume, in megabytes (MBs)",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "storage_efficiency_enabled",
						Description: "The volume's storage efficiency setting",
						Type:        schema.TypeBool,
					},
					{
						Name:        "storage_virtual_machine_id",
						Description: "The ID of the volume's storage virtual machine",
						Type:        schema.TypeString,
					},
					{
						Name:        "storage_virtual_machine_root",
						Description: "A Boolean flag indicating whether this volume is the root volume for its storage virtual machine (SVM)",
						Type:        schema.TypeBool,
					},
					{
						Name:        "tiering_policy_cooling_period",
						Description: "Specifies the number of days that user data in a volume must remain inactive before it is considered \"cold\" and moved to the capacity pool",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("TieringPolicy.CoolingPeriod"),
					},
					{
						Name:        "tiering_policy_name",
						Description: "Specifies the tiering policy used to transition data",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TieringPolicy.Name"),
					},
					{
						Name:        "uuid",
						Description: "The volume's universally unique identifier (UUID)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UUID"),
					},
				},
			},
			{
				Name:        "aws_fsx_volume_open_zfs_configuration",
				Description: "The configuration of an Amazon FSx for OpenZFS volume",
				Resolver:    schema.PathTableResolver("OpenZFSConfiguration"),
				Columns: []schema.Column{
					{
						Name:        "volume_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_volumes table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "copy_tags_to_snapshots",
						Description: "A Boolean value indicating whether tags for the volume should be copied to snapshots",
						Type:        schema.TypeBool,
					},
					{
						Name:        "data_compression_type",
						Description: "Specifies the method used to compress the data on the volume",
						Type:        schema.TypeString,
					},
					{
						Name:        "nfs_exports",
						Description: "The configuration object for mounting a Network File System (NFS) file system",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "origin_snapshot_copy_strategy",
						Description: "The strategy used when copying data from the snapshot to the new volume",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OriginSnapshot.CopyStrategy"),
					},
					{
						Name:        "origin_snapshot_arn",
						Description: "The Amazon Resource Name (ARN) for a given resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OriginSnapshot.SnapshotARN"),
					},
					{
						Name:        "parent_volume_id",
						Description: "The ID of the parent volume",
						Type:        schema.TypeString,
					},
					{
						Name:        "read_only",
						Description: "A Boolean value indicating whether the volume is read-only",
						Type:        schema.TypeBool,
					},
					{
						Name:        "record_size",
						Description: "The record size of an OpenZFS volume, in kibibytes (KiB)",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("RecordSizeKiB"),
					},
					{
						Name:        "storage_capacity_quota",
						Description: "The maximum amount of storage in gibibtyes (GiB) that the volume can use from its parent",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StorageCapacityQuotaGiB"),
					},
					{
						Name:        "storage_capacity_reservation",
						Description: "The amount of storage in gibibytes (GiB) to reserve from the parent volume",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("StorageCapacityReservationGiB"),
					},
					{
						Name:        "user_and_group_quotas",
						Description: "An object specifying how much storage users or groups can use on the volume",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "volume_path",
						Description: "The path to the volume from the root volume",
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

func fetchFsxVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeVolumesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeVolumesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.Volumes
	}
	return nil
}
