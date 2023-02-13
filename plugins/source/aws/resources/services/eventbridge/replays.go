package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Replays() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_replays",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Replay.html`,
		Resolver:    fetchEventbridgeReplays,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
		Transform:   transformers.TransformWithStruct(&types.Replay{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveReplayArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
