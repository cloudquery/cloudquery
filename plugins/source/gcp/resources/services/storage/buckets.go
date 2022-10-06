// Code generated by codegen; DO NOT EDIT.

package storage

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "acl",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ACL"),
			},
			{
				Name:     "bucket_policy_only",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BucketPolicyOnly"),
			},
			{
				Name:     "uniform_bucket_level_access",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UniformBucketLevelAccess"),
			},
			{
				Name:     "public_access_prevention",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PublicAccessPrevention"),
			},
			{
				Name:     "default_object_acl",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultObjectACL"),
			},
			{
				Name:     "default_event_based_hold",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultEventBasedHold"),
			},
			{
				Name:     "predefined_acl",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PredefinedACL"),
			},
			{
				Name:     "predefined_default_object_acl",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PredefinedDefaultObjectACL"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "custom_placement_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomPlacementConfig"),
			},
			{
				Name:     "meta_generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MetaGeneration"),
			},
			{
				Name:     "storage_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageClass"),
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "versioning_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VersioningEnabled"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "requester_pays",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequesterPays"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RetentionPolicy"),
			},
			{
				Name:     "cors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CORS"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Encryption"),
			},
			{
				Name:     "logging",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Logging"),
			},
			{
				Name:     "website",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Website"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "location_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationType"),
			},
			{
				Name:     "project_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ProjectNumber"),
			},
			{
				Name:     "rpo",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RPO"),
			},
		},

		Relations: []*schema.Table{
			BucketPolicies(),
		},
	}
}
