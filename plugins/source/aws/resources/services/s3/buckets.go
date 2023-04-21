package s3

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Buckets() *schema.Table {
	tableName := "aws_s3_buckets"
	return &schema.Table{
		Name:                "aws_s3_buckets",
		Resolver:            listS3Buckets,
		PreResourceResolver: resolveS3BucketsAttributes,
		Transform:           transformers.TransformWithStruct(&models.WrappedBucket{}),
		Multiplex:           client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveBucketARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			bucketEncryptionRules(),
			bucketLifecycles(),
			bucketGrants(),
			bucketCorsRules(),
			bucketWebsites(),
		},
	}
}

func listS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().S3
	response, err := svc.ListBuckets(ctx, nil, func(options *s3.Options) {
		options.Region = listBucketRegion(cl)
	})
	if err != nil {
		return err
	}
	res <- response.Buckets
	return nil
}

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

func resolveS3BucketsAttributes(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	bucket := r.Item.(types.Bucket)
	resource := &models.WrappedBucket{Name: bucket.Name, CreationDate: bucket.CreationDate}
	c := meta.(*client.Client)
	mgr := c.Services().S3manager

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

	resolvers := []func(context.Context, schema.ClientMeta, *models.WrappedBucket, string) error{
		resolveBucketLogging,
		resolveBucketPolicy,
		resolveBucketVersioning,
		resolveBucketPublicAccessBlock,
		resolveBucketReplication,
		resolveBucketTagging,
		resolveBucketOwnershipControls,
	}
	for _, resolver := range resolvers {
		if err := resolver(ctx, meta, resource, resource.Region); err != nil {
			r.Item = resource
			if isBucketNotFoundError(c, err) {
				return nil
			}
			return err
		}
	}
	r.Item = resource
	return nil
}

func resolveBucketLogging(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
	svc := meta.(*client.Client).Services().S3
	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
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

func resolveBucketPolicy(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
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
			return nil
		}
		return err
	}
	if policyOutput == nil || policyOutput.Policy == nil {
		return nil
	}
	var p map[string]any
	err = json.Unmarshal([]byte(*policyOutput.Policy), &p)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON policy: %v", err)
	}
	resource.Policy = p
	return nil
}

func resolveBucketVersioning(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	versioningOutput, err := svc.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	resource.VersioningStatus = versioningOutput.Status
	resource.VersioningMfaDelete = versioningOutput.MFADelete
	return nil
}

func resolveBucketPublicAccessBlock(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
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

func resolveBucketReplication(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
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

func resolveBucketTagging(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
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
	resource.Tags = tags
	return nil
}

func resolveBucketOwnershipControls(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket, bucketRegion string) error {
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

func resolveBucketARN() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.S3Service, func(resource *schema.Resource) ([]string, error) {
		return []string{*resource.Item.(*models.WrappedBucket).Name}, nil
	})
}
