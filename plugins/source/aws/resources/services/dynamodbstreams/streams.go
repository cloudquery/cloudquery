package dynamodbstreams

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Streams() *schema.Table {
	tableName := "aws_dynamodbstreams_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_StreamDescription.html`,
		Resolver:            listStreams,
		PreResourceResolver: describeStream,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "streams.dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.StreamDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Resolver:   schema.PathResolver("StreamArn"),
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
	}
}

func listStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodbstreams

	config := dynamodbstreams.ListStreamsInput{}
	// No paginator available
	for {
		output, err := svc.ListStreams(ctx, &config, func(options *dynamodbstreams.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Streams

		if aws.ToString(output.LastEvaluatedStreamArn) == "" {
			break
		}
		config.ExclusiveStartStreamArn = output.LastEvaluatedStreamArn
	}

	return nil
}

func describeStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodbstreams
	stream := resource.Item.(types.Stream)
	response, err := svc.DescribeStream(ctx, &dynamodbstreams.DescribeStreamInput{StreamArn: stream.StreamArn}, func(options *dynamodbstreams.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = response.StreamDescription
	return nil
}
