package kinesis

import (
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Streams() *schema.Table {
	tableName := "aws_kinesis_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html`,
		Resolver:            fetchKinesisStreams,
		PreResourceResolver: getStream,
		Transform:           client.TransformWithStruct(&types.StreamDescriptionSummary{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "kinesis"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKinesisStreamTags,
			},
		},
	}
}
