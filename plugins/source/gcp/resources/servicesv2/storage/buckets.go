// Code generated by codegen; DO NOT EDIT.

package storage

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "gcp_storage_buckets",
		Resolver:  fetchBuckets,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "acl",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Acl"),
			},
			{
				Name:     "autoclass",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Autoclass"),
			},
			{
				Name:     "billing",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Billing"),
			},
			{
				Name:     "cors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cors"),
			},
			{
				Name:     "custom_placement_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomPlacementConfig"),
			},
			{
				Name:     "default_event_based_hold",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultEventBasedHold"),
			},
			{
				Name:     "default_object_acl",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultObjectAcl"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Encryption"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "iam_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IamConfiguration"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "location_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationType"),
			},
			{
				Name:     "logging",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Logging"),
			},
			{
				Name:     "metageneration",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Metageneration"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Owner"),
			},
			{
				Name:     "project_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ProjectNumber"),
			},
			{
				Name:     "retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RetentionPolicy"),
			},
			{
				Name:     "rpo",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Rpo"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPZS"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
			},
			{
				Name:     "storage_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageClass"),
			},
			{
				Name:     "time_created",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TimeCreated"),
			},
			{
				Name:     "updated",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Updated"),
			},
			{
				Name:     "versioning",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Versioning"),
			},
			{
				Name:     "website",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Website"),
			},
		},
	}
}

func fetchBuckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Storage.Buckets.List(c.ProjectId).PageToken(nextPageToken).Do()
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
