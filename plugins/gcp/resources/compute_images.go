package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeImages() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_images",
		Resolver:     fetchComputeImages,
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
				Name: "archive_size_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name:     "deprecated_deleted",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Deleted"),
			},
			{
				Name:     "deprecated",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Deprecated"),
			},
			{
				Name:     "deprecated_obsolete",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Obsolete"),
			},
			{
				Name:     "deprecated_replacement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Replacement"),
			},
			{
				Name:     "deprecated_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.State"),
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "disk_size_gb",
				Type: schema.TypeBigInt,
			},
			{
				Name: "family",
				Type: schema.TypeString,
			},
			{
				Name:     "guest_os_features",
				Type:     schema.TypeStringArray,
				Resolver: resolveComputeImageGuestOsFeatures,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name:     "image_encryption_key_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageEncryptionKey.KmsKeyName"),
			},
			{
				Name:     "image_encryption_key_kms_key_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:     "image_encryption_key_raw_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageEncryptionKey.RawKey"),
			},
			{
				Name:     "image_encryption_key_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageEncryptionKey.Sha256"),
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
				Name: "licenses",
				Type: schema.TypeStringArray,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "raw_disk_container_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RawDisk.ContainerType"),
			},
			{
				Name:     "raw_disk_sha1_checksum",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RawDisk.Sha1Checksum"),
			},
			{
				Name:     "raw_disk_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RawDisk.Source"),
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
				Name:     "shielded_instance_initial_state_pk_content",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ShieldedInstanceInitialState.Pk.Content"),
			},
			{
				Name:     "shielded_instance_initial_state_pk_file_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ShieldedInstanceInitialState.Pk.FileType"),
			},
			{
				Name: "source_disk",
				Type: schema.TypeString,
			},
			{
				Name:     "source_disk_encryption_key_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDiskEncryptionKey.KmsKeyName"),
			},
			{
				Name:     "source_disk_encryption_key_kms_key_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDiskEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:     "source_disk_encryption_key_raw_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDiskEncryptionKey.RawKey"),
			},
			{
				Name:     "source_disk_encryption_key_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDiskEncryptionKey.Sha256"),
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
				Name: "source_type",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "storage_locations",
				Type: schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeImages(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Images.List(c.ProjectId)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeImageGuestOsFeatures(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*compute.Image)
	res := make([]string, len(r.GuestOsFeatures))
	for i, v := range r.GuestOsFeatures {
		res[i] = v.Type
	}
	resource.Set("guest_os_features", res)
	return nil
}
