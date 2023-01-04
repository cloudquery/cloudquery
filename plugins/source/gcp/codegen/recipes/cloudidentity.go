package recipes

import (
	"google.golang.org/api/cloudidentity/v1"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "devices",
			Struct:      &cloudidentity.GoogleAppsCloudidentityDevicesV1Device{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/identity/docs/reference/rest/v1/devices#Device",
		},
		{
			SubService:  "device_users",
			Struct:      &cloudidentity.GoogleAppsCloudidentityDevicesV1DeviceUser{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers#DeviceUser",
		},
		{
			SubService:  "client_states",
			Struct:      &cloudidentity.GoogleAppsCloudidentityDevicesV1ClientState{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers.clientStates#ClientState",
		},
	}

	for _, resource := range resources {
		resource.Service = "cloudidentity"
		resource.Template = "newapi_list"
		resource.SkipFetch = true
		resource.MockTemplate = "resource_list_mock"
		resource.MockImports = []string{"google.golang.org/api/cloudidentity/v1"}
		resource.OutputField = Caser.ToPascal(resource.SubService)
	}

	Resources = append(Resources, resources...)
}
