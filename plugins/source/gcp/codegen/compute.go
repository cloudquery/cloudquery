package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"google.golang.org/api/compute/v1"
)

var ComputeResources = []Resource{
	{
		TableFunctionName: "ComputeAddresses",
		PackageName:       "compute",
		FileName:          "addresses.go",
		ListFunction:      "Compute.Addresses.AggregatedList",
		Struct:            compute.Address{},
		DefaultColumns:    []codegen.ColumnDefinition{ProjectIdColumn},
		Template:          "resource1.go.tpl",
		Service:           "compute",
		StructField:       "Addresses",
		StructName:        "Address",
		Imports:           []string{"google.golang.org/api/compute/v1"},
	},
	{
		TableFunctionName: "ComputeAutoscalers",
		PackageName:       "compute",
		FileName:          "autoscalers.go",
		ListFunction:      "Compute.Autoscalers.AggregatedList",
		Struct:            compute.Autoscaler{},
		DefaultColumns:    []codegen.ColumnDefinition{ProjectIdColumn},
		Template:          "resource1.go.tpl",
		Service:           "compute",
		StructField:       "Autoscalers",
		StructName:        "Autoscaler",
		Imports:           []string{"google.golang.org/api/compute/v1"},
	},
	{
		TableFunctionName: "BackendServices",
		PackageName:       "compute",
		FileName:          "backend_services.go",
		ListFunction:      "Compute.BackendServices.AggregatedList",
		Struct:            compute.BackendService{},
		DefaultColumns:    []codegen.ColumnDefinition{ProjectIdColumn},
		Template:          "resource1.go.tpl",
		Service:           "compute",
		StructField:       "BackendServices",
		StructName:        "BackendService",
		Imports:           []string{"google.golang.org/api/compute/v1"},
	},
	{
		TableFunctionName: "DiskTypes",
		PackageName:       "compute",
		FileName:          "disk_types.go",
		ListFunction:      "Compute.DiskTypes.AggregatedList",
		Struct:            compute.DiskType{},
		DefaultColumns:    []codegen.ColumnDefinition{ProjectIdColumn},
		Template:          "resource1.go.tpl",
		Service:           "compute",
		StructField:       "DiskTypes",
		StructName:        "DiskType",
		Imports:           []string{"google.golang.org/api/compute/v1"},
	},
}
