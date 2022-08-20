package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"google.golang.org/api/compute/v1"
)

var computeResources = []Resource{
	{
		GCPSubService: "addresses",
		GCPStruct:     &compute.Address{},
	},
	{
		GCPSubService: "autoscalers",
		GCPStruct:     &compute.Autoscaler{},
	},
	{
		GCPSubService: "backend_services",
		GCPStruct:     &compute.BackendService{},
	},
	{
		GCPSubService: "disk_types",
		GCPStruct:     &compute.DiskType{},
	},
}

func ComputeResources() []Resource {
	resources := computeResources
	// add all shared properties
	for i := range resources {
		resources[i].GCPService = "compute"
		resources[i].Imports = []string{"google.golang.org/api/compute/v1"}
		resources[i].Template = "resource1"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].GCPStructName = reflect.TypeOf(resources[i].GCPStruct).Elem().Name()
	}
	return resources
}
