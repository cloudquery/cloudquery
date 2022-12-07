// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"

func Armmaintenance() []Table {
	tables := []Table{
		{
			Name:           "configurations",
			Struct:         &armmaintenance.Configuration{},
			ResponseStruct: &armmaintenance.ConfigurationsClientListResponse{},
			Client:         &armmaintenance.ConfigurationsClient{},
			ListFunc:       (&armmaintenance.ConfigurationsClient{}).NewListPager,
			NewFunc:        armmaintenance.NewConfigurationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/maintenanceConfigurations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Maintenance")`,
		},
		{
			Name:           "public_maintenance_configurations",
			Struct:         &armmaintenance.Configuration{},
			ResponseStruct: &armmaintenance.PublicMaintenanceConfigurationsClientListResponse{},
			Client:         &armmaintenance.PublicMaintenanceConfigurationsClient{},
			ListFunc:       (&armmaintenance.PublicMaintenanceConfigurationsClient{}).NewListPager,
			NewFunc:        armmaintenance.NewPublicMaintenanceConfigurationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Maintenance")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armmaintenance"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmaintenance()...)
}
