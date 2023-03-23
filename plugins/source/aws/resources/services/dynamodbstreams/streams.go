package dynamodbstreams

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Resolver: schema.PathResolver("StreamArn"),
				Type:     schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{},
	}
}

func listStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Dynamodbstreams

	config := dynamodbstreams.ListStreamsInput{}
	for {
		output, err := svc.ListStreams(ctx, &config)
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
	svc := meta.(*client.Client).Services().Dynamodbstreams
	stream := resource.Item.(types.Stream)
	response, err := svc.DescribeStream(ctx, &dynamodbstreams.DescribeStreamInput{StreamArn: stream.StreamArn})
	if err != nil {
		return err
	}

	resource.Item = response.StreamDescription
	return nil
}
