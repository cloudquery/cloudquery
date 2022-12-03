package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

func LogicResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "workflows",
			Struct: &armlogic.Workflow{},
			ResponseStruct: &armlogic.WorkflowsClientListBySubscriptionResponse{},
			Client: &armlogic.WorkflowsClient{},
			ListFunc: (&armlogic.WorkflowsClient{}).NewListBySubscriptionPager,
			NewFunc: armlogic.NewWorkflowsClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "logic/armlogic"
		r.Service = "armlogic"
		r.Template = "list"
	}

	return resources
}