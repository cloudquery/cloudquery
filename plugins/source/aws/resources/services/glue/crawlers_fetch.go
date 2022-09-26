package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueCrawlers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.GetCrawlersInput{}
	for {
		output, err := svc.GetCrawlers(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Crawlers

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveGlueCrawlerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(crawlerARN(cl, aws.ToString(resource.Item.(types.Crawler).Name)))
	return resource.Set(c.Name, arn)
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

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func crawlerARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "crawler", name)
}
