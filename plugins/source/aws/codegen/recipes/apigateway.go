package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func APIGatewayesources() []*Resource {
	resources := []*Resource{
		{
			SubService: "api_keys",
			Struct:     &types.ApiKey{},
			SkipFields: []string{"Arn"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("apigateway")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiKeyArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "client_certificates",
			Struct:     &types.ClientCertificate{},
			SkipFields: []string{"Arn"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("apigateway")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveClientCertificateArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "vpc_links",
			Struct:     &types.ClientCertificate{},
			SkipFields: []string{"Arn"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("apigateway")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcLinkArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "apigateway"
	}
	return resources
}
