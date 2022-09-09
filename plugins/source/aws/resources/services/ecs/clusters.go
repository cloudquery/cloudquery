package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_clusters",
		Description: "A regional grouping of one or more container instances where you can run task requests",
		Resolver:    fetchEcsClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
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
				Name:        "active_services_count",
				Description: "The number of services that are running on the cluster in an ACTIVE state",
				Type:        schema.TypeInt,
			},
			{
				Name:          "attachments_status",
				Description:   "The status of the capacity providers associated with the cluster",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "capacity_providers",
				Description: "The capacity providers associated with the cluster.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) that identifies the cluster",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "A user-generated string that you use to identify your cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterName"),
			},
			{
				Name:          "configuration",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Configuration"),
				IgnoreInTests: true,
			},
			{
				Name:        "execute_config_logging",
				Description: "The log setting to use for redirecting logs for your execute command results. The following log settings are available.  * NONE: The execute command session is not logged.  * DEFAULT: The awslogs configuration in the task definition is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.Logging"),
			},
			{
				Name:        "default_capacity_provider_strategy",
				Description: "The default capacity provider strategy for the cluster",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DefaultCapacityProviderStrategy"),
			},
			{
				Name:        "pending_tasks_count",
				Description: "The number of tasks in the cluster that are in the PENDING state.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "registered_container_instances_count",
				Description: "The number of container instances registered into the cluster",
				Type:        schema.TypeInt,
			},
			{
				Name:        "running_tasks_count",
				Description: "The number of tasks in the cluster that are in the RUNNING state.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "settings",
				Description: "The settings for the cluster",
				Type:        schema.TypeJSON,
				Resolver:    resolveClustersSettings,
			},
			{
				Name:        "statistics",
				Description: "Additional information about your clusters that are separated by launch type. They include the following:  * runningEC2TasksCount  * RunningFargateTasksCount  * pendingEC2TasksCount  * pendingFargateTasksCount  * activeEC2ServiceCount  * activeFargateServiceCount  * drainingEC2ServiceCount  * drainingFargateServiceCount",
				Type:        schema.TypeJSON,
				Resolver:    resolveClustersStatistics,
			},
			{
				Name:        "status",
				Description: "The status of the cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The metadata that you apply to the cluster to help you categorize and organize them",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "attachments",
				Description: "An object representing a container instance or task attachment.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Attachments"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ecs_cluster_tasks",
				Description:   "Details on a task in a cluster.",
				Resolver:      fetchEcsClusterTasks,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_ecs_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "attributes",
						Description: "The attributes of the task",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Attributes"),
					},
					{
						Name:        "availability_zone",
						Description: "The Availability Zone for the task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "capacity_provider_name",
						Description: "The capacity provider that's associated with the task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cluster_arn",
						Description: "The ARN of the cluster that hosts the task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "connectivity",
						Description: "The connectivity status of a task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "connectivity_at",
						Description: "The Unix timestamp for the time when the task last went into CONNECTED status.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "container_instance_arn",
						Description: "The ARN of the container instances that host the task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cpu",
						Description: "The number of CPU units used by the task as expressed in a task definition",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The Unix timestamp for the time when the task was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "desired_status",
						Description: "The desired status of the task",
						Type:        schema.TypeString,
					},
					{
						Name:        "enable_execute_command",
						Description: "Determines whether execute command functionality is enabled for this task",
						Type:        schema.TypeBool,
					},
					{
						Name:     "ephemeral_storage",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("EphemeralStorage"),
					},
					{
						Name:        "execution_stopped_at",
						Description: "The Unix timestamp for the time when the task execution stopped.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "group",
						Description: "The name of the task group that's associated with the task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "health_status",
						Description: "The health status for the task",
						Type:        schema.TypeString,
					},
					{
						Name:        "inference_accelerators",
						Description: "The Elastic Inference accelerator that's associated with the task.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("InferenceAccelerators"),
					},
					{
						Name:        "last_status",
						Description: "The last known status for the task",
						Type:        schema.TypeString,
					},
					{
						Name:        "launch_type",
						Description: "The infrastructure where your task runs on",
						Type:        schema.TypeString,
					},
					{
						Name:        "memory",
						Description: "The amount of memory (in MiB) that the task uses as expressed in a task definition",
						Type:        schema.TypeString,
					},
					{
						Name:        "overrides",
						Description: "One or more container overrides.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Overrides"),
					},
					{
						Name:        "platform_family",
						Description: "The operating system that your tasks are running on",
						Type:        schema.TypeString,
					},
					{
						Name:        "platform_version",
						Description: "The platform version where your task runs on",
						Type:        schema.TypeString,
					},
					{
						Name:        "pull_started_at",
						Description: "The Unix timestamp for the time when the container image pull began.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "pull_stopped_at",
						Description: "The Unix timestamp for the time when the container image pull completed.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "started_at",
						Description: "The Unix timestamp for the time when the task started",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "started_by",
						Description: "The tag specified when a task is started",
						Type:        schema.TypeString,
					},
					{
						Name:        "stop_code",
						Description: "The stop code indicating why a task was stopped",
						Type:        schema.TypeString,
					},
					{
						Name:        "stopped_at",
						Description: "The Unix timestamp for the time when the task was stopped",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "stopped_reason",
						Description: "The reason that the task was stopped.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stopping_at",
						Description: "The Unix timestamp for the time when the task stops",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "tags",
						Description: "The metadata that you apply to the task to help you categorize and organize the task",
						Type:        schema.TypeJSON,
						Resolver:    client.ResolveTags,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the task.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TaskArn"),
					},
					{
						Name:        "task_definition_arn",
						Description: "The ARN of the task definition that creates the task.",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "The version counter for the task",
						Type:        schema.TypeInt,
					},
					{
						Name:        "attachments",
						Description: "An object representing a container instance or task attachment.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Attachments"),
					},
					{
						Name:        "containers",
						Description: "A Docker container that's part of a task.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Containers"),
					},
				},
			},
			{
				Name:        "aws_ecs_cluster_services",
				Description: "Details on a service within a cluster",
				Resolver:    fetchEcsClusterServices,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_ecs_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "capacity_provider_strategy",
						Description:   "The capacity provider strategy the service uses",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("CapacityProviderStrategy"),
						IgnoreInTests: true,
					},
					{
						Name:        "cluster_arn",
						Description: "The Amazon Resource Name (ARN) of the cluster that hosts the service.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The Unix timestamp for the time when the service was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "created_by",
						Description: "The principal that created the service.",
						Type:        schema.TypeString,
					},
					{
						Name:     "deployment_configuration",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("DeploymentConfiguration"),
					},
					{
						Name:     "deployment_controller",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("DeploymentController"),
					},
					{
						Name:        "desired_count",
						Description: "The desired number of instantiations of the task definition to keep running on the service",
						Type:        schema.TypeInt,
					},
					{
						Name:        "enable_ecs_managed_tags",
						Description: "Determines whether to use Amazon ECS managed tags for the tasks in the service. For more information, see Tagging Your Amazon ECS Resources (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the Amazon Elastic Container Service Developer Guide.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EnableECSManagedTags"),
					},
					{
						Name:        "enable_execute_command",
						Description: "Determines whether the execute command functionality is enabled for the service. If true, the execute command functionality is enabled for all containers in tasks as part of the service.",
						Type:        schema.TypeBool,
					},
					{
						Name:          "health_check_grace_period_seconds",
						Description:   "The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.",
						Type:          schema.TypeInt,
						IgnoreInTests: true,
					},
					{
						Name:        "launch_type",
						Description: "The launch type the service is using",
						Type:        schema.TypeString,
					},
					{
						Name:     "network_configuration",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("NetworkConfiguration"),
					},
					{
						Name:        "pending_count",
						Description: "The number of tasks in the cluster that are in the PENDING state.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "placement_constraints",
						Description: "The placement constraints for the tasks in the service.",
						Type:        schema.TypeJSON,
						Resolver:    resolveClusterServicesPlacementConstraints,
					},
					{
						Name:        "placement_strategy",
						Description: "The placement strategy that determines how tasks for the service are placed.",
						Type:        schema.TypeJSON,
						Resolver:    resolveClusterServicesPlacementStrategy,
					},
					{
						Name:          "platform_family",
						Description:   "The operating system that your tasks in the service run on",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "platform_version",
						Description:   "The platform version to run your service on",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "propagate_tags",
						Description: "Determines whether to propagate the tags from the task definition or the service to the task",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_arn",
						Description: "The ARN of the IAM role that's associated with the service",
						Type:        schema.TypeString,
					},
					{
						Name:        "running_count",
						Description: "The number of tasks in the cluster that are in the RUNNING state.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "scheduling_strategy",
						Description: "The scheduling strategy to use for the service",
						Type:        schema.TypeString,
					},
					{
						Name:        "arn",
						Description: "The ARN that identifies the service",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServiceArn"),
					},
					{
						Name:        "name",
						Description: "The name of your service",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServiceName"),
					},
					{
						Name:        "status",
						Description: "The status of the service",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The metadata that you apply to the service to help you categorize and organize them",
						Type:        schema.TypeJSON,
						Resolver:    client.ResolveTags,
					},
					{
						Name:        "task_definition",
						Description: "The task definition to use for tasks in the service",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployments",
						Description: "The details of an Amazon ECS service deployment",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Deployments"),
					},
					{
						Name:        "events",
						Description: "The details for an event that's associated with a service.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Events"),
					},
					{
						Name:        "load_balancers",
						Description: "The load balancer configuration to use with a service or task set",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("LoadBalancers"),
					},
					{
						Name:        "service_registries",
						Description: "The details for the service registry",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("ServiceRegistries"),
					},
					{
						Name:        "task_sets",
						Description: "Information about a set of Amazon ECS tasks in either an CodeDeploy or an EXTERNAL deployment",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("TaskSets"),
					},
				},
			},
			{
				Name:          "aws_ecs_cluster_container_instances",
				Description:   "An EC2 instance that's running the Amazon ECS agent and has been registered with a cluster.",
				Resolver:      fetchEcsClusterContainerInstances,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_ecs_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "agent_connected",
						Description: "This parameter returns true if the agent is connected to Amazon ECS",
						Type:        schema.TypeBool,
					},
					{
						Name:        "agent_update_status",
						Description: "The status of the most recent agent update",
						Type:        schema.TypeString,
					},
					{
						Name:        "capacity_provider_name",
						Description: "The capacity provider that's associated with the container instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "container_instance_arn",
						Description: "The Amazon Resource Name (ARN) of the container instance",
						Type:        schema.TypeString,
					},
					{
						Name:        "ec2_instance_id",
						Description: "The ID of the container instance",
						Type:        schema.TypeString,
					},
					{
						Name:     "health_status",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("HealthStatus"),
					},
					{
						Name:        "pending_tasks_count",
						Description: "The number of tasks on the container instance that are in the PENDING status.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "registered_at",
						Description: "The Unix timestamp for the time when the container instance was registered.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "running_tasks_count",
						Description: "The number of tasks on the container instance that are in the RUNNING status.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "status",
						Description: "The status of the container instance",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_reason",
						Description: "The reason that the container instance reached its current status.",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The metadata that you apply to the container instance to help you categorize and organize them",
						Type:        schema.TypeJSON,
						Resolver:    client.ResolveTags,
					},
					{
						Name:        "version",
						Description: "The version counter for the container instance",
						Type:        schema.TypeInt,
					},
					{
						Name:        "version_info_agent_hash",
						Description: "The Git commit hash for the Amazon ECS container agent build on the amazon-ecs-agent  (https://github.com/aws/amazon-ecs-agent/commits/master) GitHub repository.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionInfo.AgentHash"),
					},
					{
						Name:     "version_info",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("VersionInfo"),
					},
					{
						Name:        "version_info_docker_version",
						Description: "The Docker version that's running on the container instance.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionInfo.DockerVersion"),
					},
					{
						Name:        "attachments",
						Description: "An object representing a container instance or task attachment.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Attachments"),
					},
					{
						Name:        "attributes",
						Description: "An attribute is a name-value pair that's associated with an Amazon ECS object. Use attributes to extend the Amazon ECS data model by adding custom metadata to your resources",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Attributes"),
					},
					{
						Name:        "registered_resources",
						Description: "Describes the resources available for a container instance.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("RegisteredResources"),
					},
					{
						Name:        "remaining_resources",
						Description: "Describes the resources available for a container instance.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("RemainingResources"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ecs.ListClustersInput
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	for {
		listClustersOutput, err := svc.ListClusters(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listClustersOutput.ClusterArns) == 0 {
			return nil
		}
		describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: listClustersOutput.ClusterArns,
			Include:  []types.ClusterField{types.ClusterFieldTags},
		}, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		res <- describeClusterOutput.Clusters

		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}

func resolveClustersSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", resource.Item)
	}
	settings := make(map[string]*string)
	for _, s := range cluster.Settings {
		settings[string(s.Name)] = s.Value
	}
	return resource.Set(c.Name, settings)
}
func resolveClustersStatistics(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", resource.Item)
	}
	stats := make(map[string]*string)
	for _, s := range cluster.Statistics {
		stats[*s.Name] = s.Value
	}
	return resource.Set(c.Name, stats)
}
func fetchEcsClusterTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", parent.Item)
	}
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	config := ecs.ListTasksInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listTasks, err := svc.ListTasks(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listTasks.TaskArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeTasksInput{
			Cluster: cluster.ClusterArn,
			Tasks:   listTasks.TaskArns,
			Include: []types.TaskField{types.TaskFieldTags},
		}
		describeTasks, err := svc.DescribeTasks(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		res <- describeTasks.Tasks

		if listTasks.NextToken == nil {
			break
		}
		config.NextToken = listTasks.NextToken
	}
	return nil
}
func fetchEcsClusterServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	config := ecs.ListServicesInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listServicesOutput, err := svc.ListServices(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listServicesOutput.ServiceArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeServicesInput{
			Cluster:  cluster.ClusterArn,
			Services: listServicesOutput.ServiceArns,
			Include:  []types.ServiceField{types.ServiceFieldTags},
		}
		describeServicesOutput, err := svc.DescribeServices(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		res <- describeServicesOutput.Services

		if listServicesOutput.NextToken == nil {
			break
		}
		config.NextToken = listServicesOutput.NextToken
	}
	return nil
}

func resolveClusterServicesPlacementConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.PlacementConstraints {
		j[string(i.Type)] = aws.ToString(i.Expression)
	}

	return resource.Set(c.Name, j)
}

func resolveClusterServicesPlacementStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.PlacementStrategy {
		j[string(i.Type)] = aws.ToString(i.Field)
	}

	return resource.Set(c.Name, j)
}

func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	config := ecs.ListContainerInstancesInput{
		Cluster: cluster.ClusterArn,
	}
	for {
		listContainerInstances, err := svc.ListContainerInstances(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		if len(listContainerInstances.ContainerInstanceArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeContainerInstancesInput{
			Cluster:            cluster.ClusterArn,
			ContainerInstances: listContainerInstances.ContainerInstanceArns,
			Include:            []types.ContainerInstanceField{types.ContainerInstanceFieldTags},
		}
		describeContainerInstances, err := svc.DescribeContainerInstances(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}

		res <- describeContainerInstances.ContainerInstances

		if listContainerInstances.NextToken == nil {
			break
		}
		config.NextToken = listContainerInstances.NextToken
	}
	return nil
}
