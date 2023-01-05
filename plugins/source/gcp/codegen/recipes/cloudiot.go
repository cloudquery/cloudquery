package recipes

import (
	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "device_registries",
			Struct:              &pb.DeviceRegistry{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Relations:           []string{"Devices()"},
			Description:         "https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries#DeviceRegistry",
		},
		{
			SubService:          "devices",
			Struct:              &pb.Device{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.DeviceRegistry).Name,`,
			Description:         "https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices#Device",
			ChildTable:          true,
			Relations:           []string{"DeviceConfigs(), DeviceStates()"},
		},
		{
			SubService:          "device_configs",
			Struct:              &pb.DeviceConfig{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "device_name"},
			ExtraColumns:        []codegen.ColumnDefinition{{Name: "device_name", Type: schema.TypeString, Resolver: `schema.ParentColumnResolver("name")`}},
			RequestStructFields: `Name: parent.Item.(*pb.Device).Name,`,
			Description:         "https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.configVersions#DeviceConfig",
			ChildTable:          true,
		},
		{
			SubService:          "device_states",
			Struct:              &pb.DeviceState{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "device_name"},
			ExtraColumns:        []codegen.ColumnDefinition{{Name: "device_name", Type: schema.TypeString, Resolver: `schema.ParentColumnResolver("name")`}},
			RequestStructFields: `Name: parent.Item.(*pb.Device).Name,`,
			Description:         "https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.states#DeviceState",
			ChildTable:          true,
		},
	}

	for _, resource := range resources {
		resource.Service = "cloudiot"
		resource.ServiceAPIOverride = "iot"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/iot/apiv1/iotpb"
		resource.MockImports = []string{"cloud.google.com/go/iot/apiv1"}
		resource.NewFunction = iot.NewDeviceManagerClient
		resource.RegisterServer = pb.RegisterDeviceManagerServer
		resource.SkipFetch = true
		resource.SkipMock = true
	}

	Resources = append(Resources, resources...)
}
