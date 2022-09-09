package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type TaskDefinitionWrapper struct {
	*types.TaskDefinition
	Tags []types.Tag
}

func EcsTaskDefinitions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_task_definitions",
		Description: "The details of a task definition which describes the container and volume definitions of an Amazon Elastic Container Service task",
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
			return client.ListAndDetailResolver(ctx, meta, res, listEcsTaskDefinitions, ecsTaskDefinitionDetail)
		},
		Multiplex:     client.ServiceAccountRegionMultiplexer("ecs"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "The metadata that you apply to the service to help you categorize and organize them",
				Type:        schema.TypeJSON,
				Resolver:    resolveEcsTaskDefinitionTags,
			},
			{
				Name:        "compatibilities",
				Description: "The task launch types the task definition validated against during task definition registration",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "cpu",
				Description: "The number of cpu units used by the task",
				Type:        schema.TypeString,
			},
			{
				Name:        "deregistered_at",
				Description: "The Unix timestamp for when the task definition was deregistered.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "ephemeral_storage_size",
				Description: "The total amount, in GiB, of ephemeral storage to set for the task.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EphemeralStorage.SizeInGiB"),
			},
			{
				Name:        "execution_role_arn",
				Description: "The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf",
				Type:        schema.TypeString,
			},
			{
				Name:        "family",
				Description: "The name of a family that this task definition is registered to",
				Type:        schema.TypeString,
			},
			{
				Name:        "inference_accelerators",
				Description: "The Elastic Inference accelerator associated with the task.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEcsTaskDefinitionsInferenceAccelerators,
			},
			{
				Name:        "ipc_mode",
				Description: "The IPC resource namespace to use for the containers in the task",
				Type:        schema.TypeString,
			},
			{
				Name:        "memory",
				Description: "The amount (in MiB) of memory used by the task",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_mode",
				Description: "The Docker networking mode to use for the containers in the task",
				Type:        schema.TypeString,
			},
			{
				Name:        "pid_mode",
				Description: "The process namespace to use for the containers in the task",
				Type:        schema.TypeString,
			},
			{
				Name:        "placement_constraints",
				Description: "An array of placement constraint objects to use for tasks",
				Type:        schema.TypeJSON,
				Resolver:    resolveEcsTaskDefinitionsPlacementConstraints,
			},
			{
				Name:        "proxy_configuration_container_name",
				Description: "The name of the container that will serve as the App Mesh proxy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProxyConfiguration.ContainerName"),
			},
			{
				Name:        "proxy_configuration_properties",
				Description: "The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified as key-value pairs.  * IgnoredUID - (Required) The user ID (UID) of the proxy container as defined by the user parameter in a container definition",
				Type:        schema.TypeJSON,
				Resolver:    resolveEcsTaskDefinitionsProxyConfigurationProperties,
			},
			{
				Name:        "proxy_configuration_type",
				Description: "The proxy type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProxyConfiguration.Type"),
			},
			{
				Name:        "registered_at",
				Description: "The Unix timestamp for when the task definition was registered.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "registered_by",
				Description: "The principal that registered the task definition.",
				Type:        schema.TypeString,
			},
			{
				Name:        "requires_attributes",
				Description: "The container instance attributes required by your task",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("RequiresAttributes"),
			},
			{
				Name:        "requires_compatibilities",
				Description: "The task launch types the task definition was validated against",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "revision",
				Description: "The revision of the task in a particular family",
				Type:        schema.TypeInt,
			},
			{
				Name:        "runtime_platform_cpu_architecture",
				Description: "The CPU architecture.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuntimePlatform.CpuArchitecture"),
			},
			{
				Name:        "runtime_platform_os_family",
				Description: "The operating system.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuntimePlatform.OperatingSystemFamily"),
			},
			{
				Name:        "status",
				Description: "The status of the task definition.",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The full Amazon Resource Name (ARN) of the task definition.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("TaskDefinitionArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "task_role_arn",
				Description: "The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants containers in the task permission to call AWS APIs on your behalf",
				Type:        schema.TypeString,
			},
			{
				Name:        "container_definitions",
				Description: "Container definitions are used in task definitions to describe the different containers that are launched as part of a task.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ContainerDefinitions"),
			},
			{
				Name:        "volumes",
				Description: "A data volume used in a task definition",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Volumes"),
			},
		},
	}
}
func ecsTaskDefinitionDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, detail interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().ECS
	taskArn := detail.(string)
	describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	})
	if err != nil {
		errorChan <- err
		return
	}
	if describeTaskDefinitionOutput.TaskDefinition == nil {
		return
	}
	resultsChan <- TaskDefinitionWrapper{
		TaskDefinition: describeTaskDefinitionOutput.TaskDefinition,
		Tags:           describeTaskDefinitionOutput.Tags,
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func listEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}) error {
	var config ecs.ListTaskDefinitionsInput
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	for {
		listClustersOutput, err := svc.ListTaskDefinitions(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		for _, taskDefinitionArn := range listClustersOutput.TaskDefinitionArns {
			res <- taskDefinitionArn
		}
		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}
func resolveEcsTaskDefinitionsInferenceAccelerators(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]interface{}{}
	for _, a := range r.InferenceAccelerators {
		if a.DeviceName == nil {
			continue
		}
		j[*a.DeviceName] = aws.ToString(a.DeviceType)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionsPlacementConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]interface{}{}
	for _, p := range r.PlacementConstraints {
		if p.Expression == nil {
			continue
		}
		j[*p.Expression] = p.Type
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionsProxyConfigurationProperties(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]interface{}{}
	if r.ProxyConfiguration == nil {
		return nil
	}
	for _, p := range r.ProxyConfiguration.Properties {
		if p.Name == nil {
			continue
		}
		j[*p.Name] = aws.ToString(p.Value)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]string{}
	for _, a := range r.Tags {
		if a.Key == nil {
			continue
		}
		j[*a.Key] = aws.ToString(a.Value)
	}
	return resource.Set(c.Name, j)
}
