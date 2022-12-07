// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"

func Armhealthcareapis() []Table {
	tables := []Table{
		{
			Name:           "services",
			Struct:         &armhealthcareapis.ServicesDescription{},
			ResponseStruct: &armhealthcareapis.ServicesClientListResponse{},
			Client:         &armhealthcareapis.ServicesClient{},
			ListFunc:       (&armhealthcareapis.ServicesClient{}).NewListPager,
			NewFunc:        armhealthcareapis.NewServicesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/services",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.HealthcareApis")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armhealthcareapis"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armhealthcareapis()...)
}
