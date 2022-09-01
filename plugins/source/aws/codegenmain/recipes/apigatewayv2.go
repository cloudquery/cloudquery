package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var APIGatewayv2Resources = combine(parentize(&Resource{
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
),
	parentize(&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:      &types.DomainName{},
		AWSService:     "Apigatewayv2",
		Template:       "resource_get",
		ItemsStruct:    &apigatewayv2.GetDomainNamesOutput{},
		ItemsCustomOptionsBlock: `
			// NOTE: Swapping OperationDeserializer until this is fixed: https://github.com/aws/aws-sdk-go-v2/issues/1282
			opts.APIOptions = append(opts.APIOptions, apigatewayv2fix.SwapGetDomainNamesOperationDeserializer)
`,
		Imports: []string{
			`apigatewayv2fix "github.com/cloudquery/cloudquery/plugins/source/aws/resources/forks/apigatewayv2"`,
		},
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolvers.ResolveDomainNameArn",
			},
		},
	},
		&Resource{
			AWSStruct:       &types.ApiMapping{},
			Template:        "resource_get",
			ItemsStruct:     &apigatewayv2.GetApiMappingsOutput{},
			ParentFieldName: "DomainName",
			ColumnOverrides: map[string]codegen.ColumnDefinition{
				"arn": {
					Type:     schema.TypeString,
					Resolver: "resolvers.ResolveApiMappingArn",
				},
			},
		},
	),
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:      &types.VpcLink{},
		AWSService:     "Apigatewayv2",
		Template:       "resource_get",
		ItemsStruct:    &apigatewayv2.GetVpcLinksOutput{},
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolvers.ResolveVPCLinkArn",
			},
		},
	},
)
