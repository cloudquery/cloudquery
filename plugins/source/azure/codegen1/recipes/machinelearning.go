// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning"

func Armmachinelearning() []*Table {
	tables := []*Table{
		{
			NewFunc: armmachinelearning.NewWorkspaceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces/skus",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmachinelearning())
}