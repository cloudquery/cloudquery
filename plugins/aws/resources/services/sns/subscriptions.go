package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SnsSubscriptions() *schema.Table {
	return &schema.Table{
		Name:          "aws_sns_subscriptions",
		Description:   "A wrapper type for the attributes of an Amazon SNS subscription.",
		Resolver:      fetchSnsSubscriptions,
		Multiplex:     client.ServiceAccountRegionMultiplexer("sns"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"endpoint", "owner", "protocol", "arn", "topic_arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "endpoint",
				Description: "The subscription's endpoint (format depends on the protocol).",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner",
				Description: "The subscription's owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol",
				Description: "The subscription's protocol.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The subscription's ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionArn"),
			},
			{
				Name:        "topic_arn",
				Description: "The ARN of the subscription's topic.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	config := sns.ListSubscriptionsInput{}
	for {
		output, err := svc.ListSubscriptions(ctx, &config, func(o *sns.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Subscriptions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
