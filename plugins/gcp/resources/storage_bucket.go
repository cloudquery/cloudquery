package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/storage/v1"
)

func StorageBucket() *schema.Table {
	return &schema.Table{
		Name:         "gcp_storage_buckets",
		Resolver:     fetchStorageBuckets,
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
				Name:     "billing_requester_pays",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Billing.RequesterPays"),
			},
			{
				Name: "default_event_based_hold",
				Type: schema.TypeBool,
			},
			{
				Name:     "encryption_default_kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Encryption.DefaultKmsKeyName"),
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name:     "iam_configuration_bucket_policy_only_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IamConfiguration.BucketPolicyOnly.Enabled"),
			},
			{
				Name:     "iam_configuration_bucket_policy_only_locked_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamConfiguration.BucketPolicyOnly.LockedTime"),
			},
			{
				Name:     "iam_configuration_public_access_prevention",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamConfiguration.PublicAccessPrevention"),
			},
			{
				Name:     "iam_configuration_uniform_bucket_level_access_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IamConfiguration.UniformBucketLevelAccess.Enabled"),
			},
			{
				Name:     "iam_configuration_uniform_bucket_level_access_locked_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamConfiguration.UniformBucketLevelAccess.LockedTime"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "location_type",
				Type: schema.TypeString,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name:     "logging_log_bucket",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Logging.LogBucket"),
			},
			{
				Name:     "logging_log_object_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Logging.LogObjectPrefix"),
			},
			{
				Name: "metageneration",
				Type: schema.TypeBigInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "owner_entity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Entity"),
			},
			{
				Name:     "owner_entity_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.EntityId"),
			},
			{
				Name: "project_number",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "retention_policy_effective_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RetentionPolicy.EffectiveTime"),
			},
			{
				Name:     "retention_policy_is_locked",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RetentionPolicy.IsLocked"),
			},
			{
				Name:     "retention_policy_retention_period",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RetentionPolicy.RetentionPeriod"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPZS"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "storage_class",
				Type: schema.TypeString,
			},
			{
				Name: "time_created",
				Type: schema.TypeString,
			},
			{
				Name: "updated",
				Type: schema.TypeString,
			},
			{
				Name:     "versioning_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Versioning.Enabled"),
			},
			{
				Name:     "website_main_page_suffix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Website.MainPageSuffix"),
			},
			{
				Name:     "website_not_found_page",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Website.NotFoundPage"),
			},
			{
				Name: "zone_affinity",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_storage_bucket_cors",
				Resolver: fetchStorageBucketCors,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "max_age_seconds",
						Type: schema.TypeBigInt,
					},
					{
						Name: "method",
						Type: schema.TypeStringArray,
					},
					{
						Name: "origin",
						Type: schema.TypeStringArray,
					},
					{
						Name: "response_header",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "gcp_storage_bucket_acls",
				Resolver: fetchStorageBucketAcls,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "bucket",
						Type: schema.TypeString,
					},
					{
						Name: "domain",
						Type: schema.TypeString,
					},
					{
						Name: "email",
						Type: schema.TypeString,
					},
					{
						Name: "entity",
						Type: schema.TypeString,
					},
					{
						Name: "entity_id",
						Type: schema.TypeString,
					},
					{
						Name: "etag",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name:     "project_team_project_number",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ProjectTeam.ProjectNumber"),
					},
					{
						Name:     "project_team",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ProjectTeam.Team"),
					},
					{
						Name: "role",
						Type: schema.TypeString,
					},
					{
						Name: "self_link",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_storage_bucket_lifecycle_rules",
				Resolver: fetchStorageBucketLifecycleRules,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "action_storage_class",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Action.StorageClass"),
					},
					{
						Name:     "action_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Action.Type"),
					},
					{
						Name:     "condition_age",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Condition.Age"),
					},
					{
						Name:     "condition_created_before",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Condition.CreatedBefore"),
					},
					{
						Name:     "condition_custom_time_before",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Condition.CustomTimeBefore"),
					},
					{
						Name:     "condition_days_since_custom_time",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Condition.DaysSinceCustomTime"),
					},
					{
						Name:     "condition_days_since_noncurrent_time",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Condition.DaysSinceNoncurrentTime"),
					},
					{
						Name:     "condition_is_live",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Condition.IsLive"),
					},
					{
						Name:     "condition_matches_pattern",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Condition.MatchesPattern"),
					},
					{
						Name:     "condition_matches_storage_class",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Condition.MatchesStorageClass"),
					},
					{
						Name:     "condition_noncurrent_time_before",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Condition.NoncurrentTimeBefore"),
					},
					{
						Name:     "condition_num_newer_versions",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Condition.NumNewerVersions"),
					},
				},
			},
			{
				Name:     "gcp_storage_bucket_default_object_acls",
				Resolver: fetchStorageBucketDefaultObjectAcls,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "bucket",
						Type: schema.TypeString,
					},
					{
						Name: "domain",
						Type: schema.TypeString,
					},
					{
						Name: "email",
						Type: schema.TypeString,
					},
					{
						Name: "entity",
						Type: schema.TypeString,
					},
					{
						Name: "entity_id",
						Type: schema.TypeString,
					},
					{
						Name: "etag",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name:     "project_team_project_number",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ProjectTeam.ProjectNumber"),
					},
					{
						Name:     "project_team",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ProjectTeam.Team"),
					},
					{
						Name: "role",
						Type: schema.TypeString,
					},
					{
						Name: "self_link",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageBuckets(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Storage.Buckets.List(c.ProjectId)
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

func fetchStorageBucketAcls(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.Acl
	return nil
}

func fetchStorageBucketCors(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.Cors
	return nil
}

func fetchStorageBucketLifecycleRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	if bucket.Lifecycle != nil {
		res <- bucket.Lifecycle.Rule
	}
	return nil
}
func fetchStorageBucketDefaultObjectAcls(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.DefaultObjectAcl
	return nil
}
