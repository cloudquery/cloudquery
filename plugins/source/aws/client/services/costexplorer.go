// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

//go:generate mockgen -package=mocks -destination=../mocks/costexplorer.go -source=costexplorer.go CostexplorerClient
type CostexplorerClient interface {
	DescribeCostCategoryDefinition(context.Context, *costexplorer.DescribeCostCategoryDefinitionInput, ...func(*costexplorer.Options)) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	GetAnomalies(context.Context, *costexplorer.GetAnomaliesInput, ...func(*costexplorer.Options)) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomalyMonitors(context.Context, *costexplorer.GetAnomalyMonitorsInput, ...func(*costexplorer.Options)) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalySubscriptions(context.Context, *costexplorer.GetAnomalySubscriptionsInput, ...func(*costexplorer.Options)) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetCostAndUsage(context.Context, *costexplorer.GetCostAndUsageInput, ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageWithResources(context.Context, *costexplorer.GetCostAndUsageWithResourcesInput, ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostCategories(context.Context, *costexplorer.GetCostCategoriesInput, ...func(*costexplorer.Options)) (*costexplorer.GetCostCategoriesOutput, error)
	GetCostForecast(context.Context, *costexplorer.GetCostForecastInput, ...func(*costexplorer.Options)) (*costexplorer.GetCostForecastOutput, error)
	GetDimensionValues(context.Context, *costexplorer.GetDimensionValuesInput, ...func(*costexplorer.Options)) (*costexplorer.GetDimensionValuesOutput, error)
	GetReservationCoverage(context.Context, *costexplorer.GetReservationCoverageInput, ...func(*costexplorer.Options)) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationPurchaseRecommendation(context.Context, *costexplorer.GetReservationPurchaseRecommendationInput, ...func(*costexplorer.Options)) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationUtilization(context.Context, *costexplorer.GetReservationUtilizationInput, ...func(*costexplorer.Options)) (*costexplorer.GetReservationUtilizationOutput, error)
	GetRightsizingRecommendation(context.Context, *costexplorer.GetRightsizingRecommendationInput, ...func(*costexplorer.Options)) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetSavingsPlanPurchaseRecommendationDetails(context.Context, *costexplorer.GetSavingsPlanPurchaseRecommendationDetailsInput, ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlanPurchaseRecommendationDetailsOutput, error)
	GetSavingsPlansCoverage(context.Context, *costexplorer.GetSavingsPlansCoverageInput, ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansPurchaseRecommendation(context.Context, *costexplorer.GetSavingsPlansPurchaseRecommendationInput, ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansUtilization(context.Context, *costexplorer.GetSavingsPlansUtilizationInput, ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationDetails(context.Context, *costexplorer.GetSavingsPlansUtilizationDetailsInput, ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetTags(context.Context, *costexplorer.GetTagsInput, ...func(*costexplorer.Options)) (*costexplorer.GetTagsOutput, error)
	GetUsageForecast(context.Context, *costexplorer.GetUsageForecastInput, ...func(*costexplorer.Options)) (*costexplorer.GetUsageForecastOutput, error)
	ListCostAllocationTags(context.Context, *costexplorer.ListCostAllocationTagsInput, ...func(*costexplorer.Options)) (*costexplorer.ListCostAllocationTagsOutput, error)
	ListCostCategoryDefinitions(context.Context, *costexplorer.ListCostCategoryDefinitionsInput, ...func(*costexplorer.Options)) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListSavingsPlansPurchaseRecommendationGeneration(context.Context, *costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput, ...func(*costexplorer.Options)) (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput, error)
	ListTagsForResource(context.Context, *costexplorer.ListTagsForResourceInput, ...func(*costexplorer.Options)) (*costexplorer.ListTagsForResourceOutput, error)
}
