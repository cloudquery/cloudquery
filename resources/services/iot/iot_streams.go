package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotStreams() *schema.Table {
	return &schema.Table{
		Name:          "aws_iot_streams",
		Description:   "Information about a stream.",
		Resolver:      fetchIotStreams,
		Multiplex:     client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The stream ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamArn"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iot_stream_files",
				Description: "Represents a file to stream.",
				Resolver:    fetchIotStreamFiles,
				Columns: []schema.Column{
					{
						Name:        "stream_cq_id",
						Description: "Unique CloudQuery ID of aws_iot_streams table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "file_id",
						Description: "The file ID.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "s3_location_bucket",
						Description: "The S3 bucket.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3Location.Bucket"),
					},
					{
						Name:        "s3_location_key",
						Description: "The S3 key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3Location.Key"),
					},
					{
						Name:        "s3_location_version",
						Description: "The S3 bucket version.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3Location.Version"),
					},
				},
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
		response, err := svc.ListStreams(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
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
				return err
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
func fetchIotStreamFiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	i, ok := parent.Item.(*types.StreamInfo)
	if !ok {
		return fmt.Errorf("expected types.StreamSummary but got %T", parent.Item)
	}

	res <- i.Files
	return nil
}
