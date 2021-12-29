package ecs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EcsClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_ecs_clusters",
		Description:  "A regional grouping of one or more container instances on which you can run task requests",
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
				Name:        "attachments_status",
				Description: "The status of the capacity providers associated with the cluster",
				Type:        schema.TypeString,
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
				Name:        "execute_config_kms_key_id",
				Description: "Specify an AWS Key Management Service key ID to encrypt the data between the local client and the container.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.KmsKeyId"),
			},
			{
				Name:        "execute_config_logs_cloud_watch_encryption_enabled",
				Description: "Whether or not to enable encryption on the CloudWatch logs",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled"),
			},
			{
				Name:        "execute_config_log_cloud_watch_log_group_name",
				Description: "The name of the CloudWatch log group to send logs to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName"),
			},
			{
				Name:        "execute_config_log_s3_bucket_name",
				Description: "The name of the S3 bucket to send logs to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName"),
			},
			{
				Name:        "execute_config_log_s3_encryption_enabled",
				Description: "Whether or not to enable encryption on the CloudWatch logs",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled"),
			},
			{
				Name:        "execute_config_log_s3_key_prefix",
				Description: "An optional folder in the S3 bucket to place logs in.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix"),
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
				Resolver:    resolveEcsClustersDefaultCapacityProviderStrategy,
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
				Resolver:    resolveEcsClustersSettings,
			},
			{
				Name:        "statistics",
				Description: "Additional information about your clusters that are separated by launch type, including:  * runningEC2TasksCount  * RunningFargateTasksCount  * pendingEC2TasksCount  * pendingFargateTasksCount  * activeEC2ServiceCount  * activeFargateServiceCount  * drainingEC2ServiceCount  * drainingFargateServiceCount",
				Type:        schema.TypeJSON,
				Resolver:    resolveEcsClustersStatistics,
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
				Resolver:    resolveEcsClustersTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ecs_cluster_attachments",
				Description: "An object representing a container instance or task attachment.",
				Resolver:    fetchEcsClusterAttachments,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "id"}},
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
						Resolver:    resolveEcsClusterAttachmentsDetails,
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
						Name:        "capacity_provider_strategy",
						Description: "The capacity provider strategy associated with the service.",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsClusterServicesCapacityProviderStrategy,
					},
					{
						Name:        "cluster_arn",
						Description: "The Amazon Resource Name (ARN) of the cluster that hosts the service.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The Unix timestamp for when the service was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "created_by",
						Description: "The principal that created the service.",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_configuration_deployment_circuit_breaker_enable",
						Description: "Whether to enable the deployment circuit breaker logic for the service.  This member is required.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DeploymentConfiguration.DeploymentCircuitBreaker.Enable"),
					},
					{
						Name:        "deployment_configuration_deployment_circuit_breaker_rollback",
						Description: "Whether to enable Amazon ECS to roll back the service if a service deployment fails",
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
						Description: "Specifies whether to enable Amazon ECS managed tags for the tasks in the service",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EnableECSManagedTags"),
					},
					{
						Name:        "enable_execute_command",
						Description: "Whether or not the execute command functionality is enabled for the service",
						Type:        schema.TypeBool,
					},
					{
						Name:        "health_check_grace_period_seconds",
						Description: "The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "launch_type",
						Description: "The launch type on which your service is running",
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
						Resolver:    resolveEcsClusterServicesPlacementConstraints,
					},
					{
						Name:        "placement_strategy",
						Description: "The placement strategy that determines how tasks for the service are placed.",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsClusterServicesPlacementStrategy,
					},
					{
						Name:        "platform_version",
						Description: "The platform version on which to run your service",
						Type:        schema.TypeString,
					},
					{
						Name:        "propagate_tags",
						Description: "Specifies whether to propagate the tags from the task definition or the service to the task",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_arn",
						Description: "The ARN of the IAM role associated with the service that allows the Amazon ECS container agent to register container instances with an Elastic Load Balancing load balancer.",
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
						Resolver:    resolveEcsClusterServicesTags,
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
								Name:        "capacity_provider_strategy",
								Description: "The capacity provider strategy that the deployment is using.",
								Type:        schema.TypeJSON,
								Resolver:    resolveEcsClusterServiceDeploymentsCapacityProviderStrategy,
							},
							{
								Name:        "created_at",
								Description: "The Unix timestamp for when the service deployment was created.",
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
								Name:        "platform_version",
								Description: "The platform version on which your tasks in the service are running",
								Type:        schema.TypeString,
							},
							{
								Name:        "rollout_state",
								Description: "The rolloutState of a service is only returned for services that use the rolling update (ECS) deployment type that are not behind a Classic Load Balancer",
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
								Description: "The Unix timestamp for when the service deployment was last updated.",
								Type:        schema.TypeTimestamp,
							},
						},
					},
					{
						Name:        "aws_ecs_cluster_service_events",
						Description: "Details on an event associated with a service.",
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
								Description: "The Unix timestamp for when the event was triggered.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "id",
								Description: "The ID string of the event.",
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
						Name:        "aws_ecs_cluster_service_load_balancers",
						Description: "The load balancer configuration to use with a service or task set",
						Resolver:    fetchEcsClusterServiceLoadBalancers,
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
						Name:        "aws_ecs_cluster_service_service_registries",
						Description: "Details of the service registry.",
						Resolver:    fetchEcsClusterServiceServiceRegistries,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "container_name",
								Description: "The container name value, already specified in the task definition, to be used for your service discovery service",
								Type:        schema.TypeString,
							},
							{
								Name:        "container_port",
								Description: "The port value, already specified in the task definition, to be used for your service discovery service",
								Type:        schema.TypeInt,
							},
							{
								Name:        "port",
								Description: "The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.",
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
						Name:        "aws_ecs_cluster_service_task_sets",
						Description: "Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an EXTERNAL deployment",
						Resolver:    fetchEcsClusterServiceTaskSets,
						Columns: []schema.Column{
							{
								Name:        "cluster_service_cq_id",
								Description: "Unique CloudQuery ID of aws_ecs_cluster_services table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "capacity_provider_strategy",
								Description: "The capacity provider strategy associated with the task set.",
								Type:        schema.TypeJSON,
								Resolver:    resolveEcsClusterServiceTaskSetsCapacityProviderStrategy,
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
								Description: "The Unix timestamp for when the task set was created.",
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
								Name:        "platform_version",
								Description: "The AWS Fargate platform version on which the tasks in the task set are running. A platform version is only specified for tasks run on AWS Fargate",
								Type:        schema.TypeString,
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
								Description: "The stability status, which indicates whether the task set has reached a steady state",
								Type:        schema.TypeString,
							},
							{
								Name:        "stability_status_at",
								Description: "The Unix timestamp for when the task set stability status was retrieved.",
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
								Resolver:    resolveEcsClusterServiceTaskSetsTags,
							},
							{
								Name:        "task_definition",
								Description: "The task definition the task set is using.",
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
								Description: "The Unix timestamp for when the task set was last updated.",
								Type:        schema.TypeTimestamp,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_ecs_cluster_service_task_set_load_balancers",
								Description: "The load balancer configuration to use with a service or task set",
								Resolver:    fetchEcsClusterServiceTaskSetLoadBalancers,
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
								Name:        "aws_ecs_cluster_service_task_set_service_registries",
								Description: "Details of the service registry.",
								Resolver:    fetchEcsClusterServiceTaskSetServiceRegistries,
								Columns: []schema.Column{
									{
										Name:        "cluster_service_task_set_cq_id",
										Description: "Unique CloudQuery ID of aws_ecs_cluster_service_task_sets table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "container_name",
										Description: "The container name value, already specified in the task definition, to be used for your service discovery service",
										Type:        schema.TypeString,
									},
									{
										Name:        "container_port",
										Description: "The port value, already specified in the task definition, to be used for your service discovery service",
										Type:        schema.TypeInt,
									},
									{
										Name:        "port",
										Description: "The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.",
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
				Name:        "aws_ecs_cluster_container_instances",
				Description: "An EC2 instance that is running the Amazon ECS agent and has been registered with a cluster.",
				Resolver:    fetchEcsClusterContainerInstances,
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
						Description: "The capacity provider associated with the container instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "container_instance_arn",
						Description: "The Amazon Resource Name (ARN) of the container instance",
						Type:        schema.TypeString,
					},
					{
						Name:        "ec2_instance_id",
						Description: "The EC2 instance ID of the container instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "pending_tasks_count",
						Description: "The number of tasks on the container instance that are in the PENDING status.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "registered_at",
						Description: "The Unix timestamp for when the container instance was registered.",
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
						Resolver:    resolveEcsClusterContainerInstancesTags,
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
						Description: "The Docker version running on the container instance.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionInfo.DockerVersion"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ecs_cluster_container_instance_attachments",
						Description: "An object representing a container instance or task attachment.",
						Resolver:    fetchEcsClusterContainerInstanceAttachments,
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
								Resolver:    resolveEcsClusterContainerInstanceAttachmentsDetails,
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
						Name:        "aws_ecs_cluster_container_instance_attributes",
						Description: "An attribute is a name-value pair associated with an Amazon ECS object. Attributes enable you to extend the Amazon ECS data model by adding custom metadata to your resources",
						Resolver:    fetchEcsClusterContainerInstanceAttributes,
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
								Description: "The type of the target with which to attach the attribute",
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
						Name:        "aws_ecs_cluster_container_instance_registered_resources",
						Description: "Describes the resources available for a container instance.",
						Resolver:    fetchEcsClusterContainerInstanceRegisteredResources,
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
								Description: "The type of the resource, such as INTEGER, DOUBLE, LONG, or STRINGSET.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ecs_cluster_container_instance_remaining_resources",
						Description: "Describes the resources available for a container instance.",
						Resolver:    fetchEcsClusterContainerInstanceRemainingResources,
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
								Description: "The type of the resource, such as INTEGER, DOUBLE, LONG, or STRINGSET.",
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

func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
		describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{Clusters: listClustersOutput.ClusterArns}, func(o *ecs.Options) {
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
func resolveEcsClustersDefaultCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)
	data, err := json.Marshal(cluster.DefaultCapacityProviderStrategy)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveEcsClustersSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)
	settings := make(map[string]*string)
	for _, s := range cluster.Settings {
		settings[string(s.Name)] = s.Value
	}
	return resource.Set(c.Name, settings)
}
func resolveEcsClustersStatistics(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)
	stats := make(map[string]*string)
	for _, s := range cluster.Statistics {
		stats[*s.Name] = s.Value
	}
	return resource.Set(c.Name, stats)
}
func resolveEcsClustersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	cluster, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected to have types.Cluster but got %T", resource.Item)
	}
	listTagsForResourceOutput, err := svc.ListTagsForResource(ctx, &ecs.ListTagsForResourceInput{
		ResourceArn: cluster.ClusterArn,
	}, func(o *ecs.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}
	tags := make(map[string]*string)
	for _, s := range listTagsForResourceOutput.Tags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
func fetchEcsClusterAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster := parent.Item.(types.Cluster)
	res <- cluster.Attachments
	return nil
}
func resolveEcsClusterAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	attachment := resource.Item.(types.Attachment)
	details := make(map[string]*string)
	for _, s := range attachment.Details {
		details[*s.Name] = s.Value
	}
	return resource.Set(c.Name, details)
}
func fetchEcsClusterServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func resolveEcsClusterServicesCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	data, err := json.Marshal(service.CapacityProviderStrategy)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveEcsClusterServicesPlacementConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.PlacementConstraints {
		j[string(i.Type)] = *i.Expression
	}

	return resource.Set(c.Name, j)
}
func resolveEcsClusterServicesPlacementStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.PlacementStrategy {
		j[string(i.Type)] = *i.Field
	}

	return resource.Set(c.Name, j)
}
func resolveEcsClusterServicesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Service)
	j := make(map[string]interface{})
	for _, i := range service.Tags {
		j[*i.Key] = *i.Value
	}

	return resource.Set(c.Name, j)
}
func fetchEcsClusterServiceDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(types.Service)
	res <- service.Deployments
	return nil
}
func resolveEcsClusterServiceDeploymentsCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.Deployment)
	data, err := json.Marshal(service.CapacityProviderStrategy)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func fetchEcsClusterServiceEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(types.Service)
	res <- service.Events
	return nil
}
func fetchEcsClusterServiceLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(types.Service)
	res <- service.LoadBalancers
	return nil
}
func fetchEcsClusterServiceServiceRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(types.Service)
	res <- service.ServiceRegistries
	return nil
}
func fetchEcsClusterServiceTaskSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	service := parent.Item.(types.Service)
	res <- service.TaskSets
	return nil
}
func resolveEcsClusterServiceTaskSetsCapacityProviderStrategy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service := resource.Item.(types.TaskSet)
	data, err := json.Marshal(service.CapacityProviderStrategy)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveEcsClusterServiceTaskSetsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	taskSet := resource.Item.(types.TaskSet)
	j := make(map[string]interface{})
	for _, i := range taskSet.Tags {
		j[*i.Key] = *i.Value
	}
	return resource.Set(c.Name, j)
}
func fetchEcsClusterServiceTaskSetLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	taskSet := parent.Item.(types.TaskSet)
	res <- taskSet.LoadBalancers
	return nil
}
func fetchEcsClusterServiceTaskSetServiceRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	taskSet := parent.Item.(types.TaskSet)
	res <- taskSet.ServiceRegistries
	return nil
}
func fetchEcsClusterContainerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func resolveEcsClusterContainerInstancesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.ContainerInstance)
	j := make(map[string]interface{})
	for _, i := range instance.Tags {
		j[*i.Key] = *i.Value
	}
	return resource.Set(c.Name, j)
}
func fetchEcsClusterContainerInstanceAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance := parent.Item.(types.ContainerInstance)
	res <- instance.Attachments
	return nil
}
func resolveEcsClusterContainerInstanceAttachmentsDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	attachment := resource.Item.(types.Attachment)
	details := make(map[string]*string)
	for _, s := range attachment.Details {
		details[*s.Name] = s.Value
	}
	return resource.Set(c.Name, details)
}
func fetchEcsClusterContainerInstanceAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance := parent.Item.(types.ContainerInstance)
	res <- instance.Attributes
	return nil
}
func fetchEcsClusterContainerInstanceRegisteredResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance := parent.Item.(types.ContainerInstance)
	res <- instance.RegisteredResources
	return nil
}
func fetchEcsClusterContainerInstanceRemainingResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance := parent.Item.(types.ContainerInstance)
	res <- instance.RemainingResources
	return nil
}
