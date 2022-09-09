package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_buckets",
		Description: "Describes an Amazon Lightsail bucket",
		Resolver:    fetchLightsailBuckets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:          "access_log_config",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("AccessLogConfig"),
				IgnoreInTests: true,
			},
			{
				Name:     "access_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessRules"),
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
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
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
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
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
				Resolver:    client.ResolveTags,
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
						Name:     "last_used",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("LastUsed"),
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
		response, err := svc.GetBuckets(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Buckets
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailBucketAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Bucket)
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	input := lightsail.GetBucketAccessKeysInput{
		BucketName: r.Name,
	}
	response, err := svc.GetBucketAccessKeys(ctx, &input)
	if err != nil {
		return err
	}
	res <- response.AccessKeys
	return nil
}
