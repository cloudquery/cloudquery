// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"

func init() {
	tables := []Table{
		{
			Service:        "armeventgrid",
			Name:           "topic_types",
			Struct:         &armeventgrid.TopicTypeInfo{},
			ResponseStruct: &armeventgrid.TopicTypesClientListResponse{},
			Client:         &armeventgrid.TopicTypesClient{},
			ListFunc:       (&armeventgrid.TopicTypesClient{}).NewListPager,
			NewFunc:        armeventgrid.NewTopicTypesClient,
			URL:            "/providers/Microsoft.EventGrid/topicTypes",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_EventGrid)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
