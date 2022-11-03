// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

//go:generate mockgen -package=mocks -destination=../mocks/cloudfront.go -source=cloudfront.go CloudfrontClient
type CloudfrontClient interface {
	DescribeFunction(context.Context, *cloudfront.DescribeFunctionInput, ...func(*cloudfront.Options)) (*cloudfront.DescribeFunctionOutput, error)
	GetCachePolicy(context.Context, *cloudfront.GetCachePolicyInput, ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyOutput, error)
	GetCachePolicyConfig(context.Context, *cloudfront.GetCachePolicyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyConfigOutput, error)
	GetCloudFrontOriginAccessIdentity(context.Context, *cloudfront.GetCloudFrontOriginAccessIdentityInput, ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
	GetCloudFrontOriginAccessIdentityConfig(context.Context, *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
	GetDistribution(context.Context, *cloudfront.GetDistributionInput, ...func(*cloudfront.Options)) (*cloudfront.GetDistributionOutput, error)
	GetDistributionConfig(context.Context, *cloudfront.GetDistributionConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetDistributionConfigOutput, error)
	GetFieldLevelEncryption(context.Context, *cloudfront.GetFieldLevelEncryptionInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionOutput, error)
	GetFieldLevelEncryptionConfig(context.Context, *cloudfront.GetFieldLevelEncryptionConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
	GetFieldLevelEncryptionProfile(context.Context, *cloudfront.GetFieldLevelEncryptionProfileInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
	GetFieldLevelEncryptionProfileConfig(context.Context, *cloudfront.GetFieldLevelEncryptionProfileConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
	GetFunction(context.Context, *cloudfront.GetFunctionInput, ...func(*cloudfront.Options)) (*cloudfront.GetFunctionOutput, error)
	GetInvalidation(context.Context, *cloudfront.GetInvalidationInput, ...func(*cloudfront.Options)) (*cloudfront.GetInvalidationOutput, error)
	GetKeyGroup(context.Context, *cloudfront.GetKeyGroupInput, ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupOutput, error)
	GetKeyGroupConfig(context.Context, *cloudfront.GetKeyGroupConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupConfigOutput, error)
	GetMonitoringSubscription(context.Context, *cloudfront.GetMonitoringSubscriptionInput, ...func(*cloudfront.Options)) (*cloudfront.GetMonitoringSubscriptionOutput, error)
	GetOriginAccessControl(context.Context, *cloudfront.GetOriginAccessControlInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlOutput, error)
	GetOriginAccessControlConfig(context.Context, *cloudfront.GetOriginAccessControlConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlConfigOutput, error)
	GetOriginRequestPolicy(context.Context, *cloudfront.GetOriginRequestPolicyInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyOutput, error)
	GetOriginRequestPolicyConfig(context.Context, *cloudfront.GetOriginRequestPolicyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
	GetPublicKey(context.Context, *cloudfront.GetPublicKeyInput, ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyOutput, error)
	GetPublicKeyConfig(context.Context, *cloudfront.GetPublicKeyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyConfigOutput, error)
	GetRealtimeLogConfig(context.Context, *cloudfront.GetRealtimeLogConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetRealtimeLogConfigOutput, error)
	GetResponseHeadersPolicy(context.Context, *cloudfront.GetResponseHeadersPolicyInput, ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyOutput, error)
	GetResponseHeadersPolicyConfig(context.Context, *cloudfront.GetResponseHeadersPolicyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyConfigOutput, error)
	GetStreamingDistribution(context.Context, *cloudfront.GetStreamingDistributionInput, ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionOutput, error)
	GetStreamingDistributionConfig(context.Context, *cloudfront.GetStreamingDistributionConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionConfigOutput, error)
	ListCachePolicies(context.Context, *cloudfront.ListCachePoliciesInput, ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error)
	ListCloudFrontOriginAccessIdentities(context.Context, *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, ...func(*cloudfront.Options)) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
	ListConflictingAliases(context.Context, *cloudfront.ListConflictingAliasesInput, ...func(*cloudfront.Options)) (*cloudfront.ListConflictingAliasesOutput, error)
	ListDistributions(context.Context, *cloudfront.ListDistributionsInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsByCachePolicyId(context.Context, *cloudfront.ListDistributionsByCachePolicyIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
	ListDistributionsByKeyGroup(context.Context, *cloudfront.ListDistributionsByKeyGroupInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByKeyGroupOutput, error)
	ListDistributionsByOriginRequestPolicyId(context.Context, *cloudfront.ListDistributionsByOriginRequestPolicyIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
	ListDistributionsByRealtimeLogConfig(context.Context, *cloudfront.ListDistributionsByRealtimeLogConfigInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error)
	ListDistributionsByResponseHeadersPolicyId(context.Context, *cloudfront.ListDistributionsByResponseHeadersPolicyIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput, error)
	ListDistributionsByWebACLId(context.Context, *cloudfront.ListDistributionsByWebACLIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListFieldLevelEncryptionConfigs(context.Context, *cloudfront.ListFieldLevelEncryptionConfigsInput, ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
	ListFieldLevelEncryptionProfiles(context.Context, *cloudfront.ListFieldLevelEncryptionProfilesInput, ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
	ListFunctions(context.Context, *cloudfront.ListFunctionsInput, ...func(*cloudfront.Options)) (*cloudfront.ListFunctionsOutput, error)
	ListInvalidations(context.Context, *cloudfront.ListInvalidationsInput, ...func(*cloudfront.Options)) (*cloudfront.ListInvalidationsOutput, error)
	ListKeyGroups(context.Context, *cloudfront.ListKeyGroupsInput, ...func(*cloudfront.Options)) (*cloudfront.ListKeyGroupsOutput, error)
	ListOriginAccessControls(context.Context, *cloudfront.ListOriginAccessControlsInput, ...func(*cloudfront.Options)) (*cloudfront.ListOriginAccessControlsOutput, error)
	ListOriginRequestPolicies(context.Context, *cloudfront.ListOriginRequestPoliciesInput, ...func(*cloudfront.Options)) (*cloudfront.ListOriginRequestPoliciesOutput, error)
	ListPublicKeys(context.Context, *cloudfront.ListPublicKeysInput, ...func(*cloudfront.Options)) (*cloudfront.ListPublicKeysOutput, error)
	ListRealtimeLogConfigs(context.Context, *cloudfront.ListRealtimeLogConfigsInput, ...func(*cloudfront.Options)) (*cloudfront.ListRealtimeLogConfigsOutput, error)
	ListResponseHeadersPolicies(context.Context, *cloudfront.ListResponseHeadersPoliciesInput, ...func(*cloudfront.Options)) (*cloudfront.ListResponseHeadersPoliciesOutput, error)
	ListStreamingDistributions(context.Context, *cloudfront.ListStreamingDistributionsInput, ...func(*cloudfront.Options)) (*cloudfront.ListStreamingDistributionsOutput, error)
	ListTagsForResource(context.Context, *cloudfront.ListTagsForResourceInput, ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error)
}
