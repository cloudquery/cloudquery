package sns

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("SubscriptionArn"),
				PrimaryKey: true,
			},
			{
				Name:     "delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
			{
				Name:     "filter_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("FilterPolicy"),
			},
			{
				Name:     "redrive_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
		},
	}
}

func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sns
	config := sns.ListSubscriptionsInput{}
	paginator := sns.NewListSubscriptionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sns.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Subscriptions
	}
	return nil
}

func getSnsSubscription(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sns
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

	attrs, err := svc.GetSubscriptionAttributes(ctx,
		&sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn},
		func(o *sns.Options) {
			o.Region = cl.Region
		},
	)
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
