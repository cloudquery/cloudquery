package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func SnsSubscriptions() *schema.Table {
	return &schema.Table{
		Name:         "aws_sns_subscriptions",
		Resolver:     fetchSnsSubscriptions,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "endpoint",
				Type: schema.TypeString,
			},
			{
				Name: "owner",
				Type: schema.TypeString,
			},
			{
				Name: "protocol",
				Type: schema.TypeString,
			},
			{
				Name: "subscription_arn",
				Type: schema.TypeString,
			},
			{
				Name: "topic_arn",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SNS
	config := sns.ListSubscriptionsInput{}
	for {
		output, err := svc.ListSubscriptions(ctx, &config, func(o *sns.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.Subscriptions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
