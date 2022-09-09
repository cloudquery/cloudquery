package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DeliveryStreams() *schema.Table {
	return &schema.Table{
		Name:        "aws_firehose_delivery_streams",
		Description: "Contains information about a delivery stream",
		Resolver:    fetchFirehoseDeliveryStreams,
		Multiplex:   client.ServiceAccountRegionMultiplexer("firehose"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveFirehoseDeliveryStreamTags,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the delivery stream",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("DeliveryStreamARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "delivery_stream_arn",
				Description: "The Amazon Resource Name (ARN) of the delivery stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamARN"),
			},
			{
				Name:        "delivery_stream_name",
				Description: "The name of the delivery stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_stream_status",
				Description: "The status of the delivery stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_stream_type",
				Description: "The delivery stream type",
				Type:        schema.TypeString,
			},
			{
				Name:        "version_id",
				Description: "Each time the destination is updated for a delivery stream, the version ID is changed, and the current version ID is required when updating the destination This is so that the service knows it is applying the changes to the correct version of the delivery stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_timestamp",
				Description: "The date and time that the delivery stream was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "encryption_config_failure_description_details",
				Description: "A message providing details about the error that caused the failure",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.FailureDescription.Details"),
			},
			{
				Name:        "encryption_config_failure_description_type",
				Description: "The type of error that caused the failure",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.FailureDescription.Type"),
			},
			{
				Name:        "encryption_config_key_arn",
				Description: "If KeyType is CUSTOMER_MANAGED_CMK, this field contains the ARN of the customer managed CMK",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.KeyARN"),
			},
			{
				Name:        "encryption_config_key_type",
				Description: "Indicates the type of customer master key (CMK) that is used for encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.KeyType"),
			},
			{
				Name:        "encryption_config_status",
				Description: "This is the server-side encryption (SSE) status for the delivery stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.Status"),
			},
			{
				Name:        "failure_description_details",
				Description: "A message providing details about the error that caused the failure",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDescription.Details"),
			},
			{
				Name:        "failure_description_type",
				Description: "The type of error that caused the failure",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDescription.Type"),
			},
			{
				Name:        "last_update_timestamp",
				Description: "The date and time that the delivery stream was last updated",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "source_kinesis_stream_delivery_start_timestamp",
				Description: "Kinesis Data Firehose starts retrieving records from the Kinesis data stream starting with this timestamp",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Source.KinesisStreamSourceDescription.DeliveryStartTimestamp"),
			},
			{
				Name:        "source_kinesis_stream_kinesis_stream_arn",
				Description: "The Amazon Resource Name (ARN) of the source Kinesis data stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.KinesisStreamSourceDescription.KinesisStreamARN"),
			},
			{
				Name:        "source_kinesis_stream_role_arn",
				Description: "The ARN of the role used by the source Kinesis data stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.KinesisStreamSourceDescription.RoleARN"),
			},
			{
				Name:     "destinations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Destinations"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFirehoseDeliveryStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listDeliveryStreams, deliveryStreamDetail)
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

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func listDeliveryStreams(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Firehose
	input := firehose.ListDeliveryStreamsInput{}
	for {
		response, err := svc.ListDeliveryStreams(ctx, &input)
		if err != nil {
			return err
		}
		for _, item := range response.DeliveryStreamNames {
			detailChan <- item
		}
		if !aws.ToBool(response.HasMoreDeliveryStreams) {
			break
		}
		input.ExclusiveStartDeliveryStreamName = aws.String(response.DeliveryStreamNames[len(response.DeliveryStreamNames)-1])
	}
	return nil
}
func deliveryStreamDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	streamName := listInfo.(string)
	svc := c.Services().Firehose
	streamSummary, err := svc.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
		DeliveryStreamName: aws.String(streamName),
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- streamSummary.DeliveryStreamDescription
}
