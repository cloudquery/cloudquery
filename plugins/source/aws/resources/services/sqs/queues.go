package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "redrive_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
			{
				Name:     "redrive_allow_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedriveAllowPolicy"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSqsQueueTags,
			},
		},
	}
}

func fetchSqsQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sqs
	var params sqs.ListQueuesInput
	paginator := sqs.NewListQueuesPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.QueueUrls
	}
	return nil
}

func getQueue(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sqs
	qURL := resource.Item.(string)

	input := sqs.GetQueueAttributesInput{
		QueueUrl:       aws.String(qURL),
		AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll},
	}
	out, err := svc.GetQueueAttributes(ctx, &input)
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
	result, err := svc.ListQueueTags(ctx, &sqs.ListQueueTagsInput{QueueUrl: &q.URL})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
