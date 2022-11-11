// Code generated by codegen; DO NOT EDIT.

package ecs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TaskDefinitions() *schema.Table {
	return &schema.Table{
		Name:                "aws_ecs_task_definitions",
		Description:         `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html`,
		Resolver:            fetchEcsTaskDefinitions,
		PreResourceResolver: getTaskDefinition,
		Multiplex:           client.ServiceAccountRegionMultiplexer("ecs"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskDefinitionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEcsTaskDefinitionTags,
			},
			{
				Name:     "compatibilities",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Compatibilities"),
			},
			{
				Name:     "container_definitions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ContainerDefinitions"),
			},
			{
				Name:     "cpu",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Cpu"),
			},
			{
				Name:     "deregistered_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeregisteredAt"),
			},
			{
				Name:     "ephemeral_storage",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EphemeralStorage"),
			},
			{
				Name:     "execution_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExecutionRoleArn"),
			},
			{
				Name:     "family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Family"),
			},
			{
				Name:     "inference_accelerators",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InferenceAccelerators"),
			},
			{
				Name:     "ipc_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpcMode"),
			},
			{
				Name:     "memory",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Memory"),
			},
			{
				Name:     "network_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkMode"),
			},
			{
				Name:     "pid_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PidMode"),
			},
			{
				Name:     "placement_constraints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PlacementConstraints"),
			},
			{
				Name:     "proxy_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProxyConfiguration"),
			},
			{
				Name:     "registered_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RegisteredAt"),
			},
			{
				Name:     "registered_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegisteredBy"),
			},
			{
				Name:     "requires_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RequiresAttributes"),
			},
			{
				Name:     "requires_compatibilities",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RequiresCompatibilities"),
			},
			{
				Name:     "revision",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Revision"),
			},
			{
				Name:     "runtime_platform",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RuntimePlatform"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "task_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskRoleArn"),
			},
			{
				Name:     "volumes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Volumes"),
			},
		},
	}
}
