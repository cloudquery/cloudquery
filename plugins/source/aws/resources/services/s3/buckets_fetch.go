package s3

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

// fetchS3BucketsPoolSize describes the amount of go routines that resolve the S3 buckets
const fetchS3BucketsPoolSize = 10

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

func fetchS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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

func fetchS3BucketsWorker(ctx context.Context, meta schema.ClientMeta, buckets <-chan types.Bucket, errs chan<- error, res chan<- any, wg *sync.WaitGroup) {
	defer wg.Done()
	cl := meta.(*client.Client)
	for bucket := range buckets {
		wb := &models.WrappedBucket{Name: bucket.Name, CreationDate: bucket.CreationDate}
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

func resolveS3BucketsAttributes(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket) error {
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

func fetchS3BucketGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	svc := meta.(*client.Client).Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	aclOutput, err := svc.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = region.Str
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
func fetchS3BucketCorsRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	corsOutput, err := svc.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = region.Str
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
func fetchS3BucketEncryptionRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	aclOutput, err := svc.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = region.Str
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

func fetchS3BucketLifecycles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	lifecycleOutput, err := svc.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = region.Str
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
