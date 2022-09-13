package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchShieldSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.DescribeSubscriptionInput{}
	output, err := svc.DescribeSubscription(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- output.Subscription
	return nil
}
func resolveSubscriptionsProtectionGroupLimitsMaxProtectionGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	if r.SubscriptionLimits == nil || r.SubscriptionLimits.ProtectionGroupLimits == nil {
		return nil
	}
	return resource.Set(c.Name, int32(r.SubscriptionLimits.ProtectionGroupLimits.MaxProtectionGroups))
}
func resolveSubscriptionsProtectionGroupLimitsArbitraryPatternLimitsMaxMembers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	if r.SubscriptionLimits == nil ||
		r.SubscriptionLimits.ProtectionGroupLimits == nil ||
		r.SubscriptionLimits.ProtectionGroupLimits.PatternTypeLimits == nil ||
		r.SubscriptionLimits.ProtectionGroupLimits.PatternTypeLimits.ArbitraryPatternLimits == nil {
		return nil
	}
	return resource.Set(c.Name, int32(r.SubscriptionLimits.ProtectionGroupLimits.PatternTypeLimits.ArbitraryPatternLimits.MaxMembers))
}
func resolveSubscriptionsProtectedResourceTypeLimits(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	json := make(map[string]interface{})
	if r.SubscriptionLimits == nil || r.SubscriptionLimits.ProtectionLimits == nil {
		return nil
	}
	for _, l := range r.SubscriptionLimits.ProtectionLimits.ProtectedResourceTypeLimits {
		json[*l.Type] = l.Max
	}
	return resource.Set(c.Name, json)
}
func resolveSubscriptionsLimits(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	json := make(map[string]interface{})
	for _, l := range r.SubscriptionLimits.ProtectionLimits.ProtectedResourceTypeLimits {
		json[*l.Type] = l.Max
	}
	return resource.Set(c.Name, json)
}
func resolveSubscriptionsTimeCommitmentInSeconds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	return resource.Set(c.Name, int32(r.TimeCommitmentInSeconds))
}
