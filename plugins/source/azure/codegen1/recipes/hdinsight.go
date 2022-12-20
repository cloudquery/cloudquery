// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"

func Armhdinsight() []*Table {
	tables := []*Table{
		{
			NewFunc:        armhdinsight.NewClustersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/clusters",
			Namespace:      "microsoft.hdinsight",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_hdinsight)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ClustersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armhdinsight())
}
