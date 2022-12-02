// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

//go:generate mockgen -package=mocks -destination=../mocks/inspector.go -source=inspector.go InspectorClient
type InspectorClient interface {
	DescribeAssessmentRuns(context.Context, *inspector.DescribeAssessmentRunsInput, ...func(*inspector.Options)) (*inspector.DescribeAssessmentRunsOutput, error)
	DescribeAssessmentTargets(context.Context, *inspector.DescribeAssessmentTargetsInput, ...func(*inspector.Options)) (*inspector.DescribeAssessmentTargetsOutput, error)
	DescribeAssessmentTemplates(context.Context, *inspector.DescribeAssessmentTemplatesInput, ...func(*inspector.Options)) (*inspector.DescribeAssessmentTemplatesOutput, error)
	DescribeCrossAccountAccessRole(context.Context, *inspector.DescribeCrossAccountAccessRoleInput, ...func(*inspector.Options)) (*inspector.DescribeCrossAccountAccessRoleOutput, error)
	DescribeExclusions(context.Context, *inspector.DescribeExclusionsInput, ...func(*inspector.Options)) (*inspector.DescribeExclusionsOutput, error)
	DescribeFindings(context.Context, *inspector.DescribeFindingsInput, ...func(*inspector.Options)) (*inspector.DescribeFindingsOutput, error)
	DescribeResourceGroups(context.Context, *inspector.DescribeResourceGroupsInput, ...func(*inspector.Options)) (*inspector.DescribeResourceGroupsOutput, error)
	DescribeRulesPackages(context.Context, *inspector.DescribeRulesPackagesInput, ...func(*inspector.Options)) (*inspector.DescribeRulesPackagesOutput, error)
	GetAssessmentReport(context.Context, *inspector.GetAssessmentReportInput, ...func(*inspector.Options)) (*inspector.GetAssessmentReportOutput, error)
	GetExclusionsPreview(context.Context, *inspector.GetExclusionsPreviewInput, ...func(*inspector.Options)) (*inspector.GetExclusionsPreviewOutput, error)
	GetTelemetryMetadata(context.Context, *inspector.GetTelemetryMetadataInput, ...func(*inspector.Options)) (*inspector.GetTelemetryMetadataOutput, error)
	ListAssessmentRunAgents(context.Context, *inspector.ListAssessmentRunAgentsInput, ...func(*inspector.Options)) (*inspector.ListAssessmentRunAgentsOutput, error)
	ListAssessmentRuns(context.Context, *inspector.ListAssessmentRunsInput, ...func(*inspector.Options)) (*inspector.ListAssessmentRunsOutput, error)
	ListAssessmentTargets(context.Context, *inspector.ListAssessmentTargetsInput, ...func(*inspector.Options)) (*inspector.ListAssessmentTargetsOutput, error)
	ListAssessmentTemplates(context.Context, *inspector.ListAssessmentTemplatesInput, ...func(*inspector.Options)) (*inspector.ListAssessmentTemplatesOutput, error)
	ListEventSubscriptions(context.Context, *inspector.ListEventSubscriptionsInput, ...func(*inspector.Options)) (*inspector.ListEventSubscriptionsOutput, error)
	ListExclusions(context.Context, *inspector.ListExclusionsInput, ...func(*inspector.Options)) (*inspector.ListExclusionsOutput, error)
	ListFindings(context.Context, *inspector.ListFindingsInput, ...func(*inspector.Options)) (*inspector.ListFindingsOutput, error)
	ListRulesPackages(context.Context, *inspector.ListRulesPackagesInput, ...func(*inspector.Options)) (*inspector.ListRulesPackagesOutput, error)
	ListTagsForResource(context.Context, *inspector.ListTagsForResourceInput, ...func(*inspector.Options)) (*inspector.ListTagsForResourceOutput, error)
}
