package cloudtrail

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Channels() *schema.Table {
	tableName := "aws_cloudtrail_channels"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetChannel.html`,
		Resolver:            fetchCloudtrailChannels,
		PreResourceResolver: getChannel,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cloudtrail"),
		Transform:           transformers.TransformWithStruct(&cloudtrail.GetChannelOutput{}, transformers.WithSkipFields("ResponseMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ChannelArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchCloudtrailChannels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail
	paginator := cloudtrail.NewListChannelsPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudtrail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Channels
	}
	return nil
}

func getChannel(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail
	item := resource.Item.(types.Channel)

	params := cloudtrail.GetChannelInput{
		Channel: item.ChannelArn,
	}
	resp, err := svc.GetChannel(ctx, &params, func(options *cloudtrail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = resp
	return nil
}
