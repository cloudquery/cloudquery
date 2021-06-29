package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	storage "google.golang.org/api/storage/v1"
)

func StorageBuckets() *schema.Table {
	return &schema.Table{
		Name:         "gcp_storage_buckets",
		Description:  "The Buckets resource represents a bucket in Cloud Storage",
		Resolver:     fetchStorageBuckets,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
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
				Name:        "resource_id",
				Description: "Original Id of the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
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
				Name:        "zone_affinity",
				Description: "The zone or zones from which the bucket is intended to use zonal quota Requests for data from outside the specified affinities are still allowed but won't be able to use zonal quota The zone or zones need to be within the bucket location otherwise the requests will fail with a 400 Bad Request response",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_storage_bucket_acls",
				Description: "Access controls on the bucket.",
				Resolver:    fetchStorageBucketAcls,
				Columns: []schema.Column{
					{
						Name:        "bucket_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
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
						Name:        "resource_id",
						Description: "The ID of the access-control entry",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
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
						Name:        "bucket_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
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
						Name:        "response_header",
						Description: "The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "gcp_storage_bucket_default_object_acls",
				Description: "Default access controls to apply to new objects when no ACL is provided.",
				Resolver:    fetchStorageBucketDefaultObjectAcls,
				Columns: []schema.Column{
					{
						Name:        "bucket_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
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
						Name:        "resource_id",
						Description: "The ID of the access-control entry",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
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
						Name:        "bucket_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
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
						Name:        "condition_is_live",
						Description: "Relevant only for versioned objects If the value is true, this condition matches live objects; if the value is false, it matches archived objects",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Condition.IsLive"),
					},
					{
						Name:        "condition_matches_pattern",
						Description: "A regular expression that satisfies the RE2 syntax This condition is satisfied when the name of the object matches the RE2 pattern Note: This feature is currently in the \"Early Access\" launch stage and is only available to a whitelisted set of users; that means that this feature may be changed in backward-incompatible ways and that it is not guaranteed to be released",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Condition.MatchesPattern"),
					},
					{
						Name:        "condition_matches_storage_class",
						Description: "Objects having any of the storage classes specified by this condition will be matched Values include MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, and DURABLE_REDUCED_AVAILABILITY",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Condition.MatchesStorageClass"),
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
			{
				Name:        "gcp_storage_bucket_policies",
				Description: "A bucket/object IAM policy",
				Resolver:    fetchStorageBucketPolicies,
				Columns: []schema.Column{
					{
						Name:        "bucket_id",
						Description: "Unique ID of gcp_storage_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "etag",
						Description: "HTTP 11  Entity tag for the policy",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "The kind of item this is For policies, this is always storage#policy This field is ignored on input",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_id",
						Description: "The ID of the resource to which this policy belongs Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects A specific generation may be specified by appending #generationNumber to the end of the object name, eg projects/_/buckets/my-bucket/objects/datatxt#17 The current generation can be denoted with #0 This field is ignored on input",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "The IAM policy format version",
						Type:        schema.TypeBigInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_storage_bucket_policy_bindings",
						Description: "Represents an expression text Example: title: \"User account presence\" description: \"Determines whether the request has a user account\" expression: \"size(request",
						Resolver:    fetchStorageBucketPolicyBindings,
						Columns: []schema.Column{
							{
								Name:        "bucket_policy_id",
								Description: "Unique ID of gcp_storage_bucket_policies table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "condition_description",
								Description: "An optional description of the expression This is a longer text which describes the expression, eg when hovered over it in a UI",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Condition.Description"),
							},
							{
								Name:        "condition_expression",
								Description: "Textual representation of an expression in Common Expression Language syntax The application context of the containing message determines which well-known feature set of CEL is supported",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Condition.Expression"),
							},
							{
								Name:        "condition_location",
								Description: "An optional string indicating the location of the expression for error reporting, eg a file name and a position in the file",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Condition.Location"),
							},
							{
								Name:        "condition_title",
								Description: "An optional title for the expression, ie a short string describing its purpose This can be used eg in UIs which allow to enter the expression",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Condition.Title"),
							},
							{
								Name:        "members",
								Description: "A collection of identifiers for members who may assume the provided role Recognized identifiers are as follows: - allUsers — A special identifier that represents anyone on the internet; with or without a Google account - allAuthenticatedUsers — A special identifier that represents anyone who is authenticated with a Google account or a service account - user:emailid — An email address that represents a specific account For example, user:alice@gmailcom or user:joe@examplecom  - serviceAccount:emailid — An email address that represents a service account For example, serviceAccount:my-other-app@appspotgserviceaccountcom  - group:emailid — An email address that represents a Google group For example, group:admins@examplecom - domain:domain — A Google Apps domain name that represents all the users of that domain For example, domain:googlecom or domain:examplecom - projectOwner:projectid — Owners of the given project For example, projectOwner:my-example-project - projectEditor:projectid — Editors of the given project For example, projectEditor:my-example-project - projectViewer:projectid — Viewers of the given project",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "role",
								Description: "The role to which members belong Two types of roles are supported: new IAM roles, which grant permissions that do not map directly to those provided by ACLs, and legacy IAM roles, which do map directly to ACL permissions All roles are of the format roles/storagespecificRole The new IAM roles are: - roles/storageadmin — Full control of Google Cloud Storage resources - roles/storageobjectViewer — Read-Only access to Google Cloud Storage objects - roles/storageobjectCreator — Access to create objects in Google Cloud Storage - roles/storageobjectAdmin — Full control of Google Cloud Storage objects   The legacy IAM roles are: - roles/storagelegacyObjectReader — Read-only access to objects without listing Equivalent to an ACL entry on an object with the READER role - roles/storagelegacyObjectOwner — Read/write access to existing objects without listing Equivalent to an ACL entry on an object with the OWNER role - roles/storagelegacyBucketReader — Read access to buckets with object listing Equivalent to an ACL entry on a bucket with the READER role - roles/storagelegacyBucketWriter — Read access to buckets with object listing/creation/deletion Equivalent to an ACL entry on a bucket with the WRITER role - roles/storagelegacyBucketOwner — Read and write access to existing buckets with object listing/creation/deletion Equivalent to an ACL entry on a bucket with the OWNER role",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Storage.Buckets.List(c.ProjectId).Context(ctx).PageToken(nextPageToken)
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
func fetchStorageBucketAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.Acl
	return nil
}
func fetchStorageBucketCors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.Cors
	return nil
}
func fetchStorageBucketDefaultObjectAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	res <- bucket.DefaultObjectAcl
	return nil
}
func fetchStorageBucketLifecycleRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*storage.Bucket)
	if bucket.Lifecycle != nil {
		res <- bucket.Lifecycle.Rule
	}
	return nil
}
func fetchStorageBucketPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*storage.Bucket)
	if !ok {
		return fmt.Errorf("expected *storage.Bucket but got %T", p)
	}
	c := meta.(*client.Client)
	call := c.Services.Storage.Buckets.GetIamPolicy(p.Name).Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output
	return nil
}
func fetchStorageBucketPolicyBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*storage.Policy)
	if !ok {
		return fmt.Errorf("expected *storage.Policy but got %T", p)
	}
	res <- p.Bindings
	return nil
}
