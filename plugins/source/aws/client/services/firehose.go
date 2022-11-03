// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

//go:generate mockgen -package=mocks -destination=../mocks/firehose.go . FirehoseClient
type FirehoseClient interface {
	DescribeDeliveryStream(context.Context, *firehose.DescribeDeliveryStreamInput, ...func(*firehose.Options)) (*firehose.DescribeDeliveryStreamOutput, error)
	ListDeliveryStreams(context.Context, *firehose.ListDeliveryStreamsInput, ...func(*firehose.Options)) (*firehose.ListDeliveryStreamsOutput, error)
	ListTagsForDeliveryStream(context.Context, *firehose.ListTagsForDeliveryStreamInput, ...func(*firehose.Options)) (*firehose.ListTagsForDeliveryStreamOutput, error)
}
