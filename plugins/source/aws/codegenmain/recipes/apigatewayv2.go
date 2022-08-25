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
	AWSService:    "Apigatewayv2",
	AWSSubService: "Apis",
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
	combine(
		&Resource{
			AWSStruct:       &types.Authorizer{},
			AWSSubService:   "Authorizers",
			Template:        "resource_get",
			ParentFieldName: "ApiId",
		},
		&Resource{
			AWSStruct:       &types.Deployment{},
			AWSSubService:   "Deployments",
			Template:        "resource_get",
			ParentFieldName: "ApiId",
		},
		parentize(
			&Resource{
				AWSStruct:       &types.Integration{},
				AWSSubService:   "Integrations",
				Template:        "resource_get",
				ParentFieldName: "ApiId",
			},
			&Resource{
				AWSStruct:       &types.IntegrationResponse{},
				AWSSubService:   "IntegrationResponses",
				Template:        "resource_get",
				ParentFieldName: "IntegrationId",
			},
		),
		&Resource{
			AWSStruct:       &types.Model{},
			AWSSubService:   "Models",
			Template:        "resource_get",
			ParentFieldName: "ApiId",
			/*
				TODO this should be a resolver
					&Resource{
						AWSStruct:       aws.String(""),   // *string
						AWSSubService:   "modeltemplates",
						ItemName:        "ModelTemplate",
						Template:        "resource_get",
						ParentFieldName: "ModelId",
					},
			*/
		},
		parentize(
			&Resource{
				AWSStruct:       &types.Route{},
				AWSSubService:   "Routes",
				Template:        "resource_get",
				ParentFieldName: "ApiId",
			},
			&Resource{
				AWSStruct:       &types.RouteResponse{},
				AWSSubService:   "RouteResponses",
				Template:        "resource_get",
				ParentFieldName: "RouteId",
			},
		),
		&Resource{
			AWSStruct:       &types.Stage{},
			AWSSubService:   "Stages",
			Template:        "resource_get",
			ParentFieldName: "ApiId",
		},
	)...,
)
