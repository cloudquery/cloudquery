//go:generate mockgen -destination=./mocks/cdn.go -package=mocks . ProfilesClient,EndpointsClient,CustomDomainsClient,OriginsClient,OriginGroupsClient,RoutesClient,RuleSetsClient,RulesClient,SecurityPoliciesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	"github.com/Azure/go-autorest/autorest"
)

type ProfilesClient interface {
	List(ctx context.Context) (result cdn.ProfileListResultPage, err error)
}

type EndpointsClient interface {
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultPage, err error)
}

type CustomDomainsClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultPage, err error)
}

type OriginsClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultPage, err error)
}

type OriginGroupsClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginGroupListResultPage, err error)
}

type RoutesClient interface {
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.RouteListResultPage, err error)
}

type RuleSetsClient interface {
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.RuleSetListResultPage, err error)
}

type RulesClient interface {
	ListByRuleSet(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result cdn.RuleListResultPage, err error)
}

type SecurityPoliciesClient interface {
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SecurityPolicyListResultPage, err error)
}

type CDNClient struct {
	Profiles         ProfilesClient
	Endpoints        EndpointsClient
	CustomDomains    CustomDomainsClient
	Origins          OriginsClient
	OriginGroups     OriginGroupsClient
	Routes           RoutesClient
	RuleSets         RuleSetsClient
	Rules            RulesClient
	SecurityPolicies SecurityPoliciesClient
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
