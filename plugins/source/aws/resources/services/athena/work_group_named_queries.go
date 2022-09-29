// Code generated by codegen; DO NOT EDIT.

package athena

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkGroupNamedQueries() *schema.Table {
	return &schema.Table{
		Name:      "aws_athena_work_group_named_queries",
		Resolver:  fetchAthenaWorkGroupNamedQueries,
		Multiplex: client.ServiceAccountRegionMultiplexer("athena"),
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
				Name:     "work_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "database",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Database"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "query_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryString"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "named_query_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NamedQueryId"),
			},
			{
				Name:     "work_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkGroup"),
			},
		},
	}
}
