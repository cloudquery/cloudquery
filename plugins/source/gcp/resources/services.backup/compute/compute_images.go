package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeImages() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_images",
		Description: "Represents an Image resource  You can use images to create boot disks for your VM instances",
		Resolver:    fetchComputeImages,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "archive_size_bytes",
				Description: "Size of the image targz archive stored in Google Cloud Storage (in bytes)",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "deprecated_deleted",
				Description: "An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DELETED This is only informational and the status will not change unless the client explicitly changes it",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.Deleted"),
			},
			{
				Name:     "deprecated",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Deprecated"),
			},
			{
				Name:        "deprecated_obsolete",
				Description: "An optional RFC3339 timestamp on or after which the state of this resource is intended to change to OBSOLETE This is only informational and the status will not change unless the client explicitly changes it",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.Obsolete"),
			},
			{
				Name:        "deprecated_replacement",
				Description: "The URL of the suggested replacement for a deprecated resource The suggested replacement resource must be the same kind of resource as the deprecated resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.Replacement"),
			},
			{
				Name:        "deprecated_state",
				Description: "The deprecation state of this resource This can be ACTIVE, DEPRECATED, OBSOLETE, or DELETED Operations which communicate the end of life date for an image, can use ACTIVE Operations which create a new resource using a DEPRECATED resource will return successfully, but with a warning indicating the deprecated resource and recommending its replacement Operations which use OBSOLETE or DELETED resources will be rejected and result in an error",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.State"),
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "disk_size_gb",
				Description: "Size of the image when restored onto a persistent disk (in GB)",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "family",
				Description: "The name of the image family to which this image belongs You can create disks by specifying an image family instead of a specific image name The image family always returns its latest image that is not deprecated The name of the image family must comply with RFC1035",
				Type:        schema.TypeString,
			},
			{
				Name:        "guest_os_features",
				Description: "A list of features to enable on the guest operating system Applicable only for bootable images Read  Enabling guest operating system features to see a list of available options",
				Type:        schema.TypeStringArray,
				Resolver:    resolveComputeImageGuestOsFeatures,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "image_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ImageEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "image_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ImageEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "image_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ImageEncryptionKey.RawKey"),
			},
			{
				Name:        "image_encryption_key_sha256",
				Description: "[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ImageEncryptionKey.Sha256"),
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#image for images",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_fingerprint",
				Description: "A fingerprint for the labels being applied to this image",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Labels for this resource",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "licenses",
				Description:   "Any applicable license URI",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "raw_disk_container_type",
				Description: "The format used to encode and transmit the block device, which should be TAR This is just a container and transmission format and not a runtime format Provided by the client when the disk image is created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RawDisk.ContainerType"),
			},
			{
				Name:        "raw_disk_source",
				Description: "The full Google Cloud Storage URL where the disk image is stored You must provide either this property or the sourceDisk property but not both",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RawDisk.Source"),
			},
			{
				Name:        "satisfies_pzs",
				Description: "Reserved for future use",
				Type:        schema.TypeBool,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "shielded_instance_initial_state_pk_content",
				Description: "The raw content in the secure keys file",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ShieldedInstanceInitialState.Pk.Content"),
			},
			{
				Name:        "shielded_instance_initial_state_pk_file_type",
				Description: "The file type of source file",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ShieldedInstanceInitialState.Pk.FileType"),
			},
			{
				Name:        "source_disk",
				Description: "URL of the source disk used to create this image This can be a full or valid partial URL You must provide either this property or the rawDisksource property but not both to create an image For example, the following are valid values: - https://wwwgoogleapis",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_disk_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceDiskEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "source_disk_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceDiskEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "source_disk_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceDiskEncryptionKey.RawKey"),
			},
			{
				Name:        "source_disk_encryption_key_sha256",
				Description: "[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceDiskEncryptionKey.Sha256"),
			},
			{
				Name:        "source_disk_id",
				Description: "The ID value of the disk used to create this image This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_image",
				Description: "URL of the source image used to create this image  In order to create an image, you must provide the full or partial URL of one of the following: - The selfLink URL - This property - The rawDisk",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_image_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "source_image_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "source_image_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.RawKey"),
			},
			{
				Name:        "source_image_encryption_key_sha256",
				Description: "[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.Sha256"),
			},
			{
				Name:        "source_image_id",
				Description: "The ID value of the image used to create this image This value may be used to determine whether the image was taken from the current or a previous instance of a given image name",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_snapshot",
				Description: "URL of the source snapshot used to create this image  In order to create an image, you must provide the full or partial URL of one of the following: - The selfLink URL - This property - The sourceImage URL - The rawDisk",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_snapshot_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "source_snapshot_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "source_snapshot_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.RawKey"),
			},
			{
				Name:        "source_snapshot_encryption_key_sha256",
				Description: "[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.Sha256"),
			},
			{
				Name:        "source_snapshot_id",
				Description: "The ID value of the snapshot used to create this image This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_type",
				Description: "The type of the image used to create this disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the image An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY Possible values are FAILED, PENDING, or READY",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_locations",
				Description: "Cloud Storage bucket storage location of the image (regional or multi-regional)",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Images.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeImageGuestOsFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Image)
	res := make([]string, len(r.GuestOsFeatures))
	for i, v := range r.GuestOsFeatures {
		res[i] = v.Type
	}
	return errors.WithStack(resource.Set("guest_os_features", res))
}
