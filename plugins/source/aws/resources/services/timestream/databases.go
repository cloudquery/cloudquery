package timestream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Databases() *schema.Table {
	tableName := "aws_timestream_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/timestream/latest/developerguide/API_Database.html`,
		Resolver:    fetchTimestreamDatabases,
		Transform:   transformers.TransformWithStruct(&types.Database{}),
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
			tables(),
		},
	}
}

func fetchTimestreamDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	input := &timestreamwrite.ListDatabasesInput{MaxResults: aws.Int32(20)}
	paginator := timestreamwrite.NewListDatabasesPaginator(meta.(*client.Client).Services().Timestreamwrite, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Databases
	}
	return nil
}

func fetchDatabaseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	output, err := meta.(*client.Client).Services().Timestreamwrite.ListTagsForResource(ctx,
		&timestreamwrite.ListTagsForResourceInput{
			ResourceARN: resource.Item.(types.Database).Arn,
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
