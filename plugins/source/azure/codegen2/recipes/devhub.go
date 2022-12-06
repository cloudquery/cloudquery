// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"

func Armdevhub() []Table {
	tables := []Table{
		{
			Name:           "workflow",
			Struct:         &armdevhub.Workflow{},
			ResponseStruct: &armdevhub.WorkflowClientListResponse{},
			Client:         &armdevhub.WorkflowClient{},
			ListFunc:       (&armdevhub.WorkflowClient{}).NewListPager,
			NewFunc:        armdevhub.NewWorkflowClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/workflows",
		},
	}

	for i := range tables {
		tables[i].Service = "armdevhub"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdevhub()...)
}
