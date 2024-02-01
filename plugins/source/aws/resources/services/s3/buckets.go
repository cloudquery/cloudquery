package s3

import (
	"context"
	"errors"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Buckets() *schema.Table {
	tableName := "aws_s3_buckets"
	return &schema.Table{
		Name:                "aws_s3_buckets",
		Resolver:            listS3Buckets,
		PreResourceResolver: resolveS3BucketsAttributes,
		Description:         `https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html`,
		Transform:           transformers.TransformWithStruct(&models.WrappedBucket{}),
		Multiplex:           client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveBucketARN(),
				PrimaryKeyComponent: true,
			},
		},

		Relations: []*schema.Table{
			bucketCorsRules(),
			bucketEncryptionRules(),
			bucketGrants(),
			bucketLifecycles(),
			bucketNotificationConfigurations(),
			bucketObjectLockConfigurations(),
			bucketWebsites(),
			bucketLogging(),
			bucketOwnershipControls(),
			bucketReplications(),
			bucketPublicAccessBlock(),
			bucketVersionings(),
			bucketPolicies(),
		},
	}
}

func listS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	response, err := svc.ListBuckets(ctx, nil, func(o *s3.Options) {
		o.Region = listBucketRegion(cl)
	})
	if err != nil {
		return err
	}
	for _, bucket := range response.Buckets {
		res <- &models.WrappedBucket{
			Name:         bucket.Name,
			CreationDate: bucket.CreationDate,
		}
	}
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
	resource := r.Item.(*models.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3

	output, err := svc.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: resource.Name,
	}, func(o *s3.Options) {
		o.Region = listBucketRegion(cl)
	})
	if err != nil {
		if isBucketNotFoundError(cl, err) {
			return nil
		}
		return err
	}
	// AWS does not specify a region if bucket is in us-east-1, so as long as no error we can assume an empty string is us-east-1
	resource.Region = "us-east-1"
	if output != nil && output.LocationConstraint != "" {
		resource.Region = string(output.LocationConstraint)
	}
	if output != nil && output.LocationConstraint == "EU" {
		resource.Region = "eu-west-1"
	}
	var errAll []error

	resolvers := []func(context.Context, schema.ClientMeta, *models.WrappedBucket) error{
		resolveBucketPolicyStatus,
		resolveBucketTagging,
	}
	for _, resolver := range resolvers {
		if err := resolver(ctx, meta, resource); err != nil {
			// If we received any error other than NoSuchBucketError, we return as this indicates that the bucket has been deleted
			// and therefore no other attributes can be resolved
			if isBucketNotFoundError(cl, err) {
				r.Item = resource
				return errors.Join(errAll...)
			}
			// This enables 403 errors to be recorded, but not block subsequent resolver calls
			errAll = append(errAll, err)
		}
	}
	r.Item = resource
	return errors.Join(errAll...)
}

func resolveBucketPolicyStatus(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	policyStatusOutput, err := svc.GetBucketPolicyStatus(ctx, &s3.GetBucketPolicyStatusInput{Bucket: resource.Name}, func(o *s3.Options) {
		o.Region = resource.Region
	})
	// check if we got an error but its access denied we can continue
	if err != nil {
		if client.IsAWSError(err, "NoSuchBucketPolicy") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	if policyStatusOutput != nil {
		resource.PolicyStatus = policyStatusOutput.PolicyStatus
	}
	return nil
}

func resolveBucketTagging(ctx context.Context, meta schema.ClientMeta, resource *models.WrappedBucket) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	taggingOutput, err := svc.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{Bucket: resource.Name}, func(o *s3.Options) {
		o.Region = resource.Region
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
