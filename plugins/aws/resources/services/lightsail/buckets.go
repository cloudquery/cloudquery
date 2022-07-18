package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource buckets --config gen.hcl --output .
func Buckets() *schema.Table {
	return &schema.Table{
		Name:         "aws_lightsail_buckets",
		Description:  "Describes an Amazon Lightsail bucket",
		Resolver:     fetchLightsailBuckets,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lightsail"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "able_to_update_bundle",
				Description: "Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle",
				Type:        schema.TypeBool,
			},
			{
				Name:          "access_log_config_enabled",
				Description:   "A Boolean value that indicates whether bucket access logging is enabled for the bucket",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("AccessLogConfig.Enabled"),
			},
			{
				Name:          "access_log_config_destination",
				Description:   "The name of the bucket where the access logs are saved",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("AccessLogConfig.Destination"),
			},
			{
				Name:          "access_log_config_prefix",
				Description:   "The optional object prefix for the bucket access log",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("AccessLogConfig.Prefix"),
			},
			{
				Name:        "access_rules_allow_public_overrides",
				Description: "A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccessRules.AllowPublicOverrides"),
			},
			{
				Name:        "access_rules_get_object",
				Description: "Specifies the anonymous access to all objects in a bucket",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccessRules.GetObject"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "bundle_id",
				Description: "The ID of the bundle currently applied to the bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the distribution was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "location_availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "location_region_name",
				Description: "The AWS Region name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.RegionName"),
			},
			{
				Name:        "name",
				Description: "The name of the bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "object_versioning",
				Description: "Indicates whether object versioning is enabled for the bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "readonly_access_accounts",
				Description: "An array of strings that specify the Amazon Web Services account IDs that have read-only access to the bucket",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type of the bucket (for example, Bucket)",
				Type:        schema.TypeString,
			},
			{
				Name:          "resources_receiving_access",
				Description:   "An array of objects that describe Lightsail instances that have access to the bucket",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "state_code",
				Description: "The state code of the bucket",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:          "state_message",
				Description:   "A message that describes the state of the bucket",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("State.Message"),
			},
			{
				Name:        "support_code",
				Description: "The support code for a bucket",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the bucket",
				Type:        schema.TypeJSON,
				Resolver:    resolveBucketsTags,
			},
			{
				Name:        "url",
				Description: "The URL of the bucket",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_lightsail_bucket_access_keys",
				Description:   "Describes an access key for an Amazon Lightsail bucket",
				Resolver:      fetchLightsailBucketAccessKeys,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "bucket_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_buckets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "access_key_id",
						Description: "The ID of the access key",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The timestamp when the access key was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "last_used_date",
						Description: "The date and time when the access key was most recently used",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("LastUsed.LastUsedDate"),
					},
					{
						Name:        "last_used_region",
						Description: "The AWS Region where this access key was most recently used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LastUsed.Region"),
					},
					{
						Name:        "last_used_service_name",
						Description: "The name of the AWS service with which this access key was most recently used This value is N/A if the access key has not been used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LastUsed.ServiceName"),
					},
					{
						Name:        "secret_access_key",
						Description: "The secret access key used to sign requests",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the access key",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetBucketsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetBuckets(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Buckets
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func resolveBucketsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Bucket)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchLightsailBucketAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Bucket)
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	input := lightsail.GetBucketAccessKeysInput{
		BucketName: r.Name,
	}
	response, err := svc.GetBucketAccessKeys(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- response.AccessKeys
	return nil
}
