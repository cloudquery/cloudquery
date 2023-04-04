package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func FileCaches() *schema.Table {
	tableName := "aws_fsx_file_caches"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileCache.html`,
		Resolver:    fetchFsxFileCaches,
		Transform:   transformers.TransformWithStruct(&types.FileCache{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "fsx"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveFileCacheTags,
			},
		},
	}
}

func fetchFsxFileCaches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Fsx
	input := fsx.DescribeFileCachesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeFileCachesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.FileCaches
	}
	return nil
}

func resolveFileCacheTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.FileCache)
	cl := meta.(*client.Client)
	svc := cl.Services().Fsx
	var tags []types.Tag
	paginator := fsx.NewListTagsForResourcePaginator(svc, &fsx.ListTagsForResourceInput{ResourceARN: item.ResourceARN})
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		tags = append(tags, result.Tags...)
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
