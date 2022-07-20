package frontdoor

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/cloudquery/faker/v3"
)

type basicRouteConfigurationOption int

const (
	basicRoute basicRouteConfigurationOption = iota
	basicForwarding
	basicRedirect
)

func fakeRouteConfiguration(t *testing.T) frontdoor.RouteConfiguration {
	var route frontdoor.RouteConfiguration
	if err := faker.FakeData(&route); err != nil {
		t.Fatal(err)
	}
	return route
}

func fakeForwardingConfiguration(t *testing.T) frontdoor.ForwardingConfiguration {
	var forward frontdoor.ForwardingConfiguration
	if err := faker.FakeData(&forward); err != nil {
		t.Fatal(err)
	}
	return forward
}

func fakeRedirectConfiguration(t *testing.T) frontdoor.RedirectConfiguration {
	var redirect frontdoor.RedirectConfiguration
	if err := faker.FakeData(&redirect); err != nil {
		t.Fatal(err)
	}
	return redirect
}

func fakeBasicRouteConfiguration(t *testing.T, option basicRouteConfigurationOption) frontdoor.BasicRouteConfiguration {
	switch option {
	case basicForwarding:
		return fakeForwardingConfiguration(t)
	case basicRedirect:
		return fakeRedirectConfiguration(t)
	default:
		return fakeRouteConfiguration(t)
	}
}

func fakeHeaderAction(t *testing.T) frontdoor.HeaderAction {
	var action frontdoor.HeaderAction
	if err := faker.FakeDataSkipFields(&action, []string{"HeaderActionType"}); err != nil {
		t.Fatal(err)
	}
	action.HeaderActionType = frontdoor.HeaderActionTypeAppend
	return action
}

func fakeRulesEngineMatchCondition(t *testing.T) frontdoor.RulesEngineMatchCondition {
	var condition frontdoor.RulesEngineMatchCondition
	if err := faker.FakeDataSkipFields(&condition,
		[]string{"RulesEngineMatchVariable", "RulesEngineOperator", "Transforms"}); err != nil {
		t.Fatal(err)
	}
	condition.RulesEngineMatchVariable = frontdoor.RulesEngineMatchVariableQueryString
	condition.RulesEngineOperator = frontdoor.RulesEngineOperatorContains
	condition.Transforms = &[]frontdoor.Transform{
		frontdoor.TransformLowercase,
		frontdoor.TransformTrim,
	}
	return condition
}

func fakeRulesEngineRule(t *testing.T, option basicRouteConfigurationOption) frontdoor.RulesEngineRule {
	var rule frontdoor.RulesEngineRule
	if err := faker.FakeDataSkipFields(&rule,
		[]string{"Action", "MatchConditions", "MatchProcessingBehavior"}); err != nil {
		t.Fatal(err)
	}
	rule.Action = &frontdoor.RulesEngineAction{
		RequestHeaderActions: &[]frontdoor.HeaderAction{
			fakeHeaderAction(t),
			fakeHeaderAction(t),
			fakeHeaderAction(t),
		},
		ResponseHeaderActions: &[]frontdoor.HeaderAction{
			fakeHeaderAction(t),
			fakeHeaderAction(t),
			fakeHeaderAction(t),
		},
		RouteConfigurationOverride: fakeBasicRouteConfiguration(t, option),
	}
	rule.MatchConditions = &[]frontdoor.RulesEngineMatchCondition{
		fakeRulesEngineMatchCondition(t),
		fakeRulesEngineMatchCondition(t),
		fakeRulesEngineMatchCondition(t),
	}
	rule.MatchProcessingBehavior = frontdoor.MatchProcessingBehaviorContinue

	return rule
}

func fakeRulesEngine(t *testing.T) frontdoor.RulesEngine {
	var engine frontdoor.RulesEngine
	if err := faker.FakeDataSkipFields(&engine, []string{"RulesEngineProperties"}); err != nil {
		t.Fatal(err)
	}
	engine.RulesEngineProperties = &frontdoor.RulesEngineProperties{
		ResourceState: frontdoor.ResourceStateEnabled,
		Rules: &[]frontdoor.RulesEngineRule{
			fakeRulesEngineRule(t, basicRoute),
			fakeRulesEngineRule(t, basicRedirect),
			fakeRulesEngineRule(t, basicForwarding),
		},
	}
	return engine
}

func fakeRoutingRule(t *testing.T, option basicRouteConfigurationOption) frontdoor.RoutingRule {
	var rule frontdoor.RoutingRule
	if err := faker.FakeDataSkipFields(&rule, []string{"RoutingRuleProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.RoutingRuleProperties
	if err := faker.FakeDataSkipFields(&properties,
		[]string{
			"ResourceState",
			"AcceptedProtocols",
			"EnabledState",
			"RouteConfiguration",
		}); err != nil {
		t.Fatal(err)
	}

	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.AcceptedProtocols = &[]frontdoor.Protocol{frontdoor.ProtocolHTTPS}
	properties.EnabledState = frontdoor.RoutingRuleEnabledStateEnabled
	properties.RouteConfiguration = fakeBasicRouteConfiguration(t, option)

	rule.RoutingRuleProperties = &properties
	return rule
}

func fakeLoadBalancingSettingsModel(t *testing.T) frontdoor.LoadBalancingSettingsModel {
	var model frontdoor.LoadBalancingSettingsModel
	if err := faker.FakeDataSkipFields(&model, []string{"LoadBalancingSettingsProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.LoadBalancingSettingsProperties
	if err := faker.FakeDataSkipFields(&properties, []string{"ResourceState"}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabling
	model.LoadBalancingSettingsProperties = &properties
	return model
}

func fakeHealthProbeSettings(t *testing.T) frontdoor.HealthProbeSettingsModel {
	var model frontdoor.HealthProbeSettingsModel
	if err := faker.FakeDataSkipFields(&model, []string{"HealthProbeSettingsProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.HealthProbeSettingsProperties
	if err := faker.FakeDataSkipFields(&properties, []string{"ResourceState", "Protocol", "HealthProbeMethod", "EnabledState"}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.Protocol = frontdoor.ProtocolHTTPS
	properties.HealthProbeMethod = frontdoor.HealthProbeMethodGET
	properties.EnabledState = frontdoor.HealthProbeEnabledEnabled

	model.HealthProbeSettingsProperties = &properties
	return model
}

func fakeBackend(t *testing.T) frontdoor.Backend {
	var backend frontdoor.Backend
	if err := faker.FakeDataSkipFields(&backend, []string{"PrivateEndpointStatus", "EnabledState"}); err != nil {
		t.Fatal(err)
	}
	backend.PrivateEndpointStatus = frontdoor.PrivateEndpointStatusApproved
	backend.EnabledState = frontdoor.BackendEnabledStateEnabled

	return backend
}

func fakeBackendPools(t *testing.T) frontdoor.BackendPool {
	var pool frontdoor.BackendPool
	if err := faker.FakeDataSkipFields(&pool, []string{"BackendPoolProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.BackendPoolProperties
	if err := faker.FakeDataSkipFields(&properties, []string{"ResourceState", "Backends"}); err != nil {
		t.Fatal(err)
	}

	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.Backends = &[]frontdoor.Backend{
		fakeBackend(t),
		fakeBackend(t),
		fakeBackend(t),
	}
	pool.BackendPoolProperties = &properties
	return pool
}

func fakeCustomHTTPSConfiguration(t *testing.T) *frontdoor.CustomHTTPSConfiguration {
	var config frontdoor.CustomHTTPSConfiguration
	if err := faker.FakeDataSkipFields(&config,
		[]string{"CertificateSource", "MinimumTLSVersion", "CertificateSourceParameters"}); err != nil {
		t.Fatal(err)
	}
	config.CertificateSource = frontdoor.CertificateSourceAzureKeyVault
	config.MinimumTLSVersion = frontdoor.MinimumTLSVersionOneFullStopTwo
	config.CertificateSourceParameters = &frontdoor.CertificateSourceParameters{
		CertificateType: frontdoor.CertificateTypeDedicated,
	}

	return &config
}

func fakeFrontendEndpoint(t *testing.T) frontdoor.FrontendEndpoint {
	var endpoint frontdoor.FrontendEndpoint
	if err := faker.FakeDataSkipFields(&endpoint, []string{"FrontendEndpointProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.FrontendEndpointProperties
	if err := faker.FakeDataSkipFields(&properties,
		[]string{
			"ResourceState",
			"CustomHTTPSProvisioningState",
			"CustomHTTPSProvisioningSubstate",
			"CustomHTTPSConfiguration",
			"SessionAffinityEnabledState",
		}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.CustomHTTPSProvisioningState = frontdoor.CustomHTTPSProvisioningStateEnabling
	properties.CustomHTTPSProvisioningSubstate = frontdoor.CustomHTTPSProvisioningSubstateCertificateDeployed
	properties.CustomHTTPSConfiguration = fakeCustomHTTPSConfiguration(t)
	properties.SessionAffinityEnabledState = frontdoor.SessionAffinityEnabledStateDisabled

	endpoint.FrontendEndpointProperties = &properties
	return endpoint
}

func fakeFrontDoor(t *testing.T) frontdoor.FrontDoor {
	var frontDoor frontdoor.FrontDoor
	if err := faker.FakeDataSkipFields(&frontDoor, []string{"Properties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.Properties
	if err := faker.FakeDataSkipFields(&properties,
		[]string{"ResourceState", "RulesEngines", "RoutingRules", "LoadBalancingSettings", "HealthProbeSettings",
			"BackendPools", "FrontendEndpoints", "BackendPoolsSettings", "EnabledState"}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabled
	properties.RulesEngines = &[]frontdoor.RulesEngine{
		fakeRulesEngine(t),
		fakeRulesEngine(t),
		fakeRulesEngine(t),
	}
	properties.RoutingRules = &[]frontdoor.RoutingRule{
		fakeRoutingRule(t, basicRoute),
		fakeRoutingRule(t, basicRedirect),
		fakeRoutingRule(t, basicForwarding),
	}
	properties.LoadBalancingSettings = &[]frontdoor.LoadBalancingSettingsModel{
		fakeLoadBalancingSettingsModel(t),
		fakeLoadBalancingSettingsModel(t),
	}
	properties.HealthProbeSettings = &[]frontdoor.HealthProbeSettingsModel{
		fakeHealthProbeSettings(t),
		fakeHealthProbeSettings(t),
	}
	properties.BackendPools = &[]frontdoor.BackendPool{
		fakeBackendPools(t),
		fakeBackendPools(t),
		fakeBackendPools(t),
	}
	properties.FrontendEndpoints = &[]frontdoor.FrontendEndpoint{
		fakeFrontendEndpoint(t),
		fakeFrontendEndpoint(t),
		fakeFrontendEndpoint(t),
	}

	var backendPoolsSettings frontdoor.BackendPoolsSettings
	if err := faker.FakeDataSkipFields(&backendPoolsSettings, []string{"EnforceCertificateNameCheck"}); err != nil {
		t.Fatal(err)
	}
	backendPoolsSettings.EnforceCertificateNameCheck = frontdoor.EnforceCertificateNameCheckEnabledStateEnabled
	properties.BackendPoolsSettings = &backendPoolsSettings
	properties.EnabledState = frontdoor.EnabledStateEnabled

	frontDoor.Properties = &properties
	return frontDoor
}
