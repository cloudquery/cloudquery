// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/networkfirewall.go -source=networkfirewall.go NetworkfirewallClient
type NetworkfirewallClient interface {
	DescribeFirewall(context.Context, *networkfirewall.DescribeFirewallInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeFirewallOutput, error)
	DescribeFirewallPolicy(context.Context, *networkfirewall.DescribeFirewallPolicyInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeFirewallPolicyOutput, error)
	DescribeLoggingConfiguration(context.Context, *networkfirewall.DescribeLoggingConfigurationInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeLoggingConfigurationOutput, error)
	DescribeResourcePolicy(context.Context, *networkfirewall.DescribeResourcePolicyInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeResourcePolicyOutput, error)
	DescribeRuleGroup(context.Context, *networkfirewall.DescribeRuleGroupInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeRuleGroupOutput, error)
	DescribeRuleGroupMetadata(context.Context, *networkfirewall.DescribeRuleGroupMetadataInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeRuleGroupMetadataOutput, error)
	DescribeTLSInspectionConfiguration(context.Context, *networkfirewall.DescribeTLSInspectionConfigurationInput, ...func(*networkfirewall.Options)) (*networkfirewall.DescribeTLSInspectionConfigurationOutput, error)
	ListFirewallPolicies(context.Context, *networkfirewall.ListFirewallPoliciesInput, ...func(*networkfirewall.Options)) (*networkfirewall.ListFirewallPoliciesOutput, error)
	ListFirewalls(context.Context, *networkfirewall.ListFirewallsInput, ...func(*networkfirewall.Options)) (*networkfirewall.ListFirewallsOutput, error)
	ListRuleGroups(context.Context, *networkfirewall.ListRuleGroupsInput, ...func(*networkfirewall.Options)) (*networkfirewall.ListRuleGroupsOutput, error)
	ListTLSInspectionConfigurations(context.Context, *networkfirewall.ListTLSInspectionConfigurationsInput, ...func(*networkfirewall.Options)) (*networkfirewall.ListTLSInspectionConfigurationsOutput, error)
	ListTagsForResource(context.Context, *networkfirewall.ListTagsForResourceInput, ...func(*networkfirewall.Options)) (*networkfirewall.ListTagsForResourceOutput, error)
}
