// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"

func Armcustomerlockbox() []Table {
	tables := []Table{
		{
			Name:           "lockbox_request_response",
			Struct:         &armcustomerlockbox.LockboxRequestResponse{},
			ResponseStruct: &armcustomerlockbox.RequestsClientListResponse{},
			Client:         &armcustomerlockbox.RequestsClient{},
			ListFunc:       (&armcustomerlockbox.RequestsClient{}).NewListPager,
			NewFunc:        armcustomerlockbox.NewRequestsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerLockbox/requests",
		},
	}

	for i := range tables {
		tables[i].Service = "armcustomerlockbox"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcustomerlockbox()...)
}
