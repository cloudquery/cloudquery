// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

//go:generate mockgen -package=mocks -destination=../mocks/emr.go . EmrClient
type EmrClient interface {
	DescribeCluster(context.Context, *emr.DescribeClusterInput, ...func(*emr.Options)) (*emr.DescribeClusterOutput, error)
	DescribeJobFlows(context.Context, *emr.DescribeJobFlowsInput, ...func(*emr.Options)) (*emr.DescribeJobFlowsOutput, error)
	DescribeNotebookExecution(context.Context, *emr.DescribeNotebookExecutionInput, ...func(*emr.Options)) (*emr.DescribeNotebookExecutionOutput, error)
	DescribeReleaseLabel(context.Context, *emr.DescribeReleaseLabelInput, ...func(*emr.Options)) (*emr.DescribeReleaseLabelOutput, error)
	DescribeSecurityConfiguration(context.Context, *emr.DescribeSecurityConfigurationInput, ...func(*emr.Options)) (*emr.DescribeSecurityConfigurationOutput, error)
	DescribeStep(context.Context, *emr.DescribeStepInput, ...func(*emr.Options)) (*emr.DescribeStepOutput, error)
	DescribeStudio(context.Context, *emr.DescribeStudioInput, ...func(*emr.Options)) (*emr.DescribeStudioOutput, error)
	GetAutoTerminationPolicy(context.Context, *emr.GetAutoTerminationPolicyInput, ...func(*emr.Options)) (*emr.GetAutoTerminationPolicyOutput, error)
	GetBlockPublicAccessConfiguration(context.Context, *emr.GetBlockPublicAccessConfigurationInput, ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error)
	GetManagedScalingPolicy(context.Context, *emr.GetManagedScalingPolicyInput, ...func(*emr.Options)) (*emr.GetManagedScalingPolicyOutput, error)
	GetStudioSessionMapping(context.Context, *emr.GetStudioSessionMappingInput, ...func(*emr.Options)) (*emr.GetStudioSessionMappingOutput, error)
	ListBootstrapActions(context.Context, *emr.ListBootstrapActionsInput, ...func(*emr.Options)) (*emr.ListBootstrapActionsOutput, error)
	ListClusters(context.Context, *emr.ListClustersInput, ...func(*emr.Options)) (*emr.ListClustersOutput, error)
	ListInstanceFleets(context.Context, *emr.ListInstanceFleetsInput, ...func(*emr.Options)) (*emr.ListInstanceFleetsOutput, error)
	ListInstanceGroups(context.Context, *emr.ListInstanceGroupsInput, ...func(*emr.Options)) (*emr.ListInstanceGroupsOutput, error)
	ListInstances(context.Context, *emr.ListInstancesInput, ...func(*emr.Options)) (*emr.ListInstancesOutput, error)
	ListNotebookExecutions(context.Context, *emr.ListNotebookExecutionsInput, ...func(*emr.Options)) (*emr.ListNotebookExecutionsOutput, error)
	ListReleaseLabels(context.Context, *emr.ListReleaseLabelsInput, ...func(*emr.Options)) (*emr.ListReleaseLabelsOutput, error)
	ListSecurityConfigurations(context.Context, *emr.ListSecurityConfigurationsInput, ...func(*emr.Options)) (*emr.ListSecurityConfigurationsOutput, error)
	ListSteps(context.Context, *emr.ListStepsInput, ...func(*emr.Options)) (*emr.ListStepsOutput, error)
	ListStudioSessionMappings(context.Context, *emr.ListStudioSessionMappingsInput, ...func(*emr.Options)) (*emr.ListStudioSessionMappingsOutput, error)
	ListStudios(context.Context, *emr.ListStudiosInput, ...func(*emr.Options)) (*emr.ListStudiosOutput, error)
}
