package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var APIGatewayv2Resources = parentize(&Resource{
	DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
	AWSStruct:      &types.Api{},
	AWSService:     "Apigatewayv2",
	AWSSubService:  "Apis",
	Template:       "resource_get",
	//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	ColumnOverrides: map[string]codegen.ColumnDefinition{
		"tags": {
			Type:        schema.TypeJSON,
			Description: "A collection of tags associated with the API.",
		},
	},
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
			ColumnOverrides: map[string]codegen.ColumnDefinition{
				"model_template": {
					Type:     schema.TypeString,
					Resolver: "resolvers.ResolveApiModelTemplate",
				},
			},
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
