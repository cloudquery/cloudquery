// Code generated by codegen; DO NOT EDIT.
package packages

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"

func Armstreamanalytics() []*Table {
	tables := []*Table{
		{
			NewFunc: armstreamanalytics.NewStreamingJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armstreamanalytics())
}