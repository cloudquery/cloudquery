// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DatabaseTableIndexes() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_database_table_indexes",
		Resolver:  fetchGlueDatabaseTableIndexes,
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:     "database_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("database_arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "database_table_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "index_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IndexName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "index_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IndexStatus"),
			},
			{
				Name:     "keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Keys"),
			},
			{
				Name:     "backfill_errors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackfillErrors"),
			},
		},
	}
}
