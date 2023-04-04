package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/mitchellh/mapstructure"
)

func Subscriptions() *schema.Table {
	tableName := "aws_sns_subscriptions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html`,
		Resolver:            fetchSnsSubscriptions,
		PreResourceResolver: getSnsSubscription,
		Transform:           transformers.TransformWithStruct(&models.Subscription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "sns"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
			{
				Name:     "filter_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FilterPolicy"),
			},
			{
				Name:     "redrive_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
		},
	}
}

func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	config := sns.ListSubscriptionsInput{}
	paginator := sns.NewListSubscriptionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Subscriptions
	}
	return nil
}

func getSnsSubscription(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	item := resource.Item.(types.Subscription)
	s := models.Subscription{
		SubscriptionArn: item.SubscriptionArn,
		Owner:           item.Owner,
		Protocol:        item.Protocol,
		TopicArn:        item.TopicArn,
		Endpoint:        item.Endpoint,
	}
	// Return early if SubscriptionARN is not set because it is still pending
	if aws.ToString(item.SubscriptionArn) == "PendingConfirmation" {
		resource.Item = s
		return nil
	}

	attrs, err := svc.GetSubscriptionAttributes(ctx, &sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn})
	if err != nil {
		return err
	}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &s})
	if err != nil {
		return err
	}
	if err := dec.Decode(attrs.Attributes); err != nil {
		return err
	}
	resource.Item = s
	return nil
}
