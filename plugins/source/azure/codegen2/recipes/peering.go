// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"

func init() {
	tables := []Table{
		{
			Service:        "armpeering",
			Name:           "service_countries",
			Struct:         &armpeering.ServiceCountry{},
			ResponseStruct: &armpeering.ServiceCountriesClientListResponse{},
			Client:         &armpeering.ServiceCountriesClient{},
			ListFunc:       (&armpeering.ServiceCountriesClient{}).NewListPager,
			NewFunc:        armpeering.NewServiceCountriesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceCountries",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Peering)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armpeering",
			Name:           "service_locations",
			Struct:         &armpeering.ServiceLocation{},
			ResponseStruct: &armpeering.ServiceLocationsClientListResponse{},
			Client:         &armpeering.ServiceLocationsClient{},
			ListFunc:       (&armpeering.ServiceLocationsClient{}).NewListPager,
			NewFunc:        armpeering.NewServiceLocationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceLocations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Peering)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armpeering",
			Name:           "service_providers",
			Struct:         &armpeering.ServiceProvider{},
			ResponseStruct: &armpeering.ServiceProvidersClientListResponse{},
			Client:         &armpeering.ServiceProvidersClient{},
			ListFunc:       (&armpeering.ServiceProvidersClient{}).NewListPager,
			NewFunc:        armpeering.NewServiceProvidersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceProviders",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Peering)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
