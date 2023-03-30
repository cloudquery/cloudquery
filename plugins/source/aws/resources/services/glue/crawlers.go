package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Crawlers() *schema.Table {
	tableName := "aws_glue_crawlers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Crawler.html`,
		Resolver:    fetchGlueCrawlers,
		Transform:   transformers.TransformWithStruct(&types.Crawler{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueCrawlerArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueCrawlerTags,
			},
		},
	}
}
