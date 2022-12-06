// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"

func Armdatadog() []Table {
	tables := []Table{
		{
      Name: "agreement_resource",
      Struct: &armdatadog.AgreementResource{},
      ResponseStruct: &armdatadog.MarketplaceAgreementsClientListResponse{},
      Client: &armdatadog.MarketplaceAgreementsClient{},
      ListFunc: (&armdatadog.MarketplaceAgreementsClient{}).NewListPager,
			NewFunc: armdatadog.NewMarketplaceAgreementsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/agreements",
		},
		{
      Name: "monitor_resource",
      Struct: &armdatadog.MonitorResource{},
      ResponseStruct: &armdatadog.MonitorsClientListResponse{},
      Client: &armdatadog.MonitorsClient{},
      ListFunc: (&armdatadog.MonitorsClient{}).NewListPager,
			NewFunc: armdatadog.NewMonitorsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/monitors",
		},
	}

	for i := range tables {
		tables[i].Service = "armdatadog"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdatadog()...)
}