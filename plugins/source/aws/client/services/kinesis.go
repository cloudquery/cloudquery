// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

//go:generate mockgen -package=mocks -destination=../mocks/kinesis.go -source=kinesis.go KinesisClient
type KinesisClient interface {
	DescribeLimits(context.Context, *kinesis.DescribeLimitsInput, ...func(*kinesis.Options)) (*kinesis.DescribeLimitsOutput, error)
	DescribeStream(context.Context, *kinesis.DescribeStreamInput, ...func(*kinesis.Options)) (*kinesis.DescribeStreamOutput, error)
	DescribeStreamConsumer(context.Context, *kinesis.DescribeStreamConsumerInput, ...func(*kinesis.Options)) (*kinesis.DescribeStreamConsumerOutput, error)
	DescribeStreamSummary(context.Context, *kinesis.DescribeStreamSummaryInput, ...func(*kinesis.Options)) (*kinesis.DescribeStreamSummaryOutput, error)
	GetRecords(context.Context, *kinesis.GetRecordsInput, ...func(*kinesis.Options)) (*kinesis.GetRecordsOutput, error)
	GetResourcePolicy(context.Context, *kinesis.GetResourcePolicyInput, ...func(*kinesis.Options)) (*kinesis.GetResourcePolicyOutput, error)
	GetShardIterator(context.Context, *kinesis.GetShardIteratorInput, ...func(*kinesis.Options)) (*kinesis.GetShardIteratorOutput, error)
	ListShards(context.Context, *kinesis.ListShardsInput, ...func(*kinesis.Options)) (*kinesis.ListShardsOutput, error)
	ListStreamConsumers(context.Context, *kinesis.ListStreamConsumersInput, ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error)
	ListStreams(context.Context, *kinesis.ListStreamsInput, ...func(*kinesis.Options)) (*kinesis.ListStreamsOutput, error)
	ListTagsForStream(context.Context, *kinesis.ListTagsForStreamInput, ...func(*kinesis.Options)) (*kinesis.ListTagsForStreamOutput, error)
}
