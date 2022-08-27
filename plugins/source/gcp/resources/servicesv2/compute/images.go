// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func ComputeImages() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_images",
		Resolver:  fetchComputeImages,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "architecture",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Architecture"),
			},
			{
				Name:     "archive_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ArchiveSizeBytes"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "deprecated",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Deprecated"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disk_size_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskSizeGb"),
			},
			{
				Name:     "family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Family"),
			},
			{
				Name:     "guest_os_features",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GuestOsFeatures"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "image_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageEncryptionKey"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "label_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelFingerprint"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "license_codes",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("LicenseCodes"),
			},
			{
				Name:     "licenses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Licenses"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "raw_disk",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RawDisk"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPzs"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
			},
			{
				Name:     "shielded_instance_initial_state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShieldedInstanceInitialState"),
			},
			{
				Name:     "source_disk",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDisk"),
			},
			{
				Name:     "source_disk_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceDiskEncryptionKey"),
			},
			{
				Name:     "source_disk_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDiskId"),
			},
			{
				Name:     "source_image",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImage"),
			},
			{
				Name:     "source_image_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceImageEncryptionKey"),
			},
			{
				Name:     "source_image_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceImageId"),
			},
			{
				Name:     "source_snapshot",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshot"),
			},
			{
				Name:     "source_snapshot_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceSnapshotEncryptionKey"),
			},
			{
				Name:     "source_snapshot_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSnapshotId"),
			},
			{
				Name:     "source_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceType"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "storage_locations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("StorageLocations"),
			},
		},
	}
}

func fetchComputeImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
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
