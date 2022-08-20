package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"google.golang.org/api/cloudfunctions/v1"
)

var CloudFunctionsResources = []Resource{
	{
		TableFunctionName: "CloudFunctions",
		PackageName:       "cloudfunctions",
		FileName:          "functions.go",
		ListFunction:      "ListFunction",
		Struct:            cloudfunctions.CloudFunction{},
		DefaultColumns:    []codegen.ColumnDefinition{ProjectIdColumn},
		Template:          "resource.go.tpl",
	},
}
