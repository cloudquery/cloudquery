package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var APIGatewayv2Resources = parentize(&Resource{
	DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
	//Table:              nil, // will be "generated" at "runtime"
	AWSStruct:     &types.Api{},
	AWSService:    "apigatewayv2",
	AWSSubService: "apis",
	ItemName:      "Api",
	//DescribeFieldName:  "CertificateArn",
	Template:       "resource_get",
	Imports:        nil,
	MockImports:    nil,
	MockListStruct: "",
	SkipFields:     nil,
	//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	ColumnOverrides: map[string]codegen.ColumnDefinition{
		"tags": {
			Type:        schema.TypeJSON,
			Description: "A collection of tags associated with the API.",
		},
	},
	SkipTypesImport: true,
},
	&Resource{
		AWSStruct:       &types.Authorizer{},
		AWSSubService:   "authorizers",
		ItemName:        "Authorizer",
		Template:        "resource_get",
		ParentFieldName: "ApiId",
	},
)
