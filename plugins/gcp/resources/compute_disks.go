package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeDisks() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_disks",
		Resolver:     fetchComputeDisks,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "disk_encryption_key_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionKey.KmsKeyName"),
			},
			{
				Name:     "disk_encryption_key_kms_key_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:     "disk_encryption_key_raw_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionKey.RawKey"),
			},
			{
				Name:     "disk_encryption_key_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskEncryptionKey.Sha256"),
			},
			{
				Name:     "guest_os_features",
				Type:     schema.TypeStringArray,
				Resolver: resolveComputeDiskGuestOsFeatures,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "label_fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "last_attach_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "last_detach_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "licenses",
				Type: schema.TypeStringArray,
			},
			{
				Name: "location_hint",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "options",
				Type: schema.TypeString,
			},
			{
				Name: "physical_block_size_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "provisioned_iops",
				Type: schema.TypeBigInt,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "replica_zones",
				Type: schema.TypeStringArray,
			},
			{
				Name: "resource_policies",
				Type: schema.TypeStringArray,
			},
			{
				Name: "satisfies_pzs",
				Type: schema.TypeBool,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "size_gb",
				Type: schema.TypeBigInt,
			},
			{
				Name: "source_disk",
				Type: schema.TypeString,
			},
			{
				Name: "source_disk_id",
				Type: schema.TypeString,
			},
			{
				Name: "source_image",
				Type: schema.TypeString,
			},
			{
				Name:     "source_image_encryption_key_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImageEncryptionKey.KmsKeyName"),
			},
			{
				Name:     "source_image_encryption_key_kms_key_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImageEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:     "source_image_encryption_key_raw_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImageEncryptionKey.RawKey"),
			},
			{
				Name:     "source_image_encryption_key_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImageEncryptionKey.Sha256"),
			},
			{
				Name: "source_image_id",
				Type: schema.TypeString,
			},
			{
				Name: "source_snapshot",
				Type: schema.TypeString,
			},
			{
				Name:     "source_snapshot_encryption_key_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshotEncryptionKey.KmsKeyName"),
			},
			{
				Name:     "source_snapshot_encryption_key_kms_key_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshotEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:     "source_snapshot_encryption_key_raw_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshotEncryptionKey.RawKey"),
			},
			{
				Name:     "source_snapshot_encryption_key_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshotEncryptionKey.Sha256"),
			},
			{
				Name: "source_snapshot_id",
				Type: schema.TypeString,
			},
			{
				Name: "source_storage_object",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "users",
				Type: schema.TypeStringArray,
			},
			{
				Name: "zone",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Disks.AggregatedList(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var diskTypes []*compute.Disk
		for _, items := range output.Items {
			diskTypes = append(diskTypes, items.Disks...)
		}
		res <- diskTypes

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeDiskGuestOsFeatures(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*compute.Disk)
	res := make([]string, len(r.GuestOsFeatures))
	for i, v := range r.GuestOsFeatures {
		res[i] = v.Type
	}
	return resource.Set("guest_os_features", res)
}
