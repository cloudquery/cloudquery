// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/compute/apiv1"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_images",
		Resolver:  fetchImages,
		Multiplex: client.ProjectMultiplex("compute.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListImagesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewImagesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.List(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
