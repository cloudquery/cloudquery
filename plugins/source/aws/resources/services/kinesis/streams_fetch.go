package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listKinesisStreams, streamDetail)
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

func listKinesisStreams(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Kinesis
	input := kinesis.ListStreamsInput{}
	for {
		response, err := svc.ListStreams(ctx, &input)
		if err != nil {
			return err
		}
		for _, item := range response.StreamNames {
			detailChan <- item
		}
		if !aws.ToBool(response.HasMoreStreams) {
			break
		}
		input.ExclusiveStartStreamName = aws.String(response.StreamNames[len(response.StreamNames)-1])
	}
	return nil
}
func streamDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	streamName := listInfo.(string)
	svc := c.Services().Kinesis
	streamSummary, err := svc.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- streamSummary.StreamDescriptionSummary
}
