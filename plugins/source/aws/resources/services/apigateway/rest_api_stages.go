// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RestApiStages() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_rest_api_stages",
		Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html",
		Resolver:    fetchApigatewayRestApiStages,
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
				Name:     "rest_api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIStageArn,
			},
			{
				Name:     "access_log_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessLogSettings"),
			},
			{
				Name:     "cache_cluster_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CacheClusterEnabled"),
			},
			{
				Name:     "cache_cluster_size",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheClusterSize"),
			},
			{
				Name:     "cache_cluster_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheClusterStatus"),
			},
			{
				Name:     "canary_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CanarySettings"),
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
				Name:     "documentation_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocumentationVersion"),
			},
			{
				Name:     "last_updated_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedDate"),
			},
			{
				Name:     "method_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MethodSettings"),
			},
			{
				Name:     "stage_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StageName"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "tracing_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TracingEnabled"),
			},
			{
				Name:     "variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Variables"),
			},
			{
				Name:     "web_acl_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebAclArn"),
			},
		},
	}
}
