// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/ecr.go -source=ecr.go EcrClient
type EcrClient interface {
	BatchGetImage(context.Context, *ecr.BatchGetImageInput, ...func(*ecr.Options)) (*ecr.BatchGetImageOutput, error)
	BatchGetRepositoryScanningConfiguration(context.Context, *ecr.BatchGetRepositoryScanningConfigurationInput, ...func(*ecr.Options)) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error)
	DescribeImageReplicationStatus(context.Context, *ecr.DescribeImageReplicationStatusInput, ...func(*ecr.Options)) (*ecr.DescribeImageReplicationStatusOutput, error)
	DescribeImageScanFindings(context.Context, *ecr.DescribeImageScanFindingsInput, ...func(*ecr.Options)) (*ecr.DescribeImageScanFindingsOutput, error)
	DescribeImages(context.Context, *ecr.DescribeImagesInput, ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error)
	DescribePullThroughCacheRules(context.Context, *ecr.DescribePullThroughCacheRulesInput, ...func(*ecr.Options)) (*ecr.DescribePullThroughCacheRulesOutput, error)
	DescribeRegistry(context.Context, *ecr.DescribeRegistryInput, ...func(*ecr.Options)) (*ecr.DescribeRegistryOutput, error)
	DescribeRepositories(context.Context, *ecr.DescribeRepositoriesInput, ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error)
	GetAuthorizationToken(context.Context, *ecr.GetAuthorizationTokenInput, ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error)
	GetDownloadUrlForLayer(context.Context, *ecr.GetDownloadUrlForLayerInput, ...func(*ecr.Options)) (*ecr.GetDownloadUrlForLayerOutput, error)
	GetLifecyclePolicy(context.Context, *ecr.GetLifecyclePolicyInput, ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyPreview(context.Context, *ecr.GetLifecyclePolicyPreviewInput, ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyPreviewOutput, error)
	GetRegistryPolicy(context.Context, *ecr.GetRegistryPolicyInput, ...func(*ecr.Options)) (*ecr.GetRegistryPolicyOutput, error)
	GetRegistryScanningConfiguration(context.Context, *ecr.GetRegistryScanningConfigurationInput, ...func(*ecr.Options)) (*ecr.GetRegistryScanningConfigurationOutput, error)
	GetRepositoryPolicy(context.Context, *ecr.GetRepositoryPolicyInput, ...func(*ecr.Options)) (*ecr.GetRepositoryPolicyOutput, error)
	ListImages(context.Context, *ecr.ListImagesInput, ...func(*ecr.Options)) (*ecr.ListImagesOutput, error)
	ListTagsForResource(context.Context, *ecr.ListTagsForResourceInput, ...func(*ecr.Options)) (*ecr.ListTagsForResourceOutput, error)
}
