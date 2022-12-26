package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"

func init() {
	tables := []Table{
		{
			Service:        "armlogic",
			Name:           "workflows",
			Struct:         &armlogic.Workflow{},
			ResponseStruct: &armlogic.WorkflowsClientListBySubscriptionResponse{},
			Client:         &armlogic.WorkflowsClient{},
			ListFunc:       (&armlogic.WorkflowsClient{}).NewListBySubscriptionPager,
			NewFunc:        armlogic.NewWorkflowsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_logic)`,
		},
	}
	Tables = append(Tables, tables...)
}
