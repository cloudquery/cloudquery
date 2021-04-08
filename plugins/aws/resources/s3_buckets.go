package resources

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/aws/smithy-go"
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
				Name: "creation_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "name",
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
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchS3Buckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().S3
	response, err := svc.ListBuckets(ctx, nil)
	if err != nil {
		return err
	}
	res <- response.Buckets
	return nil
}
func resolveS3BucketsAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	var ae smithy.APIError
	log := meta.Logger()
	r := resource.Item.(types.Bucket)
	svc := meta.(*client.Client).Services().S3
	output, err := svc.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: r.Name,
	})
	if err != nil {
		if errors.As(err, &ae) && ae.ErrorCode() == "NoSuchBucket" {
			// https://aws.amazon.com/premiumsupport/knowledge-center/s3-listing-deleted-bucket/
			// deleted buckets may show up
			log.Debug("Skipping bucket (already deleted)", "bucket", *r.Name)
			return nil
		}
		return err
	}
	bucketRegion := "us-east-1"
	if output.LocationConstraint != "" {
		// This is a weird corner case by AWS API https://github.com/aws/aws-sdk-net/issues/323#issuecomment-196584538
		bucketRegion = string(output.LocationConstraint)
	}
	resource.Set("region", bucketRegion)

	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		return err
	}
	if loggingOutput.LoggingEnabled != nil {
		resource.Set("logging_target_bucket", loggingOutput.LoggingEnabled.TargetBucket)
		resource.Set("logging_target_prefix", loggingOutput.LoggingEnabled.TargetPrefix)
	}

	policyOutput, err := svc.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "NoSuchBucketPolicy") {
		return err
	}
	if policyOutput != nil {
		resource.Set("policy", policyOutput.Policy)
	}

	versioningOutput, err := svc.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		return err
	}
	resource.Set("versioning_status", versioningOutput.Status)
	resource.Set("versioning_mfa_delete", versioningOutput.MFADelete)

	return nil
}
func fetchS3BucketGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.Bucket)
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
	r := parent.Item.(types.Bucket)
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
