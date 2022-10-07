// Code generated by codegen; DO NOT EDIT.

package ecs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterContainerInstances() *schema.Table {
	return &schema.Table{
		Name:      "aws_ecs_cluster_container_instances",
		Resolver:  fetchEcsClusterContainerInstances,
		Multiplex: client.ServiceAccountRegionMultiplexer("ecs"),
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
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "agent_connected",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AgentConnected"),
			},
			{
				Name:     "agent_update_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AgentUpdateStatus"),
			},
			{
				Name:     "attachments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attachments"),
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
			{
				Name:     "capacity_provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CapacityProviderName"),
			},
			{
				Name:     "container_instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContainerInstanceArn"),
			},
			{
				Name:     "ec2_instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ec2InstanceId"),
			},
			{
				Name:     "health_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthStatus"),
			},
			{
				Name:     "pending_tasks_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingTasksCount"),
			},
			{
				Name:     "registered_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RegisteredAt"),
			},
			{
				Name:     "registered_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RegisteredResources"),
			},
			{
				Name:     "remaining_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RemainingResources"),
			},
			{
				Name:     "running_tasks_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RunningTasksCount"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusReason"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "version_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VersionInfo"),
			},
		},
	}
}
