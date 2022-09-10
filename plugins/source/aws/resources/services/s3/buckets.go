package s3

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WrappedBucket struct {
	types.Bucket
	ReplicationRole       *string
	ReplicationRules      []types.ReplicationRule
	Region                string
	LoggingTargetBucket   *string
	LoggingTargetPrefix   *string
	Policy                *string
	VersioningStatus      types.BucketVersioningStatus
	VersioningMfaDelete   types.MFADeleteStatus
	BlockPublicAcls       bool
	BlockPublicPolicy     bool
	IgnorePublicAcls      bool
	RestrictPublicBuckets bool
	Tags                  *string
	OwnershipControls     []string
}

// fetchS3BucketsPoolSize describes the amount of go routines that resolve the S3 buckets
const fetchS3BucketsPoolSize = 10

func Buckets() *schema.Table {
	return &schema.Table{
		Name:         "aws_s3_buckets",
		Description:  "An Amazon S3 bucket is a public cloud storage resource available in Amazon Web Services' (AWS) Simple Storage Service (S3)",
		Resolver:     fetchS3Buckets,
		Multiplex:    client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
			},
			{
				Name:          "logging_target_prefix",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "logging_target_bucket",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
				Name:          "policy",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "tags",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "creation_date",
				Description: "Date the bucket was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "block_public_acls",
				Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket",
				Type:        schema.TypeBool,
			},
			{
				Name:        "block_public_policy",
				Description: "Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ignore_public_acls",
				Description: "Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket",
				Type:        schema.TypeBool,
			},
			{
				Name:        "restrict_public_buckets",
				Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket",
				Type:        schema.TypeBool,
			},
			{
				Name:          "replication_role",
				Description:   "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNGlobal(client.S3Service, func(resource *schema.Resource) ([]string, error) {
					return []string{*resource.Item.(*WrappedBucket).Name}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "ownership_controls",
				Description: "The OwnershipControls (BucketOwnerEnforced, BucketOwnerPreferred, or ObjectWriter) currently in effect for this Amazon S3 bucket.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "replication_rules",
				Description:   "Specifies which Amazon S3 objects to replicate and where to store the replicas.",
				Type:          schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_s3_bucket_grants",
				Description: "Container for grant information.",
				
				Resolver:    fetchS3BucketGrants,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of grantee",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.Type"),
					},
					{
						Name:        "display_name",
						Description: "Screen name of the grantee.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.DisplayName"),
					},
					{
						Name:          "email_address",
						Description:   "Email address of the grantee",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Grantee.EmailAddress"),
						IgnoreInTests: true,
					},
					{
						Name:        "grantee_id",
						Description: "The canonical user ID of the grantee.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.ID"),
					},
					{
						Name:          "uri",
						Description:   "URI of the grantee group.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Grantee.URI"),
						IgnoreInTests: true,
					},
					{
						Name:        "permission",
						Description: "Specifies the permission given to the grantee.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_s3_bucket_cors_rules",
				Description:   "Specifies a cross-origin access rule for an Amazon S3 bucket.",
				Resolver:      fetchS3BucketCorsRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allowed_methods",
						Description: "An HTTP method that you allow the origin to execute",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_origins",
						Description: "One or more origins you want customers to be able to access the bucket from.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_headers",
						Description: "Headers that are specified in the Access-Control-Request-Headers header",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "expose_headers",
						Description: "One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "id",
						Description: "Unique identifier for the rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "max_age_seconds",
						Description: "The time in seconds that your browser is to cache the preflight response for the specified resource.",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:          "aws_s3_bucket_encryption_rules",
				Description:   "Specifies the default server-side encryption configuration.",
				Resolver:      fetchS3BucketEncryptionRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "sse_algorithm",
						Description: "Server-side encryption algorithm to use for the default encryption.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ApplyServerSideEncryptionByDefault.SSEAlgorithm"),
					},
					{
						Name:        "kms_master_key_id",
						Description: "AWS Key Management Service (KMS) customer master key ID to use for the default encryption",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ApplyServerSideEncryptionByDefault.KMSMasterKeyID"),
					},
					{
						Name:        "bucket_key_enabled",
						Description: "Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:          "aws_s3_bucket_lifecycles",
				Description:   "A lifecycle rule for individual objects in an Amazon S3 bucket.",
				Resolver:      fetchS3BucketLifecycles,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_s3_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "status",
						Description: "If 'Enabled', the rule is currently being applied",
						Type:        schema.TypeString,
					},
					{
						Name:        "abort_incomplete_multipart_upload_days_after_initiation",
						Description: "Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("AbortIncompleteMultipartUpload.DaysAfterInitiation"),
					},
					{
						Name:        "expiration_date",
						Description: "Indicates at what date the object is to be moved or deleted",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Expiration.Date"),
					},
					{
						Name:        "expiration_days",
						Description: "Indicates the lifetime, in days, of the objects that are subject to the rule. The value must be a non-zero positive integer.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Expiration.Days"),
					},
					{
						Name:        "expiration_expired_object_delete_marker",
						Description: "Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Expiration.ExpiredObjectDeleteMarker"),
					},
					{
						Name:        "filter",
						Description: "The Filter is used to identify objects that a Lifecycle Rule applies to",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Filter"),
					},
					{
						Name:        "id",
						Description: "Unique identifier for the rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "noncurrent_version_expiration_days",
						Description: "Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("NoncurrentVersionExpiration.NoncurrentDays"),
					},
					{
						Name:        "noncurrent_version_transitions",
						Description: "Specifies the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Transitions"),
					},
					{
						Name:        "prefix",
						Description: "Prefix identifying one or more objects to which the rule applies",
						Type:        schema.TypeString,
					},
					{
						Name:        "transitions",
						Description: "Specifies when an Amazon S3 object transitions to a specified storage class.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Transitions"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

// listBucketRegion identifies the canonical region for S3 based on the partition
// in the future we might want to make this configurable if users are alright with the fact that performing this
// action in different regions will return different results
func listBucketRegion(cl *client.Client) string {
	switch cl.Partition {
	case "aws-cn":
		return "cn-north-1"
	case "aws-us-gov":
		return "us-gov-west-1"
	default:
		return "us-east-1"
	}
}

func fetchS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().S3
	response, err := svc.ListBuckets(ctx, nil, func(options *s3.Options) {
		options.Region = listBucketRegion(cl)
	})
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	buckets := make(chan types.Bucket)
	errs := make(chan error)
	for i := 0; i < fetchS3BucketsPoolSize; i++ {
		wg.Add(1)
		go fetchS3BucketsWorker(ctx, meta, buckets, errs, res, &wg)
	}
	go func() {
		defer close(buckets)
		for _, bucket := range response.Buckets {
			select {
			case <-ctx.Done():
				return
			case buckets <- bucket:
			}
		}
	}()
	done := make(chan struct{})
	go func() {
		for err = range errs {
			cl.Logger().Err(err).Msg("failed to fetch s3 bucket")
		}
		close(done)
	}()
	wg.Wait()
	close(errs)
	<-done

	return nil
}

func fetchS3BucketsWorker(ctx context.Context, meta schema.ClientMeta, buckets <-chan types.Bucket, errs chan<- error, res chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	cl := meta.(*client.Client)
	for bucket := range buckets {
		wb := &WrappedBucket{Bucket: bucket}
		err := resolveS3BucketsAttributes(ctx, meta, wb)
		if err != nil {
			if !isBucketNotFoundError(cl, err) {
				errs <- err
			}
			continue
		}
		res <- wb
	}
}

func resolveS3BucketsAttributes(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket) error {
	c := meta.(*client.Client)
	mgr := c.Services().S3Manager

	output, err := mgr.GetBucketRegion(ctx, *resource.Name)
	if err != nil {
		if isBucketNotFoundError(c, err) {
			return nil
		}
		return err
	}
	// AWS does not specify a region if bucket is in us-east-1, so as long as no error we can assume an empty string is us-east-1
	resource.Region = "us-east-1"
	if output != "" {
		resource.Region = output
	}
	if err = resolveBucketLogging(ctx, meta, resource, resource.Region); err != nil {
		if isBucketNotFoundError(c, err) {
			return nil
		}
		return err
	}

	if err = resolveBucketPolicy(ctx, meta, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketVersioning(ctx, meta, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketPublicAccessBlock(ctx, meta, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketReplication(ctx, meta, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketTagging(ctx, meta, resource, resource.Region); err != nil {
		return err
	}

	return resolveBucketOwnershipControls(ctx, meta, resource, resource.Region)
}

func fetchS3BucketGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).Services().S3
	if parent.Get("region").(string) == "" {
		return nil
	}
	aclOutput, err := svc.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchBucket") {
			return nil
		}
		return err
	}
	res <- aclOutput.Grants
	return nil
}
func fetchS3BucketCorsRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	if parent.Get("region").(string) == "" {
		return nil
	}
	corsOutput, err := svc.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchCORSConfiguration", "NoSuchBucket") {
			return nil
		}
		return err
	}
	if corsOutput != nil {
		res <- corsOutput.CORSRules
	}
	return nil
}
func fetchS3BucketEncryptionRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	if parent.Get("region").(string) == "" {
		return nil
	}
	aclOutput, err := svc.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		if client.IsAWSError(err, "ServerSideEncryptionConfigurationNotFoundError") {
			return nil
		}
		return err
	}
	res <- aclOutput.ServerSideEncryptionConfiguration.Rules
	return nil
}

func fetchS3BucketLifecycles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	if parent.Get("region").(string) == "" {
		return nil
	}
	lifecycleOutput, err := svc.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = parent.Get("region").(string)
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchLifecycleConfiguration") {
			return nil
		}
		return err
	}
	res <- lifecycleOutput.Rules
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func resolveBucketLogging(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	svc := meta.(*client.Client).Services().S3
	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetBucketLogging")
			return nil
		}
		return err
	}
	if loggingOutput.LoggingEnabled == nil {
		return nil
	}
	resource.LoggingTargetBucket = loggingOutput.LoggingEnabled.TargetBucket
	resource.LoggingTargetPrefix = loggingOutput.LoggingEnabled.TargetPrefix
	return nil
}

func resolveBucketPolicy(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	policyOutput, err := svc.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	// check if we got an error but its access denied we can continue
	if err != nil {
		// if we got an error, and it's not a NoSuchBucketError, return err
		if client.IsAWSError(err, "NoSuchBucketPolicy") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetBucketPolicy")
			return nil
		}
		return err
	}
	if policyOutput == nil {
		return nil
	}
	resource.Policy = policyOutput.Policy
	return nil
}

func resolveBucketVersioning(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	versioningOutput, err := svc.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetBucketVersioning")
			return nil
		}
		return err
	}
	resource.VersioningStatus = versioningOutput.Status
	resource.VersioningMfaDelete = versioningOutput.MFADelete
	return nil
}

func resolveBucketPublicAccessBlock(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	publicAccessOutput, err := svc.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		// If we received any error other than NoSuchPublicAccessBlockConfiguration, we return and error
		if isBucketNotFoundError(c, err) {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetPublicAccessBlock")
			return nil
		}
		return err
	}
	resource.BlockPublicAcls = publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicAcls
	resource.BlockPublicPolicy = publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicPolicy
	resource.IgnorePublicAcls = publicAccessOutput.PublicAccessBlockConfiguration.IgnorePublicAcls
	resource.RestrictPublicBuckets = publicAccessOutput.PublicAccessBlockConfiguration.RestrictPublicBuckets
	return nil
}

func resolveBucketReplication(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	replicationOutput, err := svc.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})

	if err != nil {
		// If we received any error other than ReplicationConfigurationNotFoundError, we return and error
		if client.IsAWSError(err, "ReplicationConfigurationNotFoundError") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetBucketReplication")
			return nil
		}
		return err
	}
	if replicationOutput.ReplicationConfiguration == nil {
		return nil
	}
	resource.ReplicationRole = replicationOutput.ReplicationConfiguration.Role
	resource.ReplicationRules = replicationOutput.ReplicationConfiguration.Rules
	return nil
}

func resolveBucketTagging(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	taggingOutput, err := svc.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		// If buckets tags are not set it will return an error instead of empty result
		if client.IsAWSError(err, "NoSuchTagSet") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetBucketTagging")
			return nil
		}
		return err
	}
	if taggingOutput == nil {
		return nil
	}
	tags := make(map[string]*string, len(taggingOutput.TagSet))
	for _, t := range taggingOutput.TagSet {
		tags[*t.Key] = t.Value
	}

	b, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	t := string(b)
	resource.Tags = &t
	return nil
}

func resolveBucketOwnershipControls(ctx context.Context, meta schema.ClientMeta, resource *WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3

	getBucketOwnershipControlOutput, err := svc.GetBucketOwnershipControls(ctx, &s3.GetBucketOwnershipControlsInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})

	if err != nil {
		// If buckets ownership controls are not set it will return an error instead of empty result
		if client.IsAWSError(err, "OwnershipControlsNotFoundError") {
			return nil
		}

		if client.IgnoreAccessDeniedServiceDisabled(err) {
			meta.Logger().Warn().Err(err).Msg("received access denied on GetBucketOwnershipControls")
			return nil
		}

		return err
	}

	if getBucketOwnershipControlOutput == nil {
		return nil
	}

	ownershipControlRules := getBucketOwnershipControlOutput.OwnershipControls.Rules

	if len(ownershipControlRules) == 0 {
		return nil
	}

	stringArray := make([]string, 0, len(ownershipControlRules))

	for _, ownershipControlRule := range ownershipControlRules {
		stringArray = append(stringArray, string(ownershipControlRule.ObjectOwnership))
	}

	resource.OwnershipControls = stringArray
	return nil
}

func isBucketNotFoundError(cl *client.Client, err error) bool {
	if cl.IsNotFoundError(err) {
		return true
	}
	if err.Error() == "bucket not found" {
		return true
	}
	return false
}
