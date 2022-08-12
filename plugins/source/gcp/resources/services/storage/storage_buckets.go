package storage

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	storage "google.golang.org/api/storage/v1"
)

func StorageBuckets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_buckets",
		Description: "The Buckets resource represents a bucket in Cloud Storage",
		Resolver:    fetchStorageBuckets,
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
				Name:        "billing_requester_pays",
				Description: "When set to true, Requester Pays is enabled for this bucket",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Billing.RequesterPays"),
			},
			{
				Name:        "default_event_based_hold",
				Description: "The default value for event-based hold on newly created objects in this bucket Event-based hold is a way to retain objects indefinitely until an event occurs, signified by the hold's release After being released, such objects will be subject to bucket-level retention",
				Type:        schema.TypeBool,
			},
			{
				Name:        "encryption_default_kms_key_name",
				Description: "A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Encryption.DefaultKmsKeyName"),
			},
			{
				Name:        "encryption_type",
				Description: "A Cloud KMS key type. Possible values: \"CMKE\" - Customer-managed key   \"GMKE\" - Google-managed key",
				Type:        schema.TypeString,
				Resolver:    resolveBucketEncryptionType,
			},
			{
				Name:        "etag",
				Description: "HTTP 11 Entity tag for the bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_configuration_bucket_policy_only_enabled",
				Description: "If set, access is controlled only by bucket-level or above IAM policies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IamConfiguration.BucketPolicyOnly.Enabled"),
			},
			{
				Name:        "iam_configuration_bucket_policy_only_locked_time",
				Description: "The deadline for changing iamConfigurationbucketPolicyOnlyenabled from true to false in RFC 3339 format iamConfigurationbucketPolicyOnlyenabled may be changed from true to false until the locked time, after which the field is immutable",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IamConfiguration.BucketPolicyOnly.LockedTime"),
			},
			{
				Name:        "iam_configuration_public_access_prevention",
				Description: "The bucket's Public Access Prevention configuration Currently, 'unspecified' and 'enforced' are supported",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IamConfiguration.PublicAccessPrevention"),
			},
			{
				Name:        "iam_configuration_uniform_bucket_level_access_enabled",
				Description: "If set, access is controlled only by bucket-level or above IAM policies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IamConfiguration.UniformBucketLevelAccess.Enabled"),
			},
			{
				Name:        "iam_configuration_uniform_bucket_level_access_locked_time",
				Description: "The deadline for changing iamConfigurationuniformBucketLevelAccessenabled from true to false in RFC 3339  format iamConfigurationuniformBucketLevelAccessenabled may be changed from true to false until the locked time, after which the field is immutable",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IamConfiguration.UniformBucketLevelAccess.LockedTime"),
			},
			{
				Name:        "id",
				Description: "Original Id of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "The kind of item this is For buckets, this is always storage#bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "User-provided labels, in key/value pairs",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The location of the bucket Object data for objects in the bucket resides in physical storage within this region Defaults to US See the developer's guide for the authoritative list",
				Type:        schema.TypeString,
			},
			{
				Name:        "location_type",
				Description: "The type of the bucket location",
				Type:        schema.TypeString,
			},
			{
				Name:        "logging_log_bucket",
				Description: "The destination bucket where the current bucket's logs should be placed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Logging.LogBucket"),
			},
			{
				Name:        "logging_log_object_prefix",
				Description: "A prefix for log object names",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Logging.LogObjectPrefix"),
			},
			{
				Name:        "metageneration",
				Description: "The metadata generation of this bucket",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "The name of the bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner_entity",
				Description: "The entity, in the form project-owner-projectId",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.Entity"),
			},
			{
				Name:        "owner_entity_id",
				Description: "The ID for the entity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.EntityId"),
			},
			{
				Name:        "project_number",
				Description: "The project number of the project the bucket belongs to",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "retention_policy_effective_time",
				Description: "Server-determined value that indicates the time from which policy was enforced and effective This value is in RFC 3339 format",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RetentionPolicy.EffectiveTime"),
			},
			{
				Name:        "retention_policy_is_locked",
				Description: "Once locked, an object retention policy cannot be modified",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RetentionPolicy.IsLocked"),
			},
			{
				Name:        "retention_policy_retention_period",
				Description: "The duration in seconds that objects need to be retained Retention duration must be greater than zero and less than 100 years Note that enforcement of retention periods less than a day is not guaranteed Such periods should only be used for testing purposes",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RetentionPolicy.RetentionPeriod"),
			},
			{
				Name:        "satisfies_pzs",
				Description: "Reserved for future use",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SatisfiesPZS"),
			},
			{
				Name:        "self_link",
				Description: "The URI of this bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_class",
				Description: "The bucket's default storage class, used whenever no storageClass is specified for a newly-created object This defines how objects in the bucket are stored and determines the SLA and the cost of storage Values include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE, and DURABLE_REDUCED_AVAILABILITY",
				Type:        schema.TypeString,
			},
			{
				Name:        "time_created",
				Description: "The creation time of the bucket in RFC 3339 format",
				Type:        schema.TypeString,
			},
			{
				Name:        "updated",
				Description: "The modification time of the bucket in RFC 3339 format",
				Type:        schema.TypeString,
			},
			{
				Name:        "versioning_enabled",
				Description: "While set to true, versioning is fully enabled for this bucket",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Versioning.Enabled"),
			},
			{
				Name:        "website_main_page_suffix",
				Description: "If the requested object path is missing, the service will ensure the path has a trailing '/', append this suffix, and attempt to retrieve the resulting object This allows the creation of indexhtml objects to represent directory pages",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Website.MainPageSuffix"),
			},
			{
				Name:        "website_not_found_page",
				Description: "If the requested object path is missing, and any mainPageSuffix object is missing, if applicable, the service will return the named object from this bucket as the content for a 404 Not Found result",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Website.NotFoundPage"),
			},
			{
				Name:          "zone_affinity",
				Description:   "The zone or zones from which the bucket is intended to use zonal quota Requests for data from outside the specified affinities are still allowed but won't be able to use zonal quota The zone or zones need to be within the bucket location otherwise the requests will fail with a 400 Bad Request response",
				IgnoreInTests: true,
				Type:          schema.TypeStringArray,
			},
			{
				Name:        "policy",
				Description: "Bucket's policy",
				Type:        schema.TypeJSON,
				Resolver:    resolveBucketPolicy,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "gcp_storage_bucket_acls",
				Description:   "Access controls on the bucket.",
				Resolver:      fetchStorageBucketAcls,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "bucket_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "bucket",
						Description: "The name of the bucket",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain",
						Description: "The domain associated with the entity, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "email",
						Description: "The email address associated with the entity, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "entity",
						Description: "The entity holding the permission, in one of the following forms: - user-userId - user-email - group-groupId - group-email - domain-domain - project-team-projectId - allUsers - allAuthenticatedUsers Examples: - The user liz@examplecom would be user-liz@examplecom - The group example@googlegroupscom would be group-example@googlegroupscom - To refer to all members of the Google Apps for Business domain examplecom, the entity would be domain-examplecom",
						Type:        schema.TypeString,
					},
					{
						Name:        "entity_id",
						Description: "The ID for the entity, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "HTTP 11 Entity tag for the access-control entry",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The ID of the access-control entry",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "The kind of item this is For bucket access control entries, this is always storage#bucketAccessControl",
						Type:        schema.TypeString,
					},
					{
						Name:        "project_team_project_number",
						Description: "The project number",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ProjectTeam.ProjectNumber"),
					},
					{
						Name:        "project_team",
						Description: "The team",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ProjectTeam.Team"),
					},
					{
						Name:        "role",
						Description: "The access permission for the entity",
						Type:        schema.TypeString,
					},
					{
						Name:        "self_link",
						Description: "The link to this access-control entry",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_storage_bucket_cors",
				Description: "The bucket's Cross-Origin Resource Sharing (CORS) configuration.",
				Resolver:    fetchStorageBucketCors,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "bucket_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "max_age_seconds",
						Description: "The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "method",
						Description: "The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: \"*\" is permitted in the list of methods, and means \"any method\"",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "origin",
						Description: "The list of Origins eligible to receive CORS response headers Note: \"*\" is permitted in the list of origins, and means \"any Origin\"",
						Type:        schema.TypeStringArray,
					},
					{
						Name:          "response_header",
						Description:   "The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "gcp_storage_bucket_default_object_acls",
				Description:   "Default access controls to apply to new objects when no ACL is provided.",
				Resolver:      fetchStorageBucketDefaultObjectAcls,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "bucket_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "bucket",
						Description: "The name of the bucket",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain",
						Description: "The domain associated with the entity, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "email",
						Description: "The email address associated with the entity, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "entity",
						Description: "The entity holding the permission, in one of the following forms: - user-userId - user-email - group-groupId - group-email - domain-domain - project-team-projectId - allUsers - allAuthenticatedUsers Examples: - The user liz@examplecom would be user-liz@examplecom - The group example@googlegroupscom would be group-example@googlegroupscom - To refer to all members of the Google Apps for Business domain examplecom, the entity would be domain-examplecom",
						Type:        schema.TypeString,
					},
					{
						Name:        "entity_id",
						Description: "The ID for the entity, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "HTTP 11 Entity tag for the access-control entry",
						Type:        schema.TypeString,
					},
					{
						Name:        "generation",
						Description: "The content generation of the object, if applied to an object",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "id",
						Description: "The ID of the access-control entry",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "The kind of item this is For object access control entries, this is always storage#objectAccessControl",
						Type:        schema.TypeString,
					},
					{
						Name:        "object",
						Description: "The name of the object, if applied to an object",
						Type:        schema.TypeString,
					},
					{
						Name:        "project_team_project_number",
						Description: "The project number",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ProjectTeam.ProjectNumber"),
					},
					{
						Name:        "project_team",
						Description: "The team",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ProjectTeam.Team"),
					},
					{
						Name:        "role",
						Description: "The access permission for the entity",
						Type:        schema.TypeString,
					},
					{
						Name:        "self_link",
						Description: "The link to this access-control entry",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_storage_bucket_lifecycle_rules",
				Description: "A lifecycle management rule, which is made of an action to take and the condition(s) under which the action will be taken.",
				Resolver:    fetchStorageBucketLifecycleRules,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "bucket_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "action_storage_class",
						Description: "Target storage class Required iff the type of the action is SetStorageClass",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Action.StorageClass"),
					},
					{
						Name:        "action_type",
						Description: "Type of the action Currently, only Delete and SetStorageClass are supported",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Action.Type"),
					},
					{
						Name:        "condition_age",
						Description: "Age of an object (in days) This condition is satisfied when an object reaches the specified age",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Condition.Age"),
					},
					{
						Name:        "condition_created_before",
						Description: "A date in RFC 3339 format with only the date part (for instance, \"2013-01-15\") This condition is satisfied when an object is created before midnight of the specified date in UTC",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Condition.CreatedBefore"),
					},
					{
						Name:        "condition_custom_time_before",
						Description: "A date in RFC 3339 format with only the date part (for instance, \"2013-01-15\") This condition is satisfied when the custom time on an object is before this date in UTC",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Condition.CustomTimeBefore"),
					},
					{
						Name:        "condition_days_since_custom_time",
						Description: "Number of days elapsed since the user-specified timestamp set on an object The condition is satisfied if the days elapsed is at least this number If no custom timestamp is specified on an object, the condition does not apply",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Condition.DaysSinceCustomTime"),
					},
					{
						Name:        "condition_days_since_noncurrent_time",
						Description: "Number of days elapsed since the noncurrent timestamp of an object The condition is satisfied if the days elapsed is at least this number This condition is relevant only for versioned objects The value of the field must be a nonnegative integer If it's zero, the object version will become eligible for Lifecycle action as soon as it becomes noncurrent",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Condition.DaysSinceNoncurrentTime"),
					},
					{
						Name:          "condition_is_live",
						Description:   "Relevant only for versioned objects If the value is true, this condition matches live objects; if the value is false, it matches archived objects",
						Type:          schema.TypeBool,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("Condition.IsLive"),
					},
					{
						Name:        "condition_matches_pattern",
						Description: "A regular expression that satisfies the RE2 syntax This condition is satisfied when the name of the object matches the RE2 pattern Note: This feature is currently in the \"Early Access\" launch stage and is only available to a whitelisted set of users; that means that this feature may be changed in backward-incompatible ways and that it is not guaranteed to be released",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Condition.MatchesPattern"),
					},
					{
						Name:          "condition_matches_storage_class",
						Description:   "Objects having any of the storage classes specified by this condition will be matched Values include MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, and DURABLE_REDUCED_AVAILABILITY",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("Condition.MatchesStorageClass"),
					},
					{
						Name:        "condition_noncurrent_time_before",
						Description: "A date in RFC 3339 format with only the date part (for instance, \"2013-01-15\") This condition is satisfied when the noncurrent time on an object is before this date in UTC This condition is relevant only for versioned objects",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Condition.NoncurrentTimeBefore"),
					},
					{
						Name:        "condition_num_newer_versions",
						Description: "Relevant only for versioned objects If the value is N, this condition is satisfied when there are at least N versions (including the live version) newer than this version of the object",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Condition.NumNewerVersions"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchStorageBucketAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.Acl
	return nil
}
func fetchStorageBucketCors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.Cors
	return nil
}
func fetchStorageBucketDefaultObjectAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.DefaultObjectAcl
	return nil
}
func fetchStorageBucketLifecycleRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	if bucket.Lifecycle != nil {
		res <- bucket.Lifecycle.Rule
	}
	return nil
}
func resolveBucketPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*storage.Bucket)
	cl := meta.(*client.Client)
	output, err := cl.Services.Storage.Buckets.GetIamPolicy(p.Name).OptionsRequestedPolicyVersion(3).Do()
	if err != nil {
		return err
	}

	var policy map[string]interface{}
	data, err := json.Marshal(output)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, policy))
}
func resolveBucketEncryptionType(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*storage.Bucket)
	if p.Encryption == nil {
		return errors.WithStack(resource.Set(c.Name, "GMKE"))
	}
	return errors.WithStack(resource.Set(c.Name, "CMKE"))
}
