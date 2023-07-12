package sqs

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/mitchellh/mapstructure"
)

func Queues() *schema.Table {
	tableName := "aws_sqs_queues"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html`,
		Resolver:            fetchSqsQueues,
		PreResourceResolver: getQueue,
		Transform:           transformers.TransformWithStruct(&models.Queue{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "sqs"),
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
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "redrive_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
			{
				Name:     "redrive_allow_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("RedriveAllowPolicy"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveSqsQueueTags,
			},
		},
	}
}

func fetchSqsQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sqs
	var params sqs.ListQueuesInput
	paginator := sqs.NewListQueuesPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sqs.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.QueueUrls
	}
	return nil
}

func getQueue(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sqs
	qURL := resource.Item.(string)

	input := sqs.GetQueueAttributesInput{
		QueueUrl:       aws.String(qURL),
		AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll},
	}
	out, err := svc.GetQueueAttributes(ctx, &input, func(o *sqs.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	q := &models.Queue{URL: qURL}
	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: q})
	if err != nil {
		return err
	}
	if err := d.Decode(out.Attributes); err != nil {
		return err
	}

	resource.Item = q
	return nil
}

func resolveSqsQueueTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sqs
	q := resource.Item.(*models.Queue)
	result, err := svc.ListQueueTags(ctx, &sqs.ListQueueTagsInput{QueueUrl: &q.URL}, func(o *sqs.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
