package sqs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
