package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
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

func fetchGlueCrawlers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	paginator := glue.NewGetCrawlersPaginator(svc, &glue.GetCrawlersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Crawlers
	}
	return nil
}
func resolveGlueCrawlerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, crawlerARN(cl, aws.ToString(resource.Item.(types.Crawler).Name)))
}
func resolveGlueCrawlerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTagsInput{
		ResourceArn: aws.String(crawlerARN(cl, aws.ToString(resource.Item.(types.Crawler).Name))),
	}

	response, err := svc.GetTags(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}

func crawlerARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("crawler/%s", name),
	}.String()
}
