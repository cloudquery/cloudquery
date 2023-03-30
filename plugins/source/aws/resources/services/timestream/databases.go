package timestream

import (
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Databases() *schema.Table {
	tableName := "aws_timestream_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/timestream/latest/developerguide/API_Database.html`,
		Resolver:    fetchTimestreamDatabases,
		Transform:   client.TransformWithStruct(&types.Database{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ingest.timestream"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: fetchDatabaseTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Tables(),
		},
	}
}
