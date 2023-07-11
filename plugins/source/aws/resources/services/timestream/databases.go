package timestream

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: fetchDatabaseTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			tables(),
		},
	}
}

func fetchTimestreamDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Timestreamwrite
	// This should be removed once https://github.com/aws/aws-sdk-go-v2/issues/2163 is fixed
	if cl.AWSConfig != nil && cl.AWSConfig.Region != cl.Region {
		awsCfg := cl.AWSConfig.Copy()
		awsCfg.Region = cl.Region
		svc = timestreamwrite.NewFromConfig(awsCfg)
	}
	input := &timestreamwrite.ListDatabasesInput{MaxResults: aws.Int32(20)}
	paginator := timestreamwrite.NewListDatabasesPaginator(svc, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(o *timestreamwrite.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Databases
	}
	return nil
}

func fetchDatabaseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Timestreamwrite

	output, err := svc.ListTagsForResource(ctx,
		&timestreamwrite.ListTagsForResourceInput{
			ResourceARN: resource.Item.(types.Database).Arn,
		},
		func(o *timestreamwrite.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
