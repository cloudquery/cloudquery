// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
)

//go:generate mockgen -package=mocks -destination=../mocks/route53domains.go -source=route53domains.go Route53domainsClient
type Route53domainsClient interface {
	GetContactReachabilityStatus(context.Context, *route53domains.GetContactReachabilityStatusInput, ...func(*route53domains.Options)) (*route53domains.GetContactReachabilityStatusOutput, error)
	GetDomainDetail(context.Context, *route53domains.GetDomainDetailInput, ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error)
	GetDomainSuggestions(context.Context, *route53domains.GetDomainSuggestionsInput, ...func(*route53domains.Options)) (*route53domains.GetDomainSuggestionsOutput, error)
	GetOperationDetail(context.Context, *route53domains.GetOperationDetailInput, ...func(*route53domains.Options)) (*route53domains.GetOperationDetailOutput, error)
	ListDomains(context.Context, *route53domains.ListDomainsInput, ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error)
	ListOperations(context.Context, *route53domains.ListOperationsInput, ...func(*route53domains.Options)) (*route53domains.ListOperationsOutput, error)
	ListPrices(context.Context, *route53domains.ListPricesInput, ...func(*route53domains.Options)) (*route53domains.ListPricesOutput, error)
	ListTagsForDomain(context.Context, *route53domains.ListTagsForDomainInput, ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error)
}
