package models

import (
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
)

type FirewallPolicyWrapper struct {
	*types.FirewallPolicy
	*types.FirewallPolicyResponse
}

type RuleGroupWrapper struct {
	*types.RuleGroup
	*types.RuleGroupResponse
}

type FirewallWrapper struct {
	*types.FirewallStatus
	*types.Firewall
}
type TLSInspectionConfigurationWrapper struct {
	*types.TLSInspectionConfiguration
	*types.TLSInspectionConfigurationResponse
}
