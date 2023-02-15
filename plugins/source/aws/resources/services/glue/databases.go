package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_databases",
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Database.html`,
		Resolver:    fetchGlueDatabases,
		Transform:   transformers.TransformWithStruct(&types.Database{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueDatabaseArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueDatabaseTags,
			},
		},

		Relations: []*schema.Table{
			DatabaseTables(),
		},
	}
}
