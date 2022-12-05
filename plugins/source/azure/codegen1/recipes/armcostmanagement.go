// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func Armcostmanagement() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcostmanagement.NewAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "/{scope}/providers/Microsoft.CostManagement/alerts",
		},
		{
			NewFunc: armcostmanagement.NewDimensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "/{scope}/providers/Microsoft.CostManagement/dimensions",
		},
		{
			NewFunc: armcostmanagement.NewExportsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "/{scope}/providers/Microsoft.CostManagement/exports",
		},
		{
			NewFunc: armcostmanagement.NewForecastClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "",
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "",
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "",
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "",
		},
		{
			NewFunc: armcostmanagement.NewGenerateReservationDetailsReportClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "",
		},
		{
			NewFunc: armcostmanagement.NewQueryClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "",
		},
		{
			NewFunc: armcostmanagement.NewViewsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL: "/providers/Microsoft.CostManagement/views",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcostmanagement())
}