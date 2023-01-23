package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DatabaseTableIndexes() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_database_table_indexes",
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_PartitionIndexDescriptor.html`,
		Resolver:    fetchGlueDatabaseTableIndexes,
		Transform:   transformers.TransformWithStruct(&types.PartitionIndexDescriptor{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
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
		},
	}
}
