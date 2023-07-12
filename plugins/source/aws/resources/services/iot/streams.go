package iot

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Streams() *schema.Table {
	tableName := "aws_iot_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_StreamInfo.html`,
		Resolver:            fetchIotStreams,
		PreResourceResolver: getStream,
		Transform:           transformers.TransformWithStruct(&types.StreamInfo{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("StreamArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIotStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	paginator := iot.NewListStreamsPaginator(svc, &iot.ListStreamsInput{
		MaxResults: aws.Int32(250),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Streams
	}
	return nil
}

func getStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.DescribeStream(ctx, &iot.DescribeStreamInput{
		StreamId: resource.Item.(types.StreamSummary).StreamId,
	}, func(options *iot.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output.StreamInfo
	return nil
}
