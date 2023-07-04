// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
)

//go:generate mockgen -package=mocks -destination=../mocks/resiliencehub.go -source=resiliencehub.go ResiliencehubClient
type ResiliencehubClient interface {
	DescribeApp(context.Context, *resiliencehub.DescribeAppInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppOutput, error)
	DescribeAppAssessment(context.Context, *resiliencehub.DescribeAppAssessmentInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppAssessmentOutput, error)
	DescribeAppVersion(context.Context, *resiliencehub.DescribeAppVersionInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionOutput, error)
	DescribeAppVersionAppComponent(context.Context, *resiliencehub.DescribeAppVersionAppComponentInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionAppComponentOutput, error)
	DescribeAppVersionResource(context.Context, *resiliencehub.DescribeAppVersionResourceInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionResourceOutput, error)
	DescribeAppVersionResourcesResolutionStatus(context.Context, *resiliencehub.DescribeAppVersionResourcesResolutionStatusInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput, error)
	DescribeAppVersionTemplate(context.Context, *resiliencehub.DescribeAppVersionTemplateInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionTemplateOutput, error)
	DescribeDraftAppVersionResourcesImportStatus(context.Context, *resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput, error)
	DescribeResiliencyPolicy(context.Context, *resiliencehub.DescribeResiliencyPolicyInput, ...func(*resiliencehub.Options)) (*resiliencehub.DescribeResiliencyPolicyOutput, error)
	ListAlarmRecommendations(context.Context, *resiliencehub.ListAlarmRecommendationsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAlarmRecommendationsOutput, error)
	ListAppAssessments(context.Context, *resiliencehub.ListAppAssessmentsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppAssessmentsOutput, error)
	ListAppComponentCompliances(context.Context, *resiliencehub.ListAppComponentCompliancesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppComponentCompliancesOutput, error)
	ListAppComponentRecommendations(context.Context, *resiliencehub.ListAppComponentRecommendationsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppComponentRecommendationsOutput, error)
	ListAppInputSources(context.Context, *resiliencehub.ListAppInputSourcesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppInputSourcesOutput, error)
	ListAppVersionAppComponents(context.Context, *resiliencehub.ListAppVersionAppComponentsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionAppComponentsOutput, error)
	ListAppVersionResourceMappings(context.Context, *resiliencehub.ListAppVersionResourceMappingsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionResourceMappingsOutput, error)
	ListAppVersionResources(context.Context, *resiliencehub.ListAppVersionResourcesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionResourcesOutput, error)
	ListAppVersions(context.Context, *resiliencehub.ListAppVersionsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionsOutput, error)
	ListApps(context.Context, *resiliencehub.ListAppsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListAppsOutput, error)
	ListRecommendationTemplates(context.Context, *resiliencehub.ListRecommendationTemplatesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListRecommendationTemplatesOutput, error)
	ListResiliencyPolicies(context.Context, *resiliencehub.ListResiliencyPoliciesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListResiliencyPoliciesOutput, error)
	ListSopRecommendations(context.Context, *resiliencehub.ListSopRecommendationsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListSopRecommendationsOutput, error)
	ListSuggestedResiliencyPolicies(context.Context, *resiliencehub.ListSuggestedResiliencyPoliciesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListSuggestedResiliencyPoliciesOutput, error)
	ListTagsForResource(context.Context, *resiliencehub.ListTagsForResourceInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListTagsForResourceOutput, error)
	ListTestRecommendations(context.Context, *resiliencehub.ListTestRecommendationsInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListTestRecommendationsOutput, error)
	ListUnsupportedAppVersionResources(context.Context, *resiliencehub.ListUnsupportedAppVersionResourcesInput, ...func(*resiliencehub.Options)) (*resiliencehub.ListUnsupportedAppVersionResourcesOutput, error)
}
