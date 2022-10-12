package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFirehoseDeliveryStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Firehose
	input := firehose.ListDeliveryStreamsInput{}
	for {
		response, err := svc.ListDeliveryStreams(ctx, &input)
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
	c := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := c.Services().Firehose
	streamSummary, err := svc.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
		DeliveryStreamName: aws.String(streamName),
	})
	if err != nil {
		return err
	}
	resource.Item = streamSummary.DeliveryStreamDescription
	return nil
}

func resolveFirehoseDeliveryStreamTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Firehose
	summary := resource.Item.(*types.DeliveryStreamDescription)
	input := firehose.ListTagsForDeliveryStreamInput{
		DeliveryStreamName: summary.DeliveryStreamName,
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForDeliveryStream(ctx, &input)
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
