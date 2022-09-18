//go:generate mockgen -destination=./mocks/cdn.go -package=mocks . CDNProfilesClient,CDNEndpointsClient,CDNCustomDomainsClient,CDNOriginsClient,CDNOriginGroupsClient,CDNRoutesClient,CDNRuleSetsClient,CDNRulesClient,CDNSecurityPoliciesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	"github.com/Azure/go-autorest/autorest"
)

type CDNProfilesClient interface {
	List(ctx context.Context) (result cdn.ProfileListResultPage, err error)
}

type CDNEndpointsClient interface {
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultPage, err error)
}

type CDNCustomDomainsClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultPage, err error)
}

type CDNOriginsClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultPage, err error)
}

type CDNOriginGroupsClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginGroupListResultPage, err error)
}

type CDNRoutesClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.RouteListResultPage, err error)
}

type CDNRuleSetsClient interface {
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.RuleSetListResultPage, err error)
}

type CDNRulesClient interface {
	ListByRuleSet(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result cdn.RuleListResultPage, err error)
}

type CDNSecurityPoliciesClient interface {
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SecurityPolicyListResultPage, err error)
}

type CDNClient struct {
	Profiles         CDNProfilesClient
	Endpoints        CDNEndpointsClient
	CustomDomains    CDNCustomDomainsClient
	Origins          CDNOriginsClient
	OriginGroups     CDNOriginGroupsClient
	Routes           CDNRoutesClient
	RuleSets         CDNRuleSetsClient
	Rules            CDNRulesClient
	SecurityPolicies CDNSecurityPoliciesClient
}

func NewCDNClient(subscriptionId string, auth autorest.Authorizer) CDNClient {
	p := cdn.NewProfilesClient(subscriptionId)
	p.Authorizer = auth

	e := cdn.NewEndpointsClient(subscriptionId)
	e.Authorizer = auth

	cd := cdn.NewCustomDomainsClient(subscriptionId)
	cd.Authorizer = auth

	o := cdn.NewOriginsClient(subscriptionId)
	o.Authorizer = auth

	og := cdn.NewOriginGroupsClient(subscriptionId)
	og.Authorizer = auth

	r := cdn.NewRoutesClient(subscriptionId)
	r.Authorizer = auth

	rs := cdn.NewRuleSetsClient(subscriptionId)
	rs.Authorizer = auth

	rul := cdn.NewRulesClient(subscriptionId)
	rul.Authorizer = auth

	sp := cdn.NewSecurityPoliciesClient(subscriptionId)
	sp.Authorizer = auth
	return CDNClient{
		Profiles:         p,
		Endpoints:        e,
		CustomDomains:    cd,
		Origins:          o,
		Routes:           r,
		OriginGroups:     og,
		RuleSets:         rs,
		Rules:            rul,
		SecurityPolicies: sp,
	}
}
