// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "aws_elasticache_users",
		Resolver:  fetchElasticacheUsers,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "access_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessString"),
			},
			{
				Name:     "authentication",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Authentication"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "minimum_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumEngineVersion"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "user_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("UserGroupIds"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserId"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
			},
		},
	}
}
