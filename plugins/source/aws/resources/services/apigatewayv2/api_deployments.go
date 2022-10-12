// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiDeployments() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_api_deployments",
		Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Deployment.html",
		Resolver:    fetchApigatewayv2ApiDeployments,
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
				Resolver: resolveApiDeploymentArn(),
			},
			{
				Name:     "auto_deployed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoDeployed"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "deployment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentId"),
			},
			{
				Name:     "deployment_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentStatus"),
			},
			{
				Name:     "deployment_status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentStatusMessage"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
		},
	}
}
