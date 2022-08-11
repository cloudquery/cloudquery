package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Streams() *schema.Table {
	return &schema.Table{
		Name:         "aws_kinesis_streams",
		Description:  "Represents the output for DescribeStreamSummary",
		Resolver:     fetchKinesisStreams,
		Multiplex:    client.ServiceAccountRegionMultiplexer("kinesis"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamARN"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveKinesisStreamTags,
			},
			{
				Name:        "open_shard_count",
				Description: "The number of open shards in the stream",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "retention_period_hours",
				Description: "The current retention period, in hours",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "stream_arn",
				Description: "The Amazon Resource Name (ARN) for the stream being described",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamARN"),
			},
			{
				Name:        "stream_creation_timestamp",
				Description: "The approximate time that the stream was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "stream_name",
				Description: "The name of the stream being described",
				Type:        schema.TypeString,
			},
			{
				Name:        "stream_status",
				Description: "The current status of the stream being described",
				Type:        schema.TypeString,
			},
			{
				Name:        "consumer_count",
				Description: "The number of enhanced fan-out consumers registered with the stream",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "encryption_type",
				Description: "The encryption type used",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_id",
				Description: "The GUID for the customer-managed Amazon Web Services KMS key to use for encryption",
				Type:        schema.TypeString,
			},
			{
				Name:        "stream_mode_details_stream_mode",
				Description: "Specifies the capacity mode to which you want to set your data stream Currently, in Kinesis Data Streams, you can choose between an on-demand capacity mode and a provisioned capacity mode for your data streams",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamModeDetails.StreamMode"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_kinesis_stream_enhanced_monitoring",
				Description: "Represents enhanced metrics types",
				Resolver:    schema.PathTableResolver("EnhancedMonitoring"),
				Columns: []schema.Column{
					{
						Name:        "stream_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_streams table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "shard_level_metrics",
						Description: "List of shard-level metrics",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listKinesisStreams, streamDetail))
}
func ResolveKinesisStreamTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
			return diag.WrapError(err)
		}
		tags = append(tags, output.Tags...)
		if !aws.ToBool(output.HasMoreTags) {
			break
		}
		input.ExclusiveStartTagKey = aws.String(*output.Tags[len(output.Tags)-1].Key)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(tags)))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func listKinesisStreams(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Kinesis
	input := kinesis.ListStreamsInput{}
	for {
		response, err := svc.ListStreams(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
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
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- streamSummary.StreamDescriptionSummary
}
