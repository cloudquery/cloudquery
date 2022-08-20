package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"google.golang.org/api/run/v1"
)

var CloudRunResources = []Resource{

	{
		TableFunctionName: "Services",
		PackageName:       "cloudrun",
		FileName:          "services.go",
		ListFunction:      "CloudRun.Projects.Locations.Services.List",
		Struct:            run.Service{},
		DefaultColumns:    []codegen.ColumnDefinition{ProjectIdColumn},
		Template:          "resource.go.tpl",
	},
}
