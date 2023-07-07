package eventbridge

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Replays() *schema.Table {
	tableName := "aws_eventbridge_replays"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DescribeReplay.html`,
		Resolver:            fetchReplays,
		PreResourceResolver: getReplay,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "events"),
		Transform:           transformers.TransformWithStruct(&eventbridge.DescribeReplayOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ReplayArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchReplays(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input eventbridge.ListReplaysInput
	cl := meta.(*client.Client)
	svc := cl.Services().Eventbridge
	// No paginator available
	for {
		response, err := svc.ListReplays(ctx, &input, func(options *eventbridge.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Replays
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getReplay(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Eventbridge

	replay := resource.Item.(types.Replay)

	out, err := svc.DescribeReplay(ctx, &eventbridge.DescribeReplayInput{
		ReplayName: replay.ReplayName,
	}, func(options *eventbridge.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = out
	return nil
}
