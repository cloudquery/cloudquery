// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_buckets",
		Resolver:  fetchLightsailBuckets,
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "able_to_update_bundle",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AbleToUpdateBundle"),
			},
			{
				Name:     "access_log_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessLogConfig"),
			},
			{
				Name:     "access_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessRules"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "bundle_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BundleId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "object_versioning",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObjectVersioning"),
			},
			{
				Name:     "readonly_access_accounts",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReadonlyAccessAccounts"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "resources_receiving_access",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourcesReceivingAccess"),
			},
			{
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
			},
		},

		Relations: []*schema.Table{
			BucketAccessKeys(),
		},
	}
}
