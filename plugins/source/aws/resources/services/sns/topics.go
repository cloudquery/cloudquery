package sns

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/mitchellh/mapstructure"
)

func Topics() *schema.Table {
	tableName := "aws_sns_topics"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html`,
		Resolver:            fetchSnsTopics,
		PreResourceResolver: getTopic,
		Transform:           transformers.TransformWithStruct(&models.Topic{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "sns"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveSnsTopicTags,
			},
			{
				Name:     "delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
		},
	}
}

func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sns
	config := sns.ListTopicsInput{}
	paginator := sns.NewListTopicsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sns.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Topics
	}
	return nil
}

func getTopic(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sns
	topic := resource.Item.(types.Topic)

	attrs, err := svc.GetTopicAttributes(ctx,
		&sns.GetTopicAttributesInput{TopicArn: topic.TopicArn},
		func(o *sns.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	t := &models.Topic{Arn: topic.TopicArn}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: t})
	if err != nil {
		return err
	}
	if err := dec.Decode(attrs.Attributes); err != nil {
		return err
	}

	resource.Item = t
	return nil
}

func resolveSnsTopicTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	topic := resource.Item.(*models.Topic)
	cl := meta.(*client.Client)
	svc := cl.Services().Sns
	tagParams := sns.ListTagsForResourceInput{
		ResourceArn: topic.Arn,
	}
	tags, err := svc.ListTagsForResource(ctx, &tagParams, func(o *sns.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tags.Tags))
}
