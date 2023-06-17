// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

//go:generate mockgen -package=mocks -destination=../mocks/wafv2.go -source=wafv2.go Wafv2Client
type Wafv2Client interface {
	DescribeAllManagedProducts(context.Context, *wafv2.DescribeAllManagedProductsInput, ...func(*wafv2.Options)) (*wafv2.DescribeAllManagedProductsOutput, error)
	DescribeManagedProductsByVendor(context.Context, *wafv2.DescribeManagedProductsByVendorInput, ...func(*wafv2.Options)) (*wafv2.DescribeManagedProductsByVendorOutput, error)
	DescribeManagedRuleGroup(context.Context, *wafv2.DescribeManagedRuleGroupInput, ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error)
	GetDecryptedAPIKey(context.Context, *wafv2.GetDecryptedAPIKeyInput, ...func(*wafv2.Options)) (*wafv2.GetDecryptedAPIKeyOutput, error)
	GetIPSet(context.Context, *wafv2.GetIPSetInput, ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error)
	GetLoggingConfiguration(context.Context, *wafv2.GetLoggingConfigurationInput, ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error)
	GetManagedRuleSet(context.Context, *wafv2.GetManagedRuleSetInput, ...func(*wafv2.Options)) (*wafv2.GetManagedRuleSetOutput, error)
	GetMobileSdkRelease(context.Context, *wafv2.GetMobileSdkReleaseInput, ...func(*wafv2.Options)) (*wafv2.GetMobileSdkReleaseOutput, error)
	GetPermissionPolicy(context.Context, *wafv2.GetPermissionPolicyInput, ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error)
	GetRateBasedStatementManagedKeys(context.Context, *wafv2.GetRateBasedStatementManagedKeysInput, ...func(*wafv2.Options)) (*wafv2.GetRateBasedStatementManagedKeysOutput, error)
	GetRegexPatternSet(context.Context, *wafv2.GetRegexPatternSetInput, ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error)
	GetRuleGroup(context.Context, *wafv2.GetRuleGroupInput, ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error)
	GetSampledRequests(context.Context, *wafv2.GetSampledRequestsInput, ...func(*wafv2.Options)) (*wafv2.GetSampledRequestsOutput, error)
	GetWebACL(context.Context, *wafv2.GetWebACLInput, ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error)
	GetWebACLForResource(context.Context, *wafv2.GetWebACLForResourceInput, ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error)
	ListAPIKeys(context.Context, *wafv2.ListAPIKeysInput, ...func(*wafv2.Options)) (*wafv2.ListAPIKeysOutput, error)
	ListAvailableManagedRuleGroupVersions(context.Context, *wafv2.ListAvailableManagedRuleGroupVersionsInput, ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error)
	ListAvailableManagedRuleGroups(context.Context, *wafv2.ListAvailableManagedRuleGroupsInput, ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	ListIPSets(context.Context, *wafv2.ListIPSetsInput, ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error)
	ListLoggingConfigurations(context.Context, *wafv2.ListLoggingConfigurationsInput, ...func(*wafv2.Options)) (*wafv2.ListLoggingConfigurationsOutput, error)
	ListManagedRuleSets(context.Context, *wafv2.ListManagedRuleSetsInput, ...func(*wafv2.Options)) (*wafv2.ListManagedRuleSetsOutput, error)
	ListMobileSdkReleases(context.Context, *wafv2.ListMobileSdkReleasesInput, ...func(*wafv2.Options)) (*wafv2.ListMobileSdkReleasesOutput, error)
	ListRegexPatternSets(context.Context, *wafv2.ListRegexPatternSetsInput, ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error)
	ListResourcesForWebACL(context.Context, *wafv2.ListResourcesForWebACLInput, ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error)
	ListRuleGroups(context.Context, *wafv2.ListRuleGroupsInput, ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error)
	ListTagsForResource(context.Context, *wafv2.ListTagsForResourceInput, ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error)
	ListWebACLs(context.Context, *wafv2.ListWebACLsInput, ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error)
}
