// Code generated by codegen; DO NOT EDIT.

package config

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConformancePacks() *schema.Table {
	return &schema.Table{
		Name:      "aws_config_conformance_packs",
		Resolver:  fetchConfigConformancePacks,
		Multiplex: client.ServiceAccountRegionMultiplexer("config"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConformancePackArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "conformance_pack_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConformancePackId"),
			},
			{
				Name:     "conformance_pack_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConformancePackName"),
			},
			{
				Name:     "created_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedBy"),
			},
			{
				Name:     "delivery_s3_bucket",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryS3Bucket"),
			},
			{
				Name:     "delivery_s3_key_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryS3KeyPrefix"),
			},
			{
				Name:     "last_update_requested_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdateRequestedTime"),
			},
		},

		Relations: []*schema.Table{
			ConformancePackRuleCompliances(),
		},
	}
}
