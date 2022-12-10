// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"

func Armhealthcareapis() []Table {
	tables := []Table{
		{
			Service:        "armhealthcareapis",
			Name:           "services",
			Struct:         &armhealthcareapis.ServicesDescription{},
			ResponseStruct: &armhealthcareapis.ServicesClientListResponse{},
			Client:         &armhealthcareapis.ServicesClient{},
			ListFunc:       (&armhealthcareapis.ServicesClient{}).NewListPager,
			NewFunc:        armhealthcareapis.NewServicesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/services",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_HealthcareApis)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armhealthcareapis()...)
}
