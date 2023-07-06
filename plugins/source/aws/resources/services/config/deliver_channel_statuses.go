package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func deliveryChannelStatuses() *schema.Table {
	tableName := "aws_config_delivery_channel_statuses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeDeliveryChannelStatus.html`,
		Resolver:    fetchDeliveryChannelStatuses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   transformers.TransformWithStruct(&types.DeliveryChannelStatus{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchDeliveryChannelStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ruleDetail := parent.Item.(types.DeliveryChannel)
	cl := meta.(*client.Client)
	svc := cl.Services().Configservice

	input := &configservice.DescribeDeliveryChannelStatusInput{
		DeliveryChannelNames: []string{aws.ToString(ruleDetail.Name)},
	}

	response, err := svc.DescribeDeliveryChannelStatus(ctx, input, func(options *configservice.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.DeliveryChannelsStatus
	return nil
}
