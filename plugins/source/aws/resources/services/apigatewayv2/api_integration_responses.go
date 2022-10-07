// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiIntegrationResponses() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigatewayv2_api_integration_responses",
		Resolver:  fetchApigatewayv2ApiIntegrationResponses,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigateway"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "api_integration_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "integration_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("integration_id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiIntegrationResponseArn(),
			},
			{
				Name:     "integration_response_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IntegrationResponseKey"),
			},
			{
				Name:     "content_handling_strategy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContentHandlingStrategy"),
			},
			{
				Name:     "integration_response_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IntegrationResponseId"),
			},
			{
				Name:     "response_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponseParameters"),
			},
			{
				Name:     "response_templates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponseTemplates"),
			},
			{
				Name:     "template_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TemplateSelectionExpression"),
			},
		},
	}
}
