// Code generated by codegen; DO NOT EDIT.

package amp

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Workspaces() *schema.Table {
	return &schema.Table{
		Name:                "aws_amp_workspaces",
		Description:         `https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription`,
		Resolver:            fetchAmpWorkspaces,
		PreResourceResolver: describeWorkspace,
		Multiplex:           client.ServiceAccountRegionMultiplexer("amp"),
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
				Name:     "alert_manager_definition",
				Type:     schema.TypeJSON,
				Resolver: describeAlertManagerDefinition,
			},
			{
				Name:     "logging_configuration",
				Type:     schema.TypeJSON,
				Resolver: describeLoggingConfiguration,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "workspace_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkspaceId"),
			},
			{
				Name:     "alias",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Alias"),
			},
			{
				Name:     "prometheus_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrometheusEndpoint"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			RuleGroupsNamespaces(),
		},
	}
}
