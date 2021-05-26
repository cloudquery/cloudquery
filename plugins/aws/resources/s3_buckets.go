package resources

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func S3Buckets() *schema.Table {
	return &schema.Table{
		Name:                 "aws_s3_buckets",
		Resolver:             fetchS3Buckets,
		Multiplex:            client.AccountMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountFilter,
		PostResourceResolver: resolveS3BucketsAttributes,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "logging_target_prefix",
				Type: schema.TypeString,
			},
			{
				Name: "logging_target_bucket",
				Type: schema.TypeString,
			},
			{
				Name: "versioning_status",
				Type: schema.TypeString,
			},
			{
				Name: "versioning_mfa_delete",
				Type: schema.TypeString,
			},
			{
				Name: "policy",
				Type: schema.TypeJSON,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "creation_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "block_public_acls",
				Type: schema.TypeBool,
			},
			{
				Name: "block_public_policy",
				Type: schema.TypeBool,
			},
			{
				Name: "ignore_public_acls",
				Type: schema.TypeBool,
			},
			{
				Name: "restrict_public_buckets",
				Type: schema.TypeBool,
			},
			{
				Name: "replication_role",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_s3_bucket_grants",
				Resolver: fetchS3BucketGrants,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Grantee.Type"),
					},
					{
						Name:     "display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Grantee.DisplayName"),
					},
					{
						Name:     "email_address",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Grantee.EmailAddress"),
					},
					{
						Name:     "grantee_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Grantee.ID"),
					},
					{
						Name:     "uri",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Grantee.URI"),
					},
					{
						Name: "permission",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_s3_bucket_cors_rules",
				Resolver: fetchS3BucketCorsRules,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "allowed_methods",
						Type: schema.TypeStringArray,
					},
					{
						Name: "allowed_origins",
						Type: schema.TypeStringArray,
					},
					{
						Name: "allowed_headers",
						Type: schema.TypeStringArray,
					},
					{
						Name: "expose_headers",
						Type: schema.TypeStringArray,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "max_age_seconds",
						Type: schema.TypeInt,
					},
				},
			},
			{
				Name:     "aws_s3_bucket_encryption_rules",
				Resolver: fetchS3BucketEncryptionRules,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "sse_algorithm",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ApplyServerSideEncryptionByDefault.SSEAlgorithm"),
					},
					{
						Name:     "kms_master_key_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ApplyServerSideEncryptionByDefault.KMSMasterKeyID"),
					},
					{
						Name: "bucket_key_enabled",
						Type: schema.TypeBool,
					},
				},
			},
			{
				Name:     "aws_s3_bucket_replication_rules",
				Resolver: fetchS3BucketReplicationRules,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "destination_bucket",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.Bucket"),
					},
					{
						Name:     "destination_access_control_translation_owner",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.AccessControlTranslation.Owner"),
					},
					{
						Name:     "destination_account",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.Account"),
					},
					{
						Name:     "destination_encryption_configuration_replica_kms_key_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.EncryptionConfiguration.ReplicaKmsKeyID"),
					},
					{
						Name:     "destination_metrics_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.Metrics.Status"),
					},
					{
						Name:     "destination_metrics_event_threshold_minutes",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Destination.Metrics.EventThreshold.Minutes"),
					},
					{
						Name:     "destination_replication_time_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.ReplicationTime.Status"),
					},
					{
						Name:     "destination_replication_time_minutes",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Destination.ReplicationTime.Time.Minutes"),
					},
					{
						Name:     "destination_storage_class",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Destination.StorageClass"),
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name:     "delete_marker_replication_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeleteMarkerReplication.Status"),
					},
					{
						Name:     "existing_object_replication_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ExistingObjectReplication.Status"),
					},
					{
						Name:     "filter",
						Type:     schema.TypeJSON,
						Resolver: resolveS3BucketReplicationRuleFilter,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "prefix",
						Type: schema.TypeString,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "source_replica_modifications_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SourceSelectionCriteria.ReplicaModifications.Status"),
					},
					{
						Name:     "source_sse_kms_encrypted_objects_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SourceSelectionCriteria.SseKmsEncryptedObjects.Status"),
					},
				},
			},
			{
				Name:     "aws_s3_bucket_lifecycles",
				Resolver: fetchS3BucketLifecycles,
				Columns: []schema.Column{
					{
						Name:     "bucket_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name:     "abort_incomplete_multipart_upload_days_after_initiation",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("AbortIncompleteMultipartUpload.DaysAfterInitiation"),
					},
					{
						Name:     "expiration_date",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("Expiration.Date"),
					},
					{
						Name:     "expiration_days",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Expiration.Days"),
					},
					{
						Name:     "expiration_expired_object_delete_marker",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Expiration.ExpiredObjectDeleteMarker"),
					},
					{
						Name:     "filter",
						Type:     schema.TypeJSON,
						Resolver: resolveS3BucketLifecycleFilter,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "noncurrent_version_expiration_days",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("NoncurrentVersionExpiration.NoncurrentDays"),
					},
					{
						Name:     "noncurrent_version_transitions",
						Type:     schema.TypeJSON,
						Resolver: resolveS3BucketLifecycleNoncurrentVersionTransitions,
					},
					{
						Name: "prefix",
						Type: schema.TypeString,
					},
					{
						Name:     "transitions",
						Type:     schema.TypeJSON,
						Resolver: resolveS3BucketLifecycleTransitions,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

type WrappedBucket struct {
	types.Bucket
	ReplicationRole  *string
	ReplicationRules []types.ReplicationRule
}

func fetchS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().S3
	response, err := svc.ListBuckets(ctx, nil)
	if err != nil {
		return err
	}
	wb := make([]*WrappedBucket, len(response.Buckets))
	for i, b := range response.Buckets {
		wb[i] = &WrappedBucket{b, nil, nil}
	}

	res <- wb
	return nil
}
func resolveS3BucketsAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	var ae smithy.APIError
	log := meta.Logger()
	r := resource.Item.(*WrappedBucket)
	log.Info("bucket name", r.Name)
	mgr := meta.(*client.Client).Services().S3Manager
	output, err := mgr.GetBucketRegion(ctx, *r.Name)
	if err != nil {
		if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchBucket" {
			// https://aws.amazon.com/premiumsupport/knowledge-center/s3-listing-deleted-bucket/
			// deleted buckets may show up
			log.Debug("Skipping bucket (already deleted)", "bucket", *r.Name)
			return nil
		}
		return err
	}
	svc := meta.(*client.Client).Services().S3
	bucketRegion := "us-east-1"
	if output != "" {
		// This is a weird corner case by AWS API https://github.com/aws/aws-sdk-net/issues/323#issuecomment-196584538
		bucketRegion = output
	}
	if err := resource.Set("region", bucketRegion); err != nil {
		return err
	}

	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		return err
	}
	if loggingOutput.LoggingEnabled != nil {
		if err := resource.Set("logging_target_bucket", loggingOutput.LoggingEnabled.TargetBucket); err != nil {
			return err
		}
		if err := resource.Set("logging_target_prefix", loggingOutput.LoggingEnabled.TargetPrefix); err != nil {
			return err
		}
	}

	policyOutput, err := svc.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "NoSuchBucketPolicy") {
		return err
	}
	if policyOutput != nil {
		if err := resource.Set("policy", policyOutput.Policy); err != nil {
			return err
		}
	}

	versioningOutput, err := svc.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		return err
	}

	if err := resource.Set("versioning_status", versioningOutput.Status); err != nil {
		return err
	}
	if err := resource.Set("versioning_mfa_delete", versioningOutput.MFADelete); err != nil {
		return err
	}

	publicAccessOutput, err := svc.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		return err
	}
	if err := resource.Set("block_public_acls", publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicAcls); err != nil {
		return err
	}
	if err := resource.Set("block_public_policy", publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicPolicy); err != nil {
		return err
	}
	if err := resource.Set("ignore_public_acls", publicAccessOutput.PublicAccessBlockConfiguration.IgnorePublicAcls); err != nil {
		return err
	}
	if err := resource.Set("restrict_public_buckets", publicAccessOutput.PublicAccessBlockConfiguration.RestrictPublicBuckets); err != nil {
		return err
	}

	replicationOutput, err := svc.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if errors.As(err, &ae) && ae.ErrorCode() == "ReplicationConfigurationNotFoundError" {
		return nil
	}
	if replicationOutput.ReplicationConfiguration != nil {
		if err := resource.Set("replication_role", replicationOutput.ReplicationConfiguration.Role); err != nil {
			return err
		}
		// We set this here for fetchReplicationRules to get and insert
		resource.Item.(*WrappedBucket).ReplicationRules = replicationOutput.ReplicationConfiguration.Rules
	}

	taggingOutput, err := svc.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		return err
	}
	tags := make(map[string]*string, len(taggingOutput.TagSet))
	for _, t := range taggingOutput.TagSet {
		tags[*t.Key] = t.Value
	}
	if err := resource.Set("tags", tags); err != nil {
		return err
	}

	return nil
}

func fetchS3BucketGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).Services().S3
	aclOutput, err := svc.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		return err
	}
	res <- aclOutput.Grants
	return nil
}
func fetchS3BucketCorsRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var ae smithy.APIError
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).Services().S3
	CORSOutput, err := svc.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "NoSuchCORSConfiguration") {
		return err
	}
	if CORSOutput != nil {
		res <- CORSOutput.CORSRules
	}
	return nil
}
func fetchS3BucketEncryptionRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var ae smithy.APIError
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).Services().S3
	aclOutput, err := svc.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		if errors.As(err, &ae) && ae.ErrorCode() == "ServerSideEncryptionConfigurationNotFoundError" {
			return nil
		}
		return err
	}
	res <- aclOutput.ServerSideEncryptionConfiguration.Rules
	return nil
}

func fetchS3BucketReplicationRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	bucket := parent.Item.(*WrappedBucket)
	if bucket.ReplicationRules != nil {
		res <- bucket.ReplicationRules
	}
	return nil
}

func resolveS3BucketReplicationRuleFilter(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	rule := resource.Item.(types.ReplicationRule)
	if rule.Filter == nil {
		return nil
	}
	data, err := json.Marshal(rule.Filter)
	if err != nil {
		return err
	}
	return resource.Set("filter", data)
}

func fetchS3BucketLifecycles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var ae smithy.APIError
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).Services().S3
	lifecycleOutput, err := svc.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchLifecycleConfiguration" {
			return nil
		}
		return err
	}
	res <- lifecycleOutput.Rules
	return nil
}
func resolveS3BucketLifecycleNoncurrentVersionTransitions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	lc := resource.Item.(types.LifecycleRule)
	if lc.Transitions == nil {
		return nil
	}
	data, err := json.Marshal(lc.Transitions)
	if err != nil {
		return err
	}
	return resource.Set("noncurrent_version_transitions", data)
}
func resolveS3BucketLifecycleTransitions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	lc := resource.Item.(types.LifecycleRule)
	if lc.Transitions == nil {
		return nil
	}
	data, err := json.Marshal(lc.Transitions)
	if err != nil {
		return err
	}
	return resource.Set("transitions", data)
}

func resolveS3BucketLifecycleFilter(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	lc := resource.Item.(types.LifecycleRule)
	if lc.Filter == nil {
		return nil
	}
	data, err := json.Marshal(lc.Filter)
	if err != nil {
		return err
	}
	return resource.Set("filter", data)
}
