package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Streams() *schema.Table {
	tableName := "aws_iot_streams"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_StreamInfo.html`,
		Resolver:    fetchIotStreams,
		Transform:   transformers.TransformWithStruct(&types.StreamInfo{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Iot
	paginator := iot.NewListStreamsPaginator(svc, &iot.ListStreamsInput{
		MaxResults: aws.Int32(250),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, s := range page.Streams {
			// TODO: Handle resolution in parallel with PreResourceResolver
			stream, err := svc.DescribeStream(ctx, &iot.DescribeStreamInput{
				StreamId: s.StreamId,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				// A single `Describe` call error should not end resolving of table
				c.Logger().Warn().Err(err).Msg("failed to describe stream")
				continue
			}
			res <- stream.StreamInfo
		}
	}
	return nil
}
