package cdn

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockCDNProfilesClient(ctrl)

	var profile cdn.Profile
	faker.SetIgnoreInterface(true)
	if err := faker.FakeData(&profile); err != nil {
		t.Fatal(err)
	}
	id := client.FakeResourceGroup
	profile.ID = &id
	m.EXPECT().List(gomock.Any()).Return(
		cdn.NewProfileListResultPage(
			cdn.ProfileListResult{Value: &[]cdn.Profile{profile}},
			func(c context.Context, lr cdn.ProfileListResult) (cdn.ProfileListResult, error) {
				return cdn.ProfileListResult{}, nil
			},
		),
		nil,
	)

	e := mocks.NewMockCDNEndpointsClient(ctrl)
	var endpoint cdn.Endpoint
	if err := faker.FakeDataSkipFields(&endpoint, []string{"EndpointProperties"}); err != nil {
		t.Fatal(err)
	}
	var endpointProperties cdn.EndpointProperties
	if err := faker.FakeDataSkipFields(&endpointProperties, []string{"DeliveryPolicy", "ResourceState", "QueryStringCachingBehavior", "OptimizationType"}); err != nil {
		t.Fatal(err)
	}
	actions := &[]cdn.BasicDeliveryRuleAction{cdn.DeliveryRuleAction{Name: "test"}}
	conditions := &[]cdn.BasicDeliveryRuleCondition{
		cdn.DeliveryRuleCondition{Name: "test"},
	}
	endpoint.EndpointProperties = &endpointProperties
	endpoint.DeliveryPolicy = &cdn.EndpointPropertiesUpdateParametersDeliveryPolicy{
		Description: to.StringPtr("test"),
		Rules: &[]cdn.DeliveryRule{{
			Name:       to.StringPtr("test"),
			Order:      to.Int32Ptr(1),
			Conditions: conditions,
			Actions:    actions,
		},
		},
	}
	e.EXPECT().ListByProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cdn.NewEndpointListResultPage(
			cdn.EndpointListResult{Value: &[]cdn.Endpoint{endpoint}},
			func(c context.Context, lr cdn.EndpointListResult) (cdn.EndpointListResult, error) {
				return cdn.EndpointListResult{}, nil
			},
		),
		nil,
	)

	rs := mocks.NewMockCDNRuleSetsClient(ctrl)
	var ruleSet cdn.RuleSet
	if err := faker.FakeDataSkipFields(&ruleSet, []string{"RuleSetProperties"}); err != nil {
		t.Fatal(err)
	}
	ruleSet.RuleSetProperties = &cdn.RuleSetProperties{DeploymentStatus: "test", ProvisioningState: "test"}
	rs.EXPECT().ListByProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cdn.NewRuleSetListResultPage(
			cdn.RuleSetListResult{Value: &[]cdn.RuleSet{ruleSet}},
			func(c context.Context, lr cdn.RuleSetListResult) (cdn.RuleSetListResult, error) {
				return cdn.RuleSetListResult{}, nil
			},
		),
		nil,
	)

	r := mocks.NewMockCDNRulesClient(ctrl)
	var rule cdn.Rule
	if err := faker.FakeDataSkipFields(&rule, []string{"RuleProperties"}); err != nil {
		t.Fatal(err)
	}
	//todo fake it totaly
	rule.RuleProperties = &cdn.RuleProperties{
		DeploymentStatus:        "test",
		ProvisioningState:       "test",
		MatchProcessingBehavior: "test",
		Actions:                 actions,
		Conditions:              conditions,
		Order:                   to.Int32Ptr(1),
	}

	r.EXPECT().ListByRuleSet(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cdn.NewRuleListResultPage(
			cdn.RuleListResult{Value: &[]cdn.Rule{rule}},
			func(c context.Context, lr cdn.RuleListResult) (cdn.RuleListResult, error) {
				return cdn.RuleListResult{}, nil
			},
		),
		nil,
	)

	s := mocks.NewMockCDNSecurityPoliciesClient(ctrl)
	var securityPolicy cdn.SecurityPolicy
	if err := faker.FakeDataSkipFields(&securityPolicy, []string{"SecurityPolicyProperties"}); err != nil {
		t.Fatal(err)
	}
	securityPolicy.SecurityPolicyProperties = &cdn.SecurityPolicyProperties{
		DeploymentStatus:  "test",
		ProvisioningState: "test",
		Parameters:        cdn.SecurityPolicyParameters{Type: "test"},
	}
	s.EXPECT().ListByProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cdn.NewSecurityPolicyListResultPage(
			cdn.SecurityPolicyListResult{Value: &[]cdn.SecurityPolicy{securityPolicy}},
			func(c context.Context, lr cdn.SecurityPolicyListResult) (cdn.SecurityPolicyListResult, error) {
				return cdn.SecurityPolicyListResult{}, nil
			},
		),
		nil,
	)

	routes := mocks.NewMockCDNRoutesClient(ctrl)
	var route cdn.Route
	if err := faker.FakeData(&route); err != nil {
		t.Fatal(err)
	}
	//todo check  the real data
	route.CompressionSettings = "{\"test\": 1}"
	routes.EXPECT().ListByEndpoint(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cdn.NewRouteListResultPage(
			cdn.RouteListResult{Value: &[]cdn.Route{route}},
			func(c context.Context, lr cdn.RouteListResult) (cdn.RouteListResult, error) {
				return cdn.RouteListResult{}, nil
			},
		),
		nil,
	)

	cd := mocks.NewMockCDNCustomDomainsClient(ctrl)
	var customDomain cdn.CustomDomain
	if err := faker.FakeDataSkipFields(&customDomain, []string{"CustomDomainProperties"}); err != nil {
		t.Fatal(err)
	}
	customDomain.CustomDomainProperties = &cdn.CustomDomainProperties{
		ProvisioningState:               to.StringPtr("test"),
		ValidationData:                  to.StringPtr("test"),
		HostName:                        to.StringPtr("test"),
		CustomHTTPSProvisioningState:    "test",
		CustomHTTPSProvisioningSubstate: "test",
		ResourceState:                   "test",
		CustomHTTPSParameters:           &cdn.ManagedHTTPSParameters{ProtocolType: "test"},
	}
	cd.EXPECT().ListByEndpoint(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cdn.NewCustomDomainListResultPage(
			cdn.CustomDomainListResult{Value: &[]cdn.CustomDomain{customDomain}},
			func(c context.Context, lr cdn.CustomDomainListResult) (cdn.CustomDomainListResult, error) {
				return cdn.CustomDomainListResult{}, nil
			},
		),
		nil,
	)

	return services.Services{
		CDN: services.CDNClient{
			Profiles:         m,
			Endpoints:        e,
			RuleSets:         rs,
			Rules:            r,
			SecurityPolicies: s,
			Routes:           routes,
			CustomDomains:    cd,
		},
	}
}

func TestProfiles(t *testing.T) {
	client.AzureMockTestHelper(t, Profiles(), buildProfilesMock, client.TestOptions{})
}
