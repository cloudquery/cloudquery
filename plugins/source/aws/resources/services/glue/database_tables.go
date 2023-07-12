package glue

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func databaseTables() *schema.Table {
	tableName := "aws_glue_database_tables"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Table.html`,
		Resolver:    fetchGlueDatabaseTables,
		Transform:   transformers.TransformWithStruct(&types.Table{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "database_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			databaseTableIndexes(),
		},
	}
}
