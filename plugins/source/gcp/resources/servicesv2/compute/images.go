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
		Options: schema.TableCreationOptions{
			PrimaryKeys: []string{
				"project_id",
				"id",
			},
		},
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "architecture",
				Type: schema.TypeString,
			},
			{
				Name: "archive_size_bytes",
				Type: schema.TypeInt,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "deprecated",
				Type: schema.TypeJSON,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "disk_size_gb",
				Type: schema.TypeInt,
			},
			{
				Name: "family",
				Type: schema.TypeString,
			},
			{
				Name: "guest_os_features",
				Type: schema.TypeJSON,
			},
			{
				Name: "id",
				Type: schema.TypeInt,
			},
			{
				Name: "image_encryption_key",
				Type: schema.TypeJSON,
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
				Name: "license_codes",
				Type: schema.TypeIntArray,
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
				Name: "raw_disk",
				Type: schema.TypeJSON,
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
				Name: "shielded_instance_initial_state",
				Type: schema.TypeJSON,
			},
			{
				Name: "source_disk",
				Type: schema.TypeString,
			},
			{
				Name: "source_disk_encryption_key",
				Type: schema.TypeJSON,
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
				Name: "source_image_encryption_key",
				Type: schema.TypeJSON,
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
				Name: "source_snapshot_encryption_key",
				Type: schema.TypeJSON,
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
