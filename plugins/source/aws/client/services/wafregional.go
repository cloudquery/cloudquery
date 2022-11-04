// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

//go:generate mockgen -package=mocks -destination=../mocks/wafregional.go -source=wafregional.go WafregionalClient
type WafregionalClient interface {
	GetByteMatchSet(context.Context, *wafregional.GetByteMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetByteMatchSetOutput, error)
	GetChangeToken(context.Context, *wafregional.GetChangeTokenInput, ...func(*wafregional.Options)) (*wafregional.GetChangeTokenOutput, error)
	GetChangeTokenStatus(context.Context, *wafregional.GetChangeTokenStatusInput, ...func(*wafregional.Options)) (*wafregional.GetChangeTokenStatusOutput, error)
	GetGeoMatchSet(context.Context, *wafregional.GetGeoMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetGeoMatchSetOutput, error)
	GetIPSet(context.Context, *wafregional.GetIPSetInput, ...func(*wafregional.Options)) (*wafregional.GetIPSetOutput, error)
	GetLoggingConfiguration(context.Context, *wafregional.GetLoggingConfigurationInput, ...func(*wafregional.Options)) (*wafregional.GetLoggingConfigurationOutput, error)
	GetPermissionPolicy(context.Context, *wafregional.GetPermissionPolicyInput, ...func(*wafregional.Options)) (*wafregional.GetPermissionPolicyOutput, error)
	GetRateBasedRule(context.Context, *wafregional.GetRateBasedRuleInput, ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error)
	GetRateBasedRuleManagedKeys(context.Context, *wafregional.GetRateBasedRuleManagedKeysInput, ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleManagedKeysOutput, error)
	GetRegexMatchSet(context.Context, *wafregional.GetRegexMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetRegexMatchSetOutput, error)
	GetRegexPatternSet(context.Context, *wafregional.GetRegexPatternSetInput, ...func(*wafregional.Options)) (*wafregional.GetRegexPatternSetOutput, error)
	GetRule(context.Context, *wafregional.GetRuleInput, ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error)
	GetRuleGroup(context.Context, *wafregional.GetRuleGroupInput, ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error)
	GetSampledRequests(context.Context, *wafregional.GetSampledRequestsInput, ...func(*wafregional.Options)) (*wafregional.GetSampledRequestsOutput, error)
	GetSizeConstraintSet(context.Context, *wafregional.GetSizeConstraintSetInput, ...func(*wafregional.Options)) (*wafregional.GetSizeConstraintSetOutput, error)
	GetSqlInjectionMatchSet(context.Context, *wafregional.GetSqlInjectionMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetSqlInjectionMatchSetOutput, error)
	GetWebACL(context.Context, *wafregional.GetWebACLInput, ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error)
	GetWebACLForResource(context.Context, *wafregional.GetWebACLForResourceInput, ...func(*wafregional.Options)) (*wafregional.GetWebACLForResourceOutput, error)
	GetXssMatchSet(context.Context, *wafregional.GetXssMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetXssMatchSetOutput, error)
	ListActivatedRulesInRuleGroup(context.Context, *wafregional.ListActivatedRulesInRuleGroupInput, ...func(*wafregional.Options)) (*wafregional.ListActivatedRulesInRuleGroupOutput, error)
	ListByteMatchSets(context.Context, *wafregional.ListByteMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListByteMatchSetsOutput, error)
	ListGeoMatchSets(context.Context, *wafregional.ListGeoMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListGeoMatchSetsOutput, error)
	ListIPSets(context.Context, *wafregional.ListIPSetsInput, ...func(*wafregional.Options)) (*wafregional.ListIPSetsOutput, error)
	ListLoggingConfigurations(context.Context, *wafregional.ListLoggingConfigurationsInput, ...func(*wafregional.Options)) (*wafregional.ListLoggingConfigurationsOutput, error)
	ListRateBasedRules(context.Context, *wafregional.ListRateBasedRulesInput, ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error)
	ListRegexMatchSets(context.Context, *wafregional.ListRegexMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListRegexMatchSetsOutput, error)
	ListRegexPatternSets(context.Context, *wafregional.ListRegexPatternSetsInput, ...func(*wafregional.Options)) (*wafregional.ListRegexPatternSetsOutput, error)
	ListResourcesForWebACL(context.Context, *wafregional.ListResourcesForWebACLInput, ...func(*wafregional.Options)) (*wafregional.ListResourcesForWebACLOutput, error)
	ListRuleGroups(context.Context, *wafregional.ListRuleGroupsInput, ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error)
	ListRules(context.Context, *wafregional.ListRulesInput, ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error)
	ListSizeConstraintSets(context.Context, *wafregional.ListSizeConstraintSetsInput, ...func(*wafregional.Options)) (*wafregional.ListSizeConstraintSetsOutput, error)
	ListSqlInjectionMatchSets(context.Context, *wafregional.ListSqlInjectionMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListSqlInjectionMatchSetsOutput, error)
	ListSubscribedRuleGroups(context.Context, *wafregional.ListSubscribedRuleGroupsInput, ...func(*wafregional.Options)) (*wafregional.ListSubscribedRuleGroupsOutput, error)
	ListTagsForResource(context.Context, *wafregional.ListTagsForResourceInput, ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error)
	ListWebACLs(context.Context, *wafregional.ListWebACLsInput, ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error)
	ListXssMatchSets(context.Context, *wafregional.ListXssMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListXssMatchSetsOutput, error)
}
