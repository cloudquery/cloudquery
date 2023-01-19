package sqs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Queues() *schema.Table {
	return &schema.Table{
		Name:                "aws_sqs_queues",
		Description:         `https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html`,
		Resolver:            fetchSqsQueues,
		PreResourceResolver: getQueue,
		Transform:           transformers.TransformWithStruct(&models.Queue{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("sqs"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
