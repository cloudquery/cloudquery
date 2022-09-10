package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotStreams() *schema.Table {
	return &schema.Table{
		Name:          "aws_iot_streams",
		Description:   "Information about a stream.",
		Resolver:      fetchIotStreams,
		Multiplex:     client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "created_at",
				Description: "The date when the stream was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the stream.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_updated_at",
				Description: "The date when the stream was last updated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "role_arn",
				Description: "An IAM role IoT assumes to access your S3 files.",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The stream ARN.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("StreamArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "id",
				Description: "The stream ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamId"),
			},
			{
				Name:        "version",
				Description: "The stream version.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("StreamVersion"),
			},
			{
				Name:        "files",
				Type: 			schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := iot.ListStreamsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListStreams(ctx, &input)
		if err != nil {
			return err
		}
		for _, s := range response.Streams {
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
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
