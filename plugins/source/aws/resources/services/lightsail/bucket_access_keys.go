// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BucketAccessKeys() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_bucket_access_keys",
		Resolver:  fetchLightsailBucketAccessKeys,
		Multiplex: client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "access_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKeyId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "last_used",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastUsed"),
			},
			{
				Name:     "secret_access_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecretAccessKey"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
