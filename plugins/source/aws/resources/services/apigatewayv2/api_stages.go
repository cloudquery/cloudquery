// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiStages() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_api_stages",
		Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html",
		Resolver:    fetchApigatewayv2ApiStages,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Name:     "api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "api_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiStageArn(),
			},
			{
				Name:     "stage_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StageName"),
			},
			{
				Name:     "access_log_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessLogSettings"),
			},
			{
				Name:     "api_gateway_managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ApiGatewayManaged"),
			},
			{
				Name:     "auto_deploy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoDeploy"),
			},
			{
				Name:     "client_certificate_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientCertificateId"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "default_route_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultRouteSettings"),
			},
			{
				Name:     "deployment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentId"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "last_deployment_status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastDeploymentStatusMessage"),
			},
			{
				Name:     "last_updated_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedDate"),
			},
			{
				Name:     "route_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RouteSettings"),
			},
			{
				Name:     "stage_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StageVariables"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
