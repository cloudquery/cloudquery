// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

//go:generate mockgen -package=mocks -destination=../mocks/sesv2.go . Sesv2Client
type Sesv2Client interface {
	BatchGetMetricData(context.Context, *sesv2.BatchGetMetricDataInput, ...func(*sesv2.Options)) (*sesv2.BatchGetMetricDataOutput, error)
	GetAccount(context.Context, *sesv2.GetAccountInput, ...func(*sesv2.Options)) (*sesv2.GetAccountOutput, error)
	GetBlacklistReports(context.Context, *sesv2.GetBlacklistReportsInput, ...func(*sesv2.Options)) (*sesv2.GetBlacklistReportsOutput, error)
	GetConfigurationSet(context.Context, *sesv2.GetConfigurationSetInput, ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetOutput, error)
	GetConfigurationSetEventDestinations(context.Context, *sesv2.GetConfigurationSetEventDestinationsInput, ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
	GetContact(context.Context, *sesv2.GetContactInput, ...func(*sesv2.Options)) (*sesv2.GetContactOutput, error)
	GetContactList(context.Context, *sesv2.GetContactListInput, ...func(*sesv2.Options)) (*sesv2.GetContactListOutput, error)
	GetCustomVerificationEmailTemplate(context.Context, *sesv2.GetCustomVerificationEmailTemplateInput, ...func(*sesv2.Options)) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
	GetDedicatedIp(context.Context, *sesv2.GetDedicatedIpInput, ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpOutput, error)
	GetDedicatedIpPool(context.Context, *sesv2.GetDedicatedIpPoolInput, ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpPoolOutput, error)
	GetDedicatedIps(context.Context, *sesv2.GetDedicatedIpsInput, ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpsOutput, error)
	GetDeliverabilityDashboardOptions(context.Context, *sesv2.GetDeliverabilityDashboardOptionsInput, ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityTestReport(context.Context, *sesv2.GetDeliverabilityTestReportInput, ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityTestReportOutput, error)
	GetDomainDeliverabilityCampaign(context.Context, *sesv2.GetDomainDeliverabilityCampaignInput, ...func(*sesv2.Options)) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainStatisticsReport(context.Context, *sesv2.GetDomainStatisticsReportInput, ...func(*sesv2.Options)) (*sesv2.GetDomainStatisticsReportOutput, error)
	GetEmailIdentity(context.Context, *sesv2.GetEmailIdentityInput, ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityOutput, error)
	GetEmailIdentityPolicies(context.Context, *sesv2.GetEmailIdentityPoliciesInput, ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityPoliciesOutput, error)
	GetEmailTemplate(context.Context, *sesv2.GetEmailTemplateInput, ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error)
	GetImportJob(context.Context, *sesv2.GetImportJobInput, ...func(*sesv2.Options)) (*sesv2.GetImportJobOutput, error)
	GetSuppressedDestination(context.Context, *sesv2.GetSuppressedDestinationInput, ...func(*sesv2.Options)) (*sesv2.GetSuppressedDestinationOutput, error)
	ListConfigurationSets(context.Context, *sesv2.ListConfigurationSetsInput, ...func(*sesv2.Options)) (*sesv2.ListConfigurationSetsOutput, error)
	ListContactLists(context.Context, *sesv2.ListContactListsInput, ...func(*sesv2.Options)) (*sesv2.ListContactListsOutput, error)
	ListContacts(context.Context, *sesv2.ListContactsInput, ...func(*sesv2.Options)) (*sesv2.ListContactsOutput, error)
	ListCustomVerificationEmailTemplates(context.Context, *sesv2.ListCustomVerificationEmailTemplatesInput, ...func(*sesv2.Options)) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
	ListDedicatedIpPools(context.Context, *sesv2.ListDedicatedIpPoolsInput, ...func(*sesv2.Options)) (*sesv2.ListDedicatedIpPoolsOutput, error)
	ListDeliverabilityTestReports(context.Context, *sesv2.ListDeliverabilityTestReportsInput, ...func(*sesv2.Options)) (*sesv2.ListDeliverabilityTestReportsOutput, error)
	ListDomainDeliverabilityCampaigns(context.Context, *sesv2.ListDomainDeliverabilityCampaignsInput, ...func(*sesv2.Options)) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
	ListEmailIdentities(context.Context, *sesv2.ListEmailIdentitiesInput, ...func(*sesv2.Options)) (*sesv2.ListEmailIdentitiesOutput, error)
	ListEmailTemplates(context.Context, *sesv2.ListEmailTemplatesInput, ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error)
	ListImportJobs(context.Context, *sesv2.ListImportJobsInput, ...func(*sesv2.Options)) (*sesv2.ListImportJobsOutput, error)
	ListRecommendations(context.Context, *sesv2.ListRecommendationsInput, ...func(*sesv2.Options)) (*sesv2.ListRecommendationsOutput, error)
	ListSuppressedDestinations(context.Context, *sesv2.ListSuppressedDestinationsInput, ...func(*sesv2.Options)) (*sesv2.ListSuppressedDestinationsOutput, error)
	ListTagsForResource(context.Context, *sesv2.ListTagsForResourceInput, ...func(*sesv2.Options)) (*sesv2.ListTagsForResourceOutput, error)
}
