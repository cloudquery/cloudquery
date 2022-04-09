package ecs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_ecs_clusters",
		Description:  "A regional grouping of one or more container instances where you can run task requests",
		Resolver:     fetchEcsClusters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ecs"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterArn"),
			},
			{
				Name:        "name",
				Description: "A user-generated string that you use to identify your cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterName"),
			},
			{
				Name:          "execute_config_kms_key_id",
				Description:   "Specify an Key Management Service key ID to encrypt the data between the local client and the container.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ExecuteCommandConfiguration.KmsKeyId"),
				IgnoreInTests: true,
			},
			{
				Name:        "execute_config_logs_cloud_watch_encryption_enabled",
				Description: "Determines whether to use encryption on the CloudWatch logs",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled"),
			},
			{
				Name:          "execute_config_log_cloud_watch_log_group_name",
				Description:   "The name of the CloudWatch log group to send logs to",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName"),
				IgnoreInTests: true,
			},
			{
				Name:          "execute_config_log_s3_bucket_name",
				Description:   "The name of the S3 bucket to send logs to",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName"),
				IgnoreInTests: true,
			},
			{
				Name:        "execute_config_log_s3_encryption_enabled",
				Description: "Determines whether to use encryption on the S3 logs",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled"),
			},
			{
				Name:          "execute_config_log_s3_key_prefix",
				Description:   "An optional folder in the S3 bucket to place logs in.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix"),
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
				Resolver:    resolveClustersDefaultCapacityProviderStrategy,
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
				Resolver:    resolveClustersTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ecs_cluster_attachments",
				Description:   "An object representing a container instance or task attachment.",
				Resolver:      fetchEcsClusterAttachments,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_ecs_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "details",
						Description: "Details of the attachment",
						Type:        schema.TypeJSON,
						Resolver:    resolveClusterAttachmentsDetails,
					},
					{
						Name:        "id",
						Description: "The unique identifier for the attachment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the attachment",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the attachment, such as ElasticNetworkInterface.",
						Type:        schema.TypeString,
					},
				},
			},
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
						Resolver:    resolveClusterTasksAttributes,
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
						Name:        "ephemeral_storage_size_in_gib",
						Description: "The total amount, in GiB, of ephemeral storage to set for the task",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralStorage.SizeInGiB"),
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
						Resolver:    resolveClusterTasksInferenceAccelerators,
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
						Resolver:    resolveClusterTasksOverrides,
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
						Resolver:    resolveClusterTasksTags,
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
						Type:        schema.TypeBigInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ecs_cluster_task_attachments",
						Description: "An object representing a container instance or task attachment.",
						Resolver:    fetchEcsClusterTaskAttachments,
						Columns: []schema.Column{
							{
								Name:        "cluster_task_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_tasks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "details",
								Description: "Details of the attachment",
								Type:        schema.TypeJSON,
								Resolver:    resolveClusterTaskAttachmentsDetails,
							},
							{
								Name:        "id",
								Description: "The unique identifier for the attachment.",
								Type:        schema.TypeString,
							},
							{
								Name:        "status",
								Description: "The status of the attachment",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "The type of the attachment, such as ElasticNetworkInterface.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ecs_cluster_task_containers",
						Description: "A Docker container that's part of a task.",
						Resolver:    fetchEcsClusterTaskContainers,
						Columns: []schema.Column{
							{
								Name:        "cluster_task_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_tasks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "container_arn",
								Description: "The Amazon Resource Name (ARN) of the container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "cpu",
								Description: "The number of CPU units set for the container",
								Type:        schema.TypeString,
							},
							{
								Name:        "exit_code",
								Description: "The exit code returned from the container.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "gpu_ids",
								Description: "The IDs of each GPU assigned to the container.",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "health_status",
								Description: "The health status of the container",
								Type:        schema.TypeString,
							},
							{
								Name:        "image",
								Description: "The image used for the container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "image_digest",
								Description: "The container image manifest digest",
								Type:        schema.TypeString,
							},
							{
								Name:        "last_status",
								Description: "The last known status of the container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "managed_agents",
								Description: "The details of any Amazon ECS managed agents associated with the container.",
								Type:        schema.TypeJSON,
								Resolver:    resolveClusterTaskContainersManagedAgents,
							},
							{
								Name:        "memory",
								Description: "The hard limit (in MiB) of memory set for the container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "memory_reservation",
								Description: "The soft limit (in MiB) of memory set for the container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "name",
								Description: "The name of the container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "network_bindings",
								Description: "The network bindings associated with the container.",
								Type:        schema.TypeJSON,
								Resolver:    resolveClusterTaskContainersNetworkBindings,
							},
							{
								Name:        "network_interfaces",
								Description: "The network interfaces associated with the container.",
								Type:        schema.TypeJSON,
								Resolver:    resolveClusterTaskContainersNetworkInterfaces,
							},
							{
								Name:        "reason",
								Description: "A short (255 max characters) human-readable string to provide additional details about a running or stopped container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "runtime_id",
								Description: "The ID of the Docker container.",
								Type:        schema.TypeString,
							},
							{
								Name:        "task_arn",
								Description: "The ARN of the task.",
								Type:        schema.TypeString,
							},
						},
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
						Resolver:      resolveClusterServicesCapacityProviderStrategy,
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
						Name:        "deployment_configuration_deployment_circuit_breaker_enable",
						Description: "Determines whether to use the deployment circuit breaker logic for the service.  This member is required.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DeploymentConfiguration.DeploymentCircuitBreaker.Enable"),
					},
					{
						Name:        "deployment_configuration_deployment_circuit_breaker_rollback",
						Description: "Determines whether to configure Amazon ECS to roll back the service if a service deployment fails",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DeploymentConfiguration.DeploymentCircuitBreaker.Rollback"),
					},
					{
						Name:        "deployment_configuration_maximum_percent",
						Description: "If a service is using the rolling update (ECS) deployment type, the maximum percent parameter represents an upper limit on the number of tasks in a service that are allowed in the RUNNING or PENDING state during a deployment, as a percentage of the desired number of tasks (rounded down to the nearest integer), and while any container instances are in the DRAINING state if the service contains tasks using the EC2 launch type",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DeploymentConfiguration.MaximumPercent"),
					},
					{
						Name:        "deployment_configuration_minimum_healthy_percent",
						Description: "If a service is using the rolling update (ECS) deployment type, the minimum healthy percent represents a lower limit on the number of tasks in a service that must remain in the RUNNING state during a deployment, as a percentage of the desired number of tasks (rounded up to the nearest integer), and while any container instances are in the DRAINING state if the service contains tasks using the EC2 launch type",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DeploymentConfiguration.MinimumHealthyPercent"),
					},
					{
						Name:        "deployment_controller_type",
						Description: "The deployment controller type to use",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DeploymentController.Type"),
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
						Name:        "network_configuration_awsvpc_configuration_subnets",
						Description: "The IDs of the subnets associated with the task or service",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.Subnets"),
					},
					{
						Name:        "network_configuration_awsvpc_configuration_assign_public_ip",
						Description: "Whether the task's elastic network interface receives a public IP address",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.AssignPublicIp"),
					},
					{
						Name:        "network_configuration_awsvpc_configuration_security_groups",
						Description: "The IDs of the security groups associated with the task or service",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.SecurityGroups"),
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
						Resolver:    resolveClusterServicesTags,
					},
					{
						Name:        "task_definition",
						Description: "The task definition to use for tasks in the service",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ecs_cluster_service_deployments",
						Description: "The details of an Amazon ECS service deployment",
						Resolver:    fetchEcsClusterServiceDeployments,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:          "capacity_provider_strategy",
								Description:   "The capacity provider strategy that the deployment is using.",
								Type:          schema.TypeJSON,
								Resolver:      resolveClusterServiceDeploymentsCapacityProviderStrategy,
								IgnoreInTests: true,
							},
							{
								Name:        "created_at",
								Description: "The Unix timestamp for the time when the service deployment was created.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "desired_count",
								Description: "The most recent desired count of tasks that was specified for the service to deploy or maintain.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "failed_tasks",
								Description: "The number of consecutively failed tasks in the deployment",
								Type:        schema.TypeInt,
							},
							{
								Name:        "id",
								Description: "The ID of the deployment.",
								Type:        schema.TypeString,
							},
							{
								Name:        "launch_type",
								Description: "The launch type the tasks in the service are using",
								Type:        schema.TypeString,
							},
							{
								Name:        "network_configuration_awsvpc_configuration_subnets",
								Description: "The IDs of the subnets associated with the task or service",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.Subnets"),
							},
							{
								Name:        "network_configuration_awsvpc_configuration_assign_public_ip",
								Description: "Whether the task's elastic network interface receives a public IP address",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.AssignPublicIp"),
							},
							{
								Name:        "network_configuration_awsvpc_configuration_security_groups",
								Description: "The IDs of the security groups associated with the task or service",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.SecurityGroups"),
							},
							{
								Name:        "pending_count",
								Description: "The number of tasks in the deployment that are in the PENDING status.",
								Type:        schema.TypeInt,
							},
							{
								Name:          "platform_family",
								Description:   "The operating system that your tasks in the service, or tasks are running on",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "platform_version",
								Description:   "The platform version that your tasks in the service run on",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "rollout_state",
								Description: "The rolloutState of a service is only returned for services that use the rolling update (ECS) deployment type that aren't behind a Classic Load Balancer",
								Type:        schema.TypeString,
							},
							{
								Name:        "rollout_state_reason",
								Description: "A description of the rollout state of a deployment.",
								Type:        schema.TypeString,
							},
							{
								Name:        "running_count",
								Description: "The number of tasks in the deployment that are in the RUNNING status.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "status",
								Description: "The status of the deployment",
								Type:        schema.TypeString,
							},
							{
								Name:        "task_definition",
								Description: "The most recent task definition that was specified for the tasks in the service to use.",
								Type:        schema.TypeString,
							},
							{
								Name:        "updated_at",
								Description: "The Unix timestamp for the time when the service deployment was last updated.",
								Type:        schema.TypeTimestamp,
							},
						},
					},
					{
						Name:        "aws_ecs_cluster_service_events",
						Description: "The details for an event that's associated with a service.",
						Resolver:    fetchEcsClusterServiceEvents,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "created_at",
								Description: "The Unix timestamp for the time when the event was triggered.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "id",
								Description: "The ID string for the event.",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "The event message.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ecs_cluster_service_load_balancers",
						Description:   "The load balancer configuration to use with a service or task set",
						Resolver:      fetchEcsClusterServiceLoadBalancers,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "container_name",
								Description: "The name of the container (as it appears in a container definition) to associate with the load balancer.",
								Type:        schema.TypeString,
							},
							{
								Name:        "container_port",
								Description: "The port on the container to associate with the load balancer",
								Type:        schema.TypeInt,
							},
							{
								Name:          "load_balancer_name",
								Description:   "The name of the load balancer to associate with the Amazon ECS service or task set",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "target_group_arn",
								Description: "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ecs_cluster_service_service_registries",
						Description:   "The details for the service registry",
						Resolver:      fetchEcsClusterServiceServiceRegistries,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "container_name",
								Description: "The container name value to be used for your service discovery service",
								Type:        schema.TypeString,
							},
							{
								Name:        "container_port",
								Description: "The port value to be used for your service discovery service",
								Type:        schema.TypeInt,
							},
							{
								Name:        "port",
								Description: "The port value used if your service discovery service specified an SRV record. This field might be used if both the awsvpc network mode and SRV records are used.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "registry_arn",
								Description: "The Amazon Resource Name (ARN) of the service registry",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ecs_cluster_service_task_sets",
						Description:   "Information about a set of Amazon ECS tasks in either an CodeDeploy or an EXTERNAL deployment",
						Resolver:      fetchEcsClusterServiceTaskSets,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:          "capacity_provider_strategy",
								Description:   "The capacity provider strategy that are associated with the task set.",
								Type:          schema.TypeJSON,
								Resolver:      resolveClusterServiceTaskSetsCapacityProviderStrategy,
								IgnoreInTests: true,
							},
							{
								Name:        "cluster_arn",
								Description: "The Amazon Resource Name (ARN) of the cluster that the service that hosts the task set exists in.",
								Type:        schema.TypeString,
							},
							{
								Name:        "computed_desired_count",
								Description: "The computed desired count for the task set",
								Type:        schema.TypeInt,
							},
							{
								Name:        "created_at",
								Description: "The Unix timestamp for the time when the task set was created.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "external_id",
								Description: "The external ID associated with the task set",
								Type:        schema.TypeString,
							},
							{
								Name:        "id",
								Description: "The ID of the task set.",
								Type:        schema.TypeString,
							},
							{
								Name:        "launch_type",
								Description: "The launch type the tasks in the task set are using",
								Type:        schema.TypeString,
							},
							{
								Name:        "network_configuration_awsvpc_configuration_subnets",
								Description: "The IDs of the subnets associated with the task or service",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.Subnets"),
							},
							{
								Name:        "network_configuration_awsvpc_configuration_assign_public_ip",
								Description: "Whether the task's elastic network interface receives a public IP address",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.AssignPublicIp"),
							},
							{
								Name:        "network_configuration_awsvpc_configuration_security_groups",
								Description: "The IDs of the security groups associated with the task or service",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("NetworkConfiguration.AwsvpcConfiguration.SecurityGroups"),
							},
							{
								Name:        "pending_count",
								Description: "The number of tasks in the task set that are in the PENDING status during a deployment",
								Type:        schema.TypeInt,
							},
							{
								Name:          "platform_family",
								Description:   "The operating system that your tasks in the set are running on",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "platform_version",
								Description:   "The Fargate platform version where the tasks in the task set are running",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "running_count",
								Description: "The number of tasks in the task set that are in the RUNNING status during a deployment",
								Type:        schema.TypeInt,
							},
							{
								Name:        "scale_unit",
								Description: "The unit of measure for the scale value.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Scale.Unit"),
							},
							{
								Name:        "scale_value",
								Description: "The value, specified as a percent total of a service's desiredCount, to scale the task set",
								Type:        schema.TypeFloat,
								Resolver:    schema.PathResolver("Scale.Value"),
							},
							{
								Name:        "service_arn",
								Description: "The Amazon Resource Name (ARN) of the service the task set exists in.",
								Type:        schema.TypeString,
							},
							{
								Name:        "stability_status",
								Description: "The stability status",
								Type:        schema.TypeString,
							},
							{
								Name:        "stability_status_at",
								Description: "The Unix timestamp for the time when the task set stability status was retrieved.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "started_by",
								Description: "The tag specified when a task set is started",
								Type:        schema.TypeString,
							},
							{
								Name:        "status",
								Description: "The status of the task set",
								Type:        schema.TypeString,
							},
							{
								Name:        "tags",
								Description: "The metadata that you apply to the task set to help you categorize and organize them",
								Type:        schema.TypeJSON,
								Resolver:    resolveClusterServiceTaskSetsTags,
							},
							{
								Name:        "task_definition",
								Description: "The task definition that the task set is using.",
								Type:        schema.TypeString,
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) of the task set.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TaskSetArn"),
							},
							{
								Name:        "updated_at",
								Description: "The Unix timestamp for the time when the task set was last updated.",
								Type:        schema.TypeTimestamp,
							},
						},
						Relations: []*schema.Table{
							{
								Name:          "aws_ecs_cluster_service_task_set_load_balancers",
								Description:   "The load balancer configuration to use with a service or task set",
								Resolver:      fetchEcsClusterServiceTaskSetLoadBalancers,
								IgnoreInTests: true,
								Columns: []schema.Column{
									{
										Name:        "cluster_service_task_set_cq_id",
										Description: "Unique CloudQuery ID of aws_ecs_cluster_service_task_sets table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "container_name",
										Description: "The name of the container (as it appears in a container definition) to associate with the load balancer.",
										Type:        schema.TypeString,
									},
									{
										Name:        "container_port",
										Description: "The port on the container to associate with the load balancer",
										Type:        schema.TypeInt,
									},
									{
										Name:        "load_balancer_name",
										Description: "The name of the load balancer to associate with the Amazon ECS service or task set",
										Type:        schema.TypeString,
									},
									{
										Name:        "target_group_arn",
										Description: "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set",
										Type:        schema.TypeString,
									},
								},
							},
							{
								Name:          "aws_ecs_cluster_service_task_set_service_registries",
								Description:   "The details for the service registry",
								Resolver:      fetchEcsClusterServiceTaskSetServiceRegistries,
								IgnoreInTests: true,
								Columns: []schema.Column{
									{
										Name:        "cluster_service_task_set_cq_id",
										Description: "Unique CloudQuery ID of aws_ecs_cluster_service_task_sets table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "container_name",
										Description: "The container name value to be used for your service discovery service",
										Type:        schema.TypeString,
									},
									{
										Name:        "container_port",
										Description: "The port value to be used for your service discovery service",
										Type:        schema.TypeInt,
									},
									{
										Name:        "port",
										Description: "The port value used if your service discovery service specified an SRV record. This field might be used if both the awsvpc network mode and SRV records are used.",
										Type:        schema.TypeInt,
									},
									{
										Name:        "arn",
										Description: "The Amazon Resource Name (ARN) of the service registry",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("RegistryArn"),
									},
								},
							},
						},
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
						Name:        "health_status_overall_status",
						Description: "The overall health status of the container instance",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthStatus.OverallStatus"),
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
						Resolver:    resolveClusterContainerInstancesTags,
					},
					{
						Name:        "version",
						Description: "The version counter for the container instance",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "version_info_agent_hash",
						Description: "The Git commit hash for the Amazon ECS container agent build on the amazon-ecs-agent  (https://github.com/aws/amazon-ecs-agent/commits/master) GitHub repository.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionInfo.AgentHash"),
					},
					{
						Name:        "version_info_agent_version",
						Description: "The version number of the Amazon ECS container agent.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionInfo.AgentVersion"),
					},
					{
						Name:        "version_info_docker_version",
						Description: "The Docker version that's running on the container instance.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionInfo.DockerVersion"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_ecs_cluster_container_instance_attachments",
						Description:   "An object representing a container instance or task attachment.",
						Resolver:      fetchEcsClusterContainerInstanceAttachments,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_container_instance_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "details",
								Description: "Details of the attachment",
								Type:        schema.TypeJSON,
								Resolver:    resolveClusterContainerInstanceAttachmentsDetails,
							},
							{
								Name:        "id",
								Description: "The unique identifier for the attachment.",
								Type:        schema.TypeString,
							},
							{
								Name:        "status",
								Description: "The status of the attachment",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "The type of the attachment, such as ElasticNetworkInterface.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ecs_cluster_container_instance_attributes",
						Description:   "An attribute is a name-value pair that's associated with an Amazon ECS object. Use attributes to extend the Amazon ECS data model by adding custom metadata to your resources",
						Resolver:      fetchEcsClusterContainerInstanceAttributes,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_container_instance_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The name of the attribute",
								Type:        schema.TypeString,
							},
							{
								Name:        "target_id",
								Description: "The ID of the target",
								Type:        schema.TypeString,
							},
							{
								Name:        "target_type",
								Description: "The type of the target to attach the attribute with",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The value of the attribute",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ecs_cluster_container_instance_health_status_details",
						Description: "An object representing the result of a container instance health status check.",
						Resolver:    fetchEcsClusterContainerInstanceHealthStatusDetails,
						Columns: []schema.Column{
							{
								Name:        "cluster_container_instance_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "last_status_change",
								Description: "The Unix timestamp for when the container instance health status last changed.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "last_updated",
								Description: "The Unix timestamp for when the container instance health status was last updated.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "status",
								Description: "The container instance health status.",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "The type of container instance health status that was verified.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ecs_cluster_container_instance_registered_resources",
						Description:   "Describes the resources available for a container instance.",
						Resolver:      fetchEcsClusterContainerInstanceRegisteredResources,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_container_instance_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "double_value",
								Description: "When the doubleValue type is set, the value of the resource must be a double precision floating-point type.",
								Type:        schema.TypeFloat,
							},
							{
								Name:        "integer_value",
								Description: "When the integerValue type is set, the value of the resource must be an integer.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "long_value",
								Description: "When the longValue type is set, the value of the resource must be an extended precision floating-point type.",
								Type:        schema.TypeBigInt,
							},
							{
								Name:        "name",
								Description: "The name of the resource, such as CPU, MEMORY, PORTS, PORTS_UDP, or a user-defined resource.",
								Type:        schema.TypeString,
							},
							{
								Name:        "string_set_value",
								Description: "When the stringSetValue type is set, the value of the resource must be a string type.",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "type",
								Description: "The type of the resource",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ecs_cluster_container_instance_remaining_resources",
						Description:   "Describes the resources available for a container instance.",
						Resolver:      fetchEcsClusterContainerInstanceRemainingResources,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "cluster_container_instance_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "double_value",
								Description: "When the doubleValue type is set, the value of the resource must be a double precision floating-point type.",
								Type:        schema.TypeFloat,
							},
							{
								Name:        "integer_value",
								Description: "When the integerValue type is set, the value of the resource must be an integer.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "long_value",
								Description: "When the longValue type is set, the value of the resource must be an extended precision floating-point type.",
								Type:        schema.TypeBigInt,
							},
							{
								Name:        "name",
								Description: "The name of the resource, such as CPU, MEMORY, PORTS, PORTS_UDP, or a user-defined resource.",
								Type:        schema.TypeString,
							},
							{
								Name:        "string_set_value",
								Description: "When the stringSetValue type is set, the value of the resource must be a string type.",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "type",
								Description: "The type of the resource",
								Type:        schema.TypeString,
							},
						},
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
			return diag.WrapError(err)
		}
		if len(listClustersOutput.ClusterArns) == 0 {
			return nil
		}
		describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{Clusters: listClustersOutput.ClusterArns}, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- describeClusterOutput.Clusters

		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}
func resolveClustersDefaultCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", resource.Item))
	}
	data, err := json.Marshal(cluster.DefaultCapacityProviderStrategy)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClustersSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", resource.Item))
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
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", resource.Item))
	}
	stats := make(map[string]*string)
	for _, s := range cluster.Statistics {
		stats[*s.Name] = s.Value
	}
	return resource.Set(c.Name, stats)
}
func resolveClustersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", resource.Item))
	}
	listTagsForResourceOutput, err := svc.ListTagsForResource(ctx, &ecs.ListTagsForResourceInput{
		ResourceArn: cluster.ClusterArn,
	}, func(o *ecs.Options) {
		o.Region = region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	tags := make(map[string]*string)
	for _, s := range listTagsForResourceOutput.Tags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
func fetchEcsClusterAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", parent.Item))
	}
	res <- cluster.Attachments
	return nil
}
func resolveClusterAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	attachment, ok := resource.Item.(types.Attachment)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Attachment but got %T", resource.Item))
	}
	details := make(map[string]*string)
	for _, s := range attachment.Details {
		details[*s.Name] = s.Value
	}
	return resource.Set(c.Name, details)
}
func fetchEcsClusterTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", parent.Item))
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
			return diag.WrapError(err)
		}
		if len(listTasks.TaskArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeTasksInput{
			Cluster: cluster.ClusterArn,
			Tasks:   listTasks.TaskArns,
		}
		describeTasks, err := svc.DescribeTasks(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- describeTasks.Tasks

		if listTasks.NextToken == nil {
			break
		}
		config.NextToken = listTasks.NextToken
	}
	return nil
}
func resolveClusterTasksAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Task)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Task but got %T", resource.Item))
	}
	data, err := json.Marshal(p.Attributes)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterTasksInferenceAccelerators(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Task)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Task but got %T", resource.Item))
	}
	data, err := json.Marshal(p.InferenceAccelerators)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterTasksOverrides(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Task)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Task but got %T", resource.Item))
	}
	if p.Overrides == nil {
		return nil
	}
	data, err := json.Marshal(p.Overrides)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterTasksTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Task)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Task but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range p.Tags {
		j[*i.Key] = *i.Value
	}

	return resource.Set(c.Name, j)
}
func fetchEcsClusterTaskAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.Task)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Task but got %T", parent.Item))
	}
	res <- p.Attachments
	return nil
}
func resolveClusterTaskAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Attachment)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Attachment but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range p.Details {
		j[*i.Name] = *i.Value
	}

	return resource.Set(c.Name, j)
}
func fetchEcsClusterTaskContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.Task)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Task but got %T", parent.Item))
	}
	res <- p.Containers
	return nil
}
func resolveClusterTaskContainersManagedAgents(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Container)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Container but got %T", resource.Item))
	}
	data, err := json.Marshal(p.ManagedAgents)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterTaskContainersNetworkBindings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Container)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Container but got %T", resource.Item))
	}
	data, err := json.Marshal(p.NetworkBindings)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterTaskContainersNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Container)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Container but got %T", resource.Item))
	}
	data, err := json.Marshal(p.NetworkInterfaces)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func fetchEcsClusterServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", parent.Item))
	}
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
			return diag.WrapError(err)
		}
		if len(listServicesOutput.ServiceArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeServicesInput{
			Cluster:  cluster.ClusterArn,
			Services: listServicesOutput.ServiceArns,
		}
		describeServicesOutput, err := svc.DescribeServices(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- describeServicesOutput.Services

		if listServicesOutput.NextToken == nil {
			break
		}
		config.NextToken = listServicesOutput.NextToken
	}
	return nil
}
func resolveClusterServicesCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", resource.Item))
	}
	data, err := json.Marshal(service.CapacityProviderStrategy)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterServicesPlacementConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range service.PlacementConstraints {
		j[string(i.Type)] = *i.Expression
	}

	return resource.Set(c.Name, j)
}
func resolveClusterServicesPlacementStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range service.PlacementStrategy {
		j[string(i.Type)] = *i.Field
	}

	return resource.Set(c.Name, j)
}
func resolveClusterServicesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range service.Tags {
		j[*i.Key] = *i.Value
	}

	return resource.Set(c.Name, j)
}
func fetchEcsClusterServiceDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", parent.Item))
	}
	res <- service.Deployments
	return nil
}
func resolveClusterServiceDeploymentsCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	deployment, ok := resource.Item.(types.Deployment)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Deployment but got %T", resource.Item))
	}
	data, err := json.Marshal(deployment.CapacityProviderStrategy)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func fetchEcsClusterServiceEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", parent.Item))
	}
	res <- service.Events
	return nil
}
func fetchEcsClusterServiceLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", parent.Item))
	}
	res <- service.LoadBalancers
	return nil
}
func fetchEcsClusterServiceServiceRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", parent.Item))
	}
	res <- service.ServiceRegistries
	return nil
}
func fetchEcsClusterServiceTaskSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(types.Service)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Service but got %T", parent.Item))
	}
	res <- service.TaskSets
	return nil
}
func resolveClusterServiceTaskSetsCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	taskSet, ok := resource.Item.(types.TaskSet)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.TaskSet but got %T", resource.Item))
	}
	data, err := json.Marshal(taskSet.CapacityProviderStrategy)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}
func resolveClusterServiceTaskSetsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	taskSet, ok := resource.Item.(types.TaskSet)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.TaskSet but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range taskSet.Tags {
		j[*i.Key] = *i.Value
	}
	return resource.Set(c.Name, j)
}
func fetchEcsClusterServiceTaskSetLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	taskSet, ok := parent.Item.(types.TaskSet)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.TaskSet but got %T", parent.Item))
	}
	res <- taskSet.LoadBalancers
	return nil
}
func fetchEcsClusterServiceTaskSetServiceRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	taskSet, ok := parent.Item.(types.TaskSet)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.TaskSet but got %T", parent.Item))
	}
	res <- taskSet.ServiceRegistries
	return nil
}
func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.Cluster but got %T", parent.Item))
	}
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
			return diag.WrapError(err)
		}
		if len(listContainerInstances.ContainerInstanceArns) == 0 {
			return nil
		}
		describeServicesInput := ecs.DescribeContainerInstancesInput{
			Cluster:            cluster.ClusterArn,
			ContainerInstances: listContainerInstances.ContainerInstanceArns,
		}
		describeContainerInstances, err := svc.DescribeContainerInstances(ctx, &describeServicesInput, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- describeContainerInstances.ContainerInstances

		if listContainerInstances.NextToken == nil {
			break
		}
		config.NextToken = listContainerInstances.NextToken
	}
	return nil
}
func resolveClusterContainerInstancesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance, ok := resource.Item.(types.ContainerInstance)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", resource.Item))
	}
	j := make(map[string]interface{})
	for _, i := range instance.Tags {
		j[*i.Key] = *i.Value
	}
	return resource.Set(c.Name, j)
}
func fetchEcsClusterContainerInstanceAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.ContainerInstance)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", parent.Item))
	}
	res <- instance.Attachments
	return nil
}
func resolveClusterContainerInstanceAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	attachment, ok := resource.Item.(types.Attachment)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", resource.Item))
	}
	details := make(map[string]*string)
	for _, s := range attachment.Details {
		details[*s.Name] = s.Value
	}
	return resource.Set(c.Name, details)
}
func fetchEcsClusterContainerInstanceAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.ContainerInstance)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", parent.Item))
	}
	res <- instance.Attributes
	return nil
}
func fetchEcsClusterContainerInstanceHealthStatusDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.ContainerInstance)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", parent.Item))
	}
	if instance.HealthStatus == nil || instance.HealthStatus.Details == nil {
		return nil
	}
	res <- instance.HealthStatus.Details
	return nil
}
func fetchEcsClusterContainerInstanceRegisteredResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.ContainerInstance)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", parent.Item))
	}
	res <- instance.RegisteredResources
	return nil
}
func fetchEcsClusterContainerInstanceRemainingResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.ContainerInstance)
	if !ok {
		return diag.WrapError(fmt.Errorf("expected to have types.ContainerInstance but got %T", parent.Item))
	}
	res <- instance.RemainingResources
	return nil
}
