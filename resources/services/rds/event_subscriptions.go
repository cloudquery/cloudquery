package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsEventSubscriptions() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_event_subscriptions",
		Description:  "Contains the results of a successful invocation of the DescribeEventSubscriptions action.",
		Resolver:     fetchRdsEventSubscriptions,
		Multiplex:    client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "cust_subscription_id",
				Description: "The RDS event notification subscription Id.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_aws_id",
				Description: "The AWS customer account associated with the RDS event notification subscription.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "A Boolean value indicating if the subscription is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "event_categories_list",
				Description: "A list of event categories for the RDS event notification subscription.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the event subscription.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventSubscriptionArn"),
			},
			{
				Name:        "sns_topic_arn",
				Description: "The topic ARN of the RDS event notification subscription.",
				Type:        schema.TypeString,
			},
			{
				Name:          "source_ids_list",
				Description:   "A list of source IDs for the RDS event notification subscription.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "source_type",
				Description: "The source type for the RDS event notification subscription.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the RDS event notification subscription",
				Type:        schema.TypeString,
			},
			{
				Name:        "subscription_creation_time",
				Description: "The time the RDS event notification subscription was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "List of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRDSEventSubscriptionTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	var input rds.DescribeEventSubscriptionsInput
	for {
		out, err := svc.DescribeEventSubscriptions(ctx, &input, func(o *rds.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- out.EventSubscriptionsList
		if aws.ToString(out.Marker) == "" {
			break
		}
		input.Marker = out.Marker
	}
	return nil
}

func resolveRDSEventSubscriptionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.EventSubscription)
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn}, func(o *rds.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	tags := make(map[string]string, len(out.TagList))
	for _, t := range out.TagList {
		tags[aws.ToString(t.Key)] = aws.ToString(t.Value)
	}
	return resource.Set(c.Name, tags)
}
