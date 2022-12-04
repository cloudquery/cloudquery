// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func Armcostmanagement() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcostmanagement.NewForecastClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewDimensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewQueryClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewGenerateReservationDetailsReportClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewExportsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
		{
			NewFunc: armcostmanagement.NewViewsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcostmanagement())
}