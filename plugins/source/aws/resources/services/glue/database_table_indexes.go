package glue

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func databaseTableIndexes() *schema.Table {
	tableName := "aws_glue_database_table_indexes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_PartitionIndexDescriptor.html`,
		Resolver:    fetchGlueDatabaseTableIndexes,
		Transform:   transformers.TransformWithStruct(&types.PartitionIndexDescriptor{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "database_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("database_arn"),
				PrimaryKey: true,
			},
			{
				Name:       "database_table_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("name"),
				PrimaryKey: true,
			},
			{
				Name:       "index_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("IndexName"),
				PrimaryKey: true,
			},
		},
	}
}
