// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"

func Armstreamanalytics() []Table {
	tables := []Table{
		{
			Name:           "streaming_jobs",
			Struct:         &armstreamanalytics.StreamingJob{},
			ResponseStruct: &armstreamanalytics.StreamingJobsClientListResponse{},
			Client:         &armstreamanalytics.StreamingJobsClient{},
			ListFunc:       (&armstreamanalytics.StreamingJobsClient{}).NewListPager,
			NewFunc:        armstreamanalytics.NewStreamingJobsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_StreamAnalytics)`,
		},
	}

	for i := range tables {
		tables[i].Service = "armstreamanalytics"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armstreamanalytics()...)
}
