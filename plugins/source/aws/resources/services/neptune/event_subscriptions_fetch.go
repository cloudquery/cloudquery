package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchNeptuneEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	input := neptune.DescribeEventSubscriptionsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}
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

func resolveNeptuneEventSubscriptionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.EventSubscription)
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
