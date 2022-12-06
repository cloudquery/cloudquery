// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"

func Armredhatopenshift() []Table {
	tables := []Table{
		{
			Name:           "open_shift_cluster",
			Struct:         &armredhatopenshift.OpenShiftCluster{},
			ResponseStruct: &armredhatopenshift.OpenShiftClustersClientListResponse{},
			Client:         &armredhatopenshift.OpenShiftClustersClient{},
			ListFunc:       (&armredhatopenshift.OpenShiftClustersClient{}).NewListPager,
			NewFunc:        armredhatopenshift.NewOpenShiftClustersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.RedHatOpenShift/openShiftClusters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.RedHatOpenShift")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armredhatopenshift"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armredhatopenshift()...)
}
