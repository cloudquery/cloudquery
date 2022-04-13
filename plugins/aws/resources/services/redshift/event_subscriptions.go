package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EventSubscriptions() *schema.Table {
	return &schema.Table{
		Name:         "aws_redshift_event_subscriptions",
		Description:  "Describes event subscriptions.",
		Resolver:     fetchRedshiftEventSubscriptions,
		Multiplex:    client.ServiceAccountRegionMultiplexer("redshift"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "id",
				Description: "The name of the Amazon Redshift event notification subscription.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CustSubscriptionId"),
			},
			{
				Name:        "customer_aws_id",
				Description: "The AWS customer account associated with the Amazon Redshift event notification subscription.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "A boolean value indicating whether the subscription is enabled; true indicates that the subscription is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "event_categories_list",
				Description: "The list of Amazon Redshift event categories specified in the event notification subscription",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "severity",
				Description: "The event severity specified in the Amazon Redshift event notification subscription",
				Type:        schema.TypeString,
			},
			{
				Name:        "sns_topic_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon SNS topic used by the event notification subscription.",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_ids_list",
				Description: "A list of the sources that publish events to the Amazon Redshift event notification subscription.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "source_type",
				Description: "The source type of the events returned by the Amazon Redshift event notification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the Amazon Redshift event notification subscription.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subscription_creation_time",
				Description: "The date and time the Amazon Redshift event notification subscription was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "tags",
				Description: "Tags",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchRedshiftEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	var params redshift.DescribeEventSubscriptionsInput
	params.MaxRecords = aws.Int32(100)
	for {
		result, err := svc.DescribeEventSubscriptions(ctx, &params, func(o *redshift.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.EventSubscriptionsList
		if aws.ToString(result.Marker) == "" {
			break
		}
		params.Marker = result.Marker
	}
	return nil
}
