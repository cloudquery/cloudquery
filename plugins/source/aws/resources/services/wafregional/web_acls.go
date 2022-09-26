// Code generated by codegen; DO NOT EDIT.

package wafregional

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WebAcls() *schema.Table {
	return &schema.Table{
		Name:      "aws_wafregional_web_acls",
		Resolver:  fetchWafregionalWebAcls,
		Multiplex: client.ServiceAccountRegionMultiplexer("waf-regional"),
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
				Resolver: schema.PathResolver("WebACLArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalWebACLTags,
				Description: `Web ACL tags.`,
			},
			{
				Name:     "default_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultAction"),
			},
			{
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rules"),
			},
			{
				Name:     "web_acl_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLId"),
			},
			{
				Name:     "metric_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricName"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}
