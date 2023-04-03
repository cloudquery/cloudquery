package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Streams() *schema.Table {
	tableName := "aws_kinesis_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html`,
		Resolver:            fetchKinesisStreams,
		PreResourceResolver: getStream,
		Transform:           transformers.TransformWithStruct(&types.StreamDescriptionSummary{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "kinesis"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKinesisStreamTags,
			},
		},
	}
}

func fetchKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Kinesis
	input := kinesis.ListStreamsInput{}
	// TODO: Use Paginator
	for {
		response, err := svc.ListStreams(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.StreamNames
		if !aws.ToBool(response.HasMoreStreams) {
			break
		}
		input.ExclusiveStartStreamName = aws.String(response.StreamNames[len(response.StreamNames)-1])
	}
	return nil
}

func getStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := c.Services().Kinesis
	streamSummary, err := svc.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	})
	if err != nil {
		return err
	}
	resource.Item = streamSummary.StreamDescriptionSummary
	return nil
}

func resolveKinesisStreamTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kinesis
	summary := resource.Item.(*types.StreamDescriptionSummary)
	input := kinesis.ListTagsForStreamInput{
		StreamName: summary.StreamName,
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForStream(ctx, &input)
		if err != nil {
			return err
		}
		tags = append(tags, output.Tags...)
		if !aws.ToBool(output.HasMoreTags) {
			break
		}
		input.ExclusiveStartTagKey = aws.String(*output.Tags[len(output.Tags)-1].Key)
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
