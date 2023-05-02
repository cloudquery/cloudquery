package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func DbProxies() *schema.Table {
	tableName := "aws_db_proxies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBProxy.html`,
		Resolver:    fetchDbProxies,
		Transform:   transformers.TransformWithStruct(&types.DBProxy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBProxyArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsDbProxyTags,
			},
		},
	}
}

func fetchDbProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds
	input := rds.DescribeDBProxiesInput{}
	paginator := rds.NewDescribeDBProxiesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBProxies
	}
	return nil
}

func resolveRdsDbProxyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBProxy)
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBProxyArn}, func(options *rds.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
