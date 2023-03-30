package firehose

import (
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DeliveryStreams() *schema.Table {
	tableName := "aws_firehose_delivery_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html`,
		Resolver:            fetchFirehoseDeliveryStreams,
		PreResourceResolver: getDeliveryStream,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "firehose"),
		Transform:           client.TransformWithStruct(&types.DeliveryStreamDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveFirehoseDeliveryStreamTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryStreamARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
