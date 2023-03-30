package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EventSubscriptions() *schema.Table {
	tableName := "aws_redshift_event_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_EventSubscription.html`,
		Resolver:    fetchEventSubscriptions,
		Transform:   transformers.TransformWithStruct(&types.EventSubscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveEventSubscriptionARN,
				Description: `ARN of the event subscription.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	var params redshift.DescribeEventSubscriptionsInput
	params.MaxRecords = aws.Int32(100)
	for {
		result, err := svc.DescribeEventSubscriptions(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.EventSubscriptionsList
		if aws.ToString(result.Marker) == "" {
			break
		}
		params.Marker = result.Marker
	}
	return nil
}

func resolveEventSubscriptionARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	sub := resource.Item.(types.EventSubscription)
	return resource.Set(c.Name, eventSubscriptionARN(cl, *sub.CustSubscriptionId))
}

func eventSubscriptionARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.RedshiftService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("eventsubscription:%s", name),
	}.String()
}
