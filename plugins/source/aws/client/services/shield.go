// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/shield.go -source=shield.go ShieldClient
type ShieldClient interface {
	DescribeAttack(context.Context, *shield.DescribeAttackInput, ...func(*shield.Options)) (*shield.DescribeAttackOutput, error)
	DescribeAttackStatistics(context.Context, *shield.DescribeAttackStatisticsInput, ...func(*shield.Options)) (*shield.DescribeAttackStatisticsOutput, error)
	DescribeDRTAccess(context.Context, *shield.DescribeDRTAccessInput, ...func(*shield.Options)) (*shield.DescribeDRTAccessOutput, error)
	DescribeEmergencyContactSettings(context.Context, *shield.DescribeEmergencyContactSettingsInput, ...func(*shield.Options)) (*shield.DescribeEmergencyContactSettingsOutput, error)
	DescribeProtection(context.Context, *shield.DescribeProtectionInput, ...func(*shield.Options)) (*shield.DescribeProtectionOutput, error)
	DescribeProtectionGroup(context.Context, *shield.DescribeProtectionGroupInput, ...func(*shield.Options)) (*shield.DescribeProtectionGroupOutput, error)
	DescribeSubscription(context.Context, *shield.DescribeSubscriptionInput, ...func(*shield.Options)) (*shield.DescribeSubscriptionOutput, error)
	GetSubscriptionState(context.Context, *shield.GetSubscriptionStateInput, ...func(*shield.Options)) (*shield.GetSubscriptionStateOutput, error)
	ListAttacks(context.Context, *shield.ListAttacksInput, ...func(*shield.Options)) (*shield.ListAttacksOutput, error)
	ListProtectionGroups(context.Context, *shield.ListProtectionGroupsInput, ...func(*shield.Options)) (*shield.ListProtectionGroupsOutput, error)
	ListProtections(context.Context, *shield.ListProtectionsInput, ...func(*shield.Options)) (*shield.ListProtectionsOutput, error)
	ListResourcesInProtectionGroup(context.Context, *shield.ListResourcesInProtectionGroupInput, ...func(*shield.Options)) (*shield.ListResourcesInProtectionGroupOutput, error)
	ListTagsForResource(context.Context, *shield.ListTagsForResourceInput, ...func(*shield.Options)) (*shield.ListTagsForResourceOutput, error)
}
