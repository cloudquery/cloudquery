// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

//go:generate mockgen -package=mocks -destination=../mocks/waf.go . WafClient
type WafClient interface {
	GetByteMatchSet(context.Context, *waf.GetByteMatchSetInput, ...func(*waf.Options)) (*waf.GetByteMatchSetOutput, error)
	GetChangeToken(context.Context, *waf.GetChangeTokenInput, ...func(*waf.Options)) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenStatus(context.Context, *waf.GetChangeTokenStatusInput, ...func(*waf.Options)) (*waf.GetChangeTokenStatusOutput, error)
	GetGeoMatchSet(context.Context, *waf.GetGeoMatchSetInput, ...func(*waf.Options)) (*waf.GetGeoMatchSetOutput, error)
	GetIPSet(context.Context, *waf.GetIPSetInput, ...func(*waf.Options)) (*waf.GetIPSetOutput, error)
	GetLoggingConfiguration(context.Context, *waf.GetLoggingConfigurationInput, ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error)
	GetPermissionPolicy(context.Context, *waf.GetPermissionPolicyInput, ...func(*waf.Options)) (*waf.GetPermissionPolicyOutput, error)
	GetRateBasedRule(context.Context, *waf.GetRateBasedRuleInput, ...func(*waf.Options)) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleManagedKeys(context.Context, *waf.GetRateBasedRuleManagedKeysInput, ...func(*waf.Options)) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRegexMatchSet(context.Context, *waf.GetRegexMatchSetInput, ...func(*waf.Options)) (*waf.GetRegexMatchSetOutput, error)
	GetRegexPatternSet(context.Context, *waf.GetRegexPatternSetInput, ...func(*waf.Options)) (*waf.GetRegexPatternSetOutput, error)
	GetRule(context.Context, *waf.GetRuleInput, ...func(*waf.Options)) (*waf.GetRuleOutput, error)
	GetRuleGroup(context.Context, *waf.GetRuleGroupInput, ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error)
	GetSampledRequests(context.Context, *waf.GetSampledRequestsInput, ...func(*waf.Options)) (*waf.GetSampledRequestsOutput, error)
	GetSizeConstraintSet(context.Context, *waf.GetSizeConstraintSetInput, ...func(*waf.Options)) (*waf.GetSizeConstraintSetOutput, error)
	GetSqlInjectionMatchSet(context.Context, *waf.GetSqlInjectionMatchSetInput, ...func(*waf.Options)) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetWebACL(context.Context, *waf.GetWebACLInput, ...func(*waf.Options)) (*waf.GetWebACLOutput, error)
	GetXssMatchSet(context.Context, *waf.GetXssMatchSetInput, ...func(*waf.Options)) (*waf.GetXssMatchSetOutput, error)
	ListActivatedRulesInRuleGroup(context.Context, *waf.ListActivatedRulesInRuleGroupInput, ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListByteMatchSets(context.Context, *waf.ListByteMatchSetsInput, ...func(*waf.Options)) (*waf.ListByteMatchSetsOutput, error)
	ListGeoMatchSets(context.Context, *waf.ListGeoMatchSetsInput, ...func(*waf.Options)) (*waf.ListGeoMatchSetsOutput, error)
	ListIPSets(context.Context, *waf.ListIPSetsInput, ...func(*waf.Options)) (*waf.ListIPSetsOutput, error)
	ListLoggingConfigurations(context.Context, *waf.ListLoggingConfigurationsInput, ...func(*waf.Options)) (*waf.ListLoggingConfigurationsOutput, error)
	ListRateBasedRules(context.Context, *waf.ListRateBasedRulesInput, ...func(*waf.Options)) (*waf.ListRateBasedRulesOutput, error)
	ListRegexMatchSets(context.Context, *waf.ListRegexMatchSetsInput, ...func(*waf.Options)) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexPatternSets(context.Context, *waf.ListRegexPatternSetsInput, ...func(*waf.Options)) (*waf.ListRegexPatternSetsOutput, error)
	ListRuleGroups(context.Context, *waf.ListRuleGroupsInput, ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error)
	ListRules(context.Context, *waf.ListRulesInput, ...func(*waf.Options)) (*waf.ListRulesOutput, error)
	ListSizeConstraintSets(context.Context, *waf.ListSizeConstraintSetsInput, ...func(*waf.Options)) (*waf.ListSizeConstraintSetsOutput, error)
	ListSqlInjectionMatchSets(context.Context, *waf.ListSqlInjectionMatchSetsInput, ...func(*waf.Options)) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSubscribedRuleGroups(context.Context, *waf.ListSubscribedRuleGroupsInput, ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListTagsForResource(context.Context, *waf.ListTagsForResourceInput, ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error)
	ListWebACLs(context.Context, *waf.ListWebACLsInput, ...func(*waf.Options)) (*waf.ListWebACLsOutput, error)
	ListXssMatchSets(context.Context, *waf.ListXssMatchSetsInput, ...func(*waf.Options)) (*waf.ListXssMatchSetsOutput, error)
}
