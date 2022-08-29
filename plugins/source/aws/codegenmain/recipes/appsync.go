package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var AppsyncResources = []*Resource{
	{
		DefaultColumns:    []codegen.ColumnDefinition{AccountIdColumn, RegionColumn, NamespaceColumn},
		AWSStruct:         &types.GraphqlApi{},
		AWSService:        "AppSync",
		AWSSubService:     "GraphqlApis",
		Template:          "resource_get",
		Verb:              "List",
		ResponseItemsName: "GraphqlApis",
		Imports:           nil,
		MockImports:       nil,
		MockListStruct:    "",
		SkipFields:        nil,
		SkipTypesImport:   true,
		TrimPrefix:        "api_",
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
}
