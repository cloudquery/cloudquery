package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRdsEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	var input rds.DescribeEventSubscriptionsInput
	for {
		out, err := svc.DescribeEventSubscriptions(ctx, &input)
		if err != nil {
			return err
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
	svc := cl.Services().Rds
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
