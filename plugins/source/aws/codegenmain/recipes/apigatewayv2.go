package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var APIGatewayv2Resources = parentize(&Resource{
	DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
	AWSStruct:      &types.Api{},
	AWSService:     "Apigatewayv2",
	Template:       "resource_get",
	ItemsStruct:    &apigatewayv2.GetApisOutput{},
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
			Template:        "resource_get",
			ParentFieldName: "ApiId",
			ItemsStruct:     &apigatewayv2.GetAuthorizersOutput{},
		},
		&Resource{
			AWSStruct:       &types.Deployment{},
			Template:        "resource_get",
			ParentFieldName: "ApiId",
			ItemsStruct:     &apigatewayv2.GetDeploymentsOutput{},
		},
		parentize(
			&Resource{
				AWSStruct:       &types.Integration{},
				Template:        "resource_get",
				ParentFieldName: "ApiId",
				ItemsStruct:     &apigatewayv2.GetIntegrationsOutput{},
			},
			&Resource{
				AWSStruct:       &types.IntegrationResponse{},
				Template:        "resource_get",
				ParentFieldName: "IntegrationId",
				ItemsStruct:     &apigatewayv2.GetIntegrationResponsesOutput{},
			},
		),
		&Resource{
			AWSStruct:       &types.Model{},
			Template:        "resource_get",
			ParentFieldName: "ApiId",
			ItemsStruct:     &apigatewayv2.GetModelsOutput{},
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
				Template:        "resource_get",
				ParentFieldName: "ApiId",
				ItemsStruct:     &apigatewayv2.GetRoutesOutput{},
			},
			&Resource{
				AWSStruct:       &types.RouteResponse{},
				Template:        "resource_get",
				ParentFieldName: "RouteId",
				ItemsStruct:     &apigatewayv2.GetRouteResponsesOutput{},
			},
		),
		&Resource{
			AWSStruct:       &types.Stage{},
			Template:        "resource_get",
			ParentFieldName: "ApiId",
			ItemsStruct:     &apigatewayv2.GetStagesOutput{},
		},
	)...,
)
