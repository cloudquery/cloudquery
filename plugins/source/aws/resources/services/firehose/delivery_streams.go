package firehose

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func DeliveryStreams() *schema.Table {
	tableName := "aws_firehose_delivery_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html`,
		Resolver:            fetchFirehoseDeliveryStreams,
		PreResourceResolver: getDeliveryStream,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "firehose"),
		Transform:           transformers.TransformWithStruct(&types.DeliveryStreamDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveFirehoseDeliveryStreamTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DeliveryStreamARN"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchFirehoseDeliveryStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceFirehose).Firehose
	input := firehose.ListDeliveryStreamsInput{}
	for {
		response, err := svc.ListDeliveryStreams(ctx, &input, func(options *firehose.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.DeliveryStreamNames
		if !aws.ToBool(response.HasMoreDeliveryStreams) {
			break
		}
		input.ExclusiveStartDeliveryStreamName = aws.String(response.DeliveryStreamNames[len(response.DeliveryStreamNames)-1])
	}
	return nil
}

func getDeliveryStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := cl.Services(client.AWSServiceFirehose).Firehose
	streamSummary, err := svc.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
		DeliveryStreamName: aws.String(streamName),
	}, func(options *firehose.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = streamSummary.DeliveryStreamDescription
	return nil
}

func resolveFirehoseDeliveryStreamTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceFirehose).Firehose
	summary := resource.Item.(*types.DeliveryStreamDescription)
	input := firehose.ListTagsForDeliveryStreamInput{
		DeliveryStreamName: summary.DeliveryStreamName,
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForDeliveryStream(ctx, &input, func(options *firehose.Options) {
			options.Region = cl.Region
		})
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
