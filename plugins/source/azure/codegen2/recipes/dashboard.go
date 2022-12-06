// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"

func Armdashboard() []Table {
	tables := []Table{
		{
			Name:           "managed_grafana",
			Struct:         &armdashboard.ManagedGrafana{},
			ResponseStruct: &armdashboard.GrafanaClientListResponse{},
			Client:         &armdashboard.GrafanaClient{},
			ListFunc:       (&armdashboard.GrafanaClient{}).NewListPager,
			NewFunc:        armdashboard.NewGrafanaClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Dashboard/grafana",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Dashboard")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armdashboard"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdashboard()...)
}
