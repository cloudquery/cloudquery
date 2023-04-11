package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSnsTopicTags,
			},
			{
				Name:     "delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
		},
	}
}

func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	config := sns.ListTopicsInput{}
	paginator := sns.NewListTopicsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Topics
	}
	return nil
}

func getTopic(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	topic := resource.Item.(types.Topic)

	attrs, err := svc.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
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
	tags, err := svc.ListTagsForResource(ctx, &tagParams)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tags.Tags))
}
