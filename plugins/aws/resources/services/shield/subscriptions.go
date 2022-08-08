package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource subscriptions --config gen.hcl --output .
func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:          "aws_shield_subscriptions",
		Description:   "Information about the Shield Advanced subscription for an account",
		Resolver:      fetchShieldSubscriptions,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "protection_group_limits_max_protection_groups",
				Description: "The maximum number of protection groups that you can have at one time",
				Type:        schema.TypeInt,
				Resolver:    resolveSubscriptionsProtectionGroupLimitsMaxProtectionGroups,
			},
			{
				Name:        "protection_group_limits_arbitrary_pattern_limits_max_members",
				Description: "The maximum number of resources you can specify for a single arbitrary pattern in a protection group",
				Type:        schema.TypeInt,
				Resolver:    resolveSubscriptionsProtectionGroupLimitsArbitraryPatternLimitsMaxMembers,
			},
			{
				Name:        "protected_resource_type_limits",
				Description: "The maximum number of resource types that you can specify in a protection",
				Type:        schema.TypeJSON,
				Resolver:    resolveSubscriptionsProtectedResourceTypeLimits,
			},
			{
				Name:        "auto_renew",
				Description: "If ENABLED, the subscription will be automatically renewed at the end of the existing subscription period",
				Type:        schema.TypeString,
			},
			{
				Name:        "end_time",
				Description: "The date and time your subscription will end",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "limits",
				Description: "Specifies how many protections of a given type you can create",
				Type:        schema.TypeJSON,
				Resolver:    resolveSubscriptionsLimits,
			},
			{
				Name:        "proactive_engagement_status",
				Description: "If ENABLED, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support",
				Type:        schema.TypeString,
			},
			{
				Name:        "start_time",
				Description: "The start time of the subscription, in Unix time in seconds",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the subscription",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionArn"),
			},
			{
				Name:        "time_commitment_in_seconds",
				Description: "The length, in seconds, of the Shield Advanced subscription for the account",
				Type:        schema.TypeInt,
				Resolver:    resolveSubscriptionsTimeCommitmentInSeconds,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchShieldSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.DescribeSubscriptionInput{}
	output, err := svc.DescribeSubscription(ctx, &config, func(o *shield.Options) {
		o.Region = c.Region
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	res <- output.Subscription
	return nil
}
func resolveSubscriptionsProtectionGroupLimitsMaxProtectionGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	if r.SubscriptionLimits == nil || r.SubscriptionLimits.ProtectionGroupLimits == nil {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, int32(r.SubscriptionLimits.ProtectionGroupLimits.MaxProtectionGroups)))
}
func resolveSubscriptionsProtectionGroupLimitsArbitraryPatternLimitsMaxMembers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	if r.SubscriptionLimits == nil ||
		r.SubscriptionLimits.ProtectionGroupLimits == nil ||
		r.SubscriptionLimits.ProtectionGroupLimits.PatternTypeLimits == nil ||
		r.SubscriptionLimits.ProtectionGroupLimits.PatternTypeLimits.ArbitraryPatternLimits == nil {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, int32(r.SubscriptionLimits.ProtectionGroupLimits.PatternTypeLimits.ArbitraryPatternLimits.MaxMembers)))
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
	return diag.WrapError(resource.Set(c.Name, json))
}
func resolveSubscriptionsLimits(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	json := make(map[string]interface{})
	for _, l := range r.SubscriptionLimits.ProtectionLimits.ProtectedResourceTypeLimits {
		json[*l.Type] = l.Max
	}
	return diag.WrapError(resource.Set(c.Name, json))
}
func resolveSubscriptionsTimeCommitmentInSeconds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Subscription)
	return diag.WrapError(resource.Set(c.Name, int32(r.TimeCommitmentInSeconds)))
}
