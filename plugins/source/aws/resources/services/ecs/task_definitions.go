package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
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
			return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listEcsTaskDefinitions, ecsTaskDefinitionDetail))
		},
		Multiplex:     client.ServiceAccountRegionMultiplexer("ecs"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The full Amazon Resource Name (ARN) of the task definition.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TaskDefinitionArn"),
			},
			{
				Name:        "task_role_arn",
				Description: "The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants containers in the task permission to call AWS APIs on your behalf",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ecs_task_definition_container_definitions",
				Description:   "Container definitions are used in task definitions to describe the different containers that are launched as part of a task.",
				Resolver:      schema.PathTableResolver("ContainerDefinitions"),
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "task_definition_cq_id",
						Description: "Unique CloudQuery ID of aws_ecs_task_definitions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "command",
						Description: "The command that is passed to the container",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "cpu",
						Description: "The number of cpu units reserved for the container",
						Type:        schema.TypeInt,
					},
					{
						Name:        "depends_on",
						Description: "The dependencies defined for container startup and shutdown",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsDependsOn,
					},
					{
						Name:        "disable_networking",
						Description: "When this parameter is true, networking is disabled within the container",
						Type:        schema.TypeBool,
					},
					{
						Name:        "dns_search_domains",
						Description: "A list of DNS search domains that are presented to the container",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "dns_servers",
						Description: "A list of DNS servers that are presented to the container",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "docker_labels",
						Description: "A key/value map of labels to add to the container",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "docker_security_options",
						Description: "A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "entry_point",
						Description: "Early versions of the Amazon ECS container agent do not properly handle entryPoint parameters",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "environment",
						Description: "The environment variables to pass to a container",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsEnvironment,
					},
					{
						Name:        "environment_files",
						Description: "A list of files containing the environment variables to pass to a container. This parameter maps to the --env-file option to docker run (https://docs.docker.com/engine/reference/run/#security-configuration)",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsEnvironmentFiles,
					},
					{
						Name:        "essential",
						Description: "If the essential parameter of a container is marked as true, and that container fails or stops for any reason, all other containers that are part of the task are stopped",
						Type:        schema.TypeBool,
					},
					{
						Name:        "extra_hosts",
						Description: "A list of hostnames and IP address mappings to append to the /etc/hosts file on the container",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsExtraHosts,
					},
					{
						Name:        "firelens_configuration_type",
						Description: "The log router to use",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FirelensConfiguration.Type"),
					},
					{
						Name:        "firelens_configuration_options",
						Description: "The options to use when configuring the log router",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("FirelensConfiguration.Options"),
					},
					{
						Name:        "health_check_command",
						Description: "A string array representing the command that the container runs to determine if it is healthy",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("HealthCheck.Command"),
					},
					{
						Name:        "health_check_interval",
						Description: "The time period in seconds between each health check execution",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("HealthCheck.Interval"),
					},
					{
						Name:        "health_check_retries",
						Description: "The number of times to retry a failed health check before the container is considered unhealthy",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("HealthCheck.Retries"),
					},
					{
						Name:        "health_check_start_period",
						Description: "The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("HealthCheck.StartPeriod"),
					},
					{
						Name:        "health_check_timeout",
						Description: "The time period in seconds to wait for a health check to succeed before it is considered a failure",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("HealthCheck.Timeout"),
					},
					{
						Name:        "hostname",
						Description: "The hostname to use for your container",
						Type:        schema.TypeString,
					},
					{
						Name:        "image",
						Description: "The image used to start a container",
						Type:        schema.TypeString,
					},
					{
						Name:        "interactive",
						Description: "When this parameter is true, this allows you to deploy containerized applications that require stdin or a tty to be allocated",
						Type:        schema.TypeBool,
					},
					{
						Name:        "links",
						Description: "The links parameter allows containers to communicate with each other without the need for port mappings",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "linux_parameters_capabilities_add",
						Description: "The Linux capabilities for the container that have been added to the default configuration provided by Docker",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("LinuxParameters.Capabilities.Add"),
					},
					{
						Name:        "linux_parameters_capabilities_drop",
						Description: "The Linux capabilities for the container that have been removed from the default configuration provided by Docker",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("LinuxParameters.Capabilities.Drop"),
					},
					{
						Name:        "linux_parameters_devices",
						Description: "Any host devices to expose to the container",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("LinuxParameters.Devices"),
					},
					{
						Name:        "linux_parameters_init_process_enabled",
						Description: "Run an init process inside the container that forwards signals and reaps processes",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("LinuxParameters.InitProcessEnabled"),
					},
					{
						Name:        "linux_parameters_max_swap",
						Description: "The total amount of swap memory (in MiB) a container can use",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LinuxParameters.MaxSwap"),
					},
					{
						Name:        "linux_parameters_shared_memory_size",
						Description: "The value for the size (in MiB) of the /dev/shm volume",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LinuxParameters.SharedMemorySize"),
					},
					{
						Name:        "linux_parameters_swappiness",
						Description: "This allows you to tune a container's memory swappiness behavior",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LinuxParameters.Swappiness"),
					},
					{
						Name:        "linux_parameters_tmpfs",
						Description: "The container path, mount options, and size (in MiB) of the tmpfs mount",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("LinuxParameters.Tmpfs"),
					},
					{
						Name:        "log_configuration_log_driver",
						Description: "The log driver to use for the container",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LogConfiguration.LogDriver"),
					},
					{
						Name:        "log_configuration_options",
						Description: "The configuration options to send to the log driver",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("LogConfiguration.Options"),
					},
					{
						Name:        "log_configuration_secret_options",
						Description: "The secrets to pass to the log configuration",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptions,
					},
					{
						Name:        "memory",
						Description: "The amount (in MiB) of memory to present to the container",
						Type:        schema.TypeInt,
					},
					{
						Name:        "memory_reservation",
						Description: "The soft limit (in MiB) of memory to reserve for the container",
						Type:        schema.TypeInt,
					},
					{
						Name:        "mount_points",
						Description: "The mount points for data volumes in your container",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("MountPoints"),
					},
					{
						Name:        "name",
						Description: "The name of a container",
						Type:        schema.TypeString,
					},
					{
						Name:        "port_mappings",
						Description: "The list of port mappings for the container",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("PortMappings"),
					},
					{
						Name:        "privileged",
						Description: "When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user)",
						Type:        schema.TypeBool,
					},
					{
						Name:        "pseudo_terminal",
						Description: "When this parameter is true, a TTY is allocated",
						Type:        schema.TypeBool,
					},
					{
						Name:        "readonly_root_filesystem",
						Description: "When this parameter is true, the container is given read-only access to its root file system",
						Type:        schema.TypeBool,
					},
					{
						Name:        "repository_credentials_parameter",
						Description: "The Amazon Resource Name (ARN) of the secret containing the private repository credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RepositoryCredentials.CredentialsParameter"),
					},
					{
						Name:        "resource_requirements",
						Description: "The type and amount of a resource to assign to a container",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsResourceRequirements,
					},
					{
						Name:        "secrets",
						Description: "The secrets to pass to the container",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsSecrets,
					},
					{
						Name:        "start_timeout",
						Description: "Time duration (in seconds) to wait before giving up on resolving dependencies for a container",
						Type:        schema.TypeInt,
					},
					{
						Name:        "stop_timeout",
						Description: "Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own",
						Type:        schema.TypeInt,
					},
					{
						Name:        "system_controls",
						Description: "A list of namespaced kernel parameters to set in the container",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsSystemControls,
					},
					{
						Name:        "ulimits",
						Description: "A list of ulimits to set in the container",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Ulimits"),
					},
					{
						Name:        "user",
						Description: "The user to use inside the container",
						Type:        schema.TypeString,
					},
					{
						Name:        "volumes_from",
						Description: "Data volumes to mount from another container",
						Type:        schema.TypeJSON,
						Resolver:    resolveEcsTaskDefinitionContainerDefinitionsVolumesFrom,
					},
					{
						Name:        "working_directory",
						Description: "The working directory in which to run commands inside the container",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ecs_task_definition_volumes",
				Description:   "A data volume used in a task definition",
				Resolver:      schema.PathTableResolver("Volumes"),
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "task_definition_cq_id",
						Description: "Unique CloudQuery ID of aws_ecs_task_definitions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "docker_autoprovision",
						Description: "If this value is true, the Docker volume is created if it does not already exist",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DockerVolumeConfiguration.Autoprovision"),
					},
					{
						Name:        "docker_driver",
						Description: "The Docker volume driver to use",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DockerVolumeConfiguration.Driver"),
					},
					{
						Name:        "docker_driver_opts",
						Description: "A map of Docker driver-specific options passed through",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("DockerVolumeConfiguration.DriverOpts"),
					},
					{
						Name:        "docker_labels",
						Description: "Custom metadata to add to your Docker volume",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("DockerVolumeConfiguration.Labels"),
					},
					{
						Name:        "docker_scope",
						Description: "The scope for the Docker volume that determines its lifecycle",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DockerVolumeConfiguration.Scope"),
					},
					{
						Name:        "efs_file_system_id",
						Description: "The Amazon EFS file system ID to use.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EfsVolumeConfiguration.FileSystemId"),
					},
					{
						Name:        "efs_authorization_config_access_point_id",
						Description: "The Amazon EFS access point ID to use",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EfsVolumeConfiguration.AuthorizationConfig.AccessPointId"),
					},
					{
						Name:        "efs_authorization_config_iam",
						Description: "Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EfsVolumeConfiguration.AuthorizationConfig.Iam"),
					},
					{
						Name:        "efs_root_directory",
						Description: "The directory within the Amazon EFS file system to mount as the root directory inside the host",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EfsVolumeConfiguration.RootDirectory"),
					},
					{
						Name:        "efs_volume_configuration_transit_encryption",
						Description: "Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EfsVolumeConfiguration.TransitEncryption"),
					},
					{
						Name:        "efs_transit_encryption_port",
						Description: "The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EfsVolumeConfiguration.TransitEncryptionPort"),
					},
					{
						Name:        "fsx_wfs_authorization_config_credentials_parameter",
						Description: "The authorization credential option to use",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FsxWindowsFileServerVolumeConfiguration.AuthorizationConfig.CredentialsParameter"),
					},
					{
						Name:        "fsx_wfs_authorization_config_domain",
						Description: "A fully qualified domain name hosted by an AWS Directory Service (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html) Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FsxWindowsFileServerVolumeConfiguration.AuthorizationConfig.Domain"),
					},
					{
						Name:        "fsx_wfs_file_system_id",
						Description: "The Amazon FSx for Windows File Server file system ID to use.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FsxWindowsFileServerVolumeConfiguration.FileSystemId"),
					},
					{
						Name:        "fsx_wfs_root_directory",
						Description: "The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FsxWindowsFileServerVolumeConfiguration.RootDirectory"),
					},
					{
						Name:        "host_source_path",
						Description: "When the host parameter is used, specify a sourcePath to declare the path on the host container instance that is presented to the container",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Host.SourcePath"),
					},
					{
						Name:        "name",
						Description: "The name of the volume",
						Type:        schema.TypeString,
					},
				},
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
	}, func(o *ecs.Options) {
		o.Region = c.Region
	})
	if err != nil {
		errorChan <- diag.WrapError(err)
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
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, j))
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
	return diag.WrapError(resource.Set(c.Name, j))
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
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsDependsOn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, p := range r.DependsOn {
		if p.ContainerName == nil {
			continue
		}
		j[*p.ContainerName] = string(p.Condition)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsEnvironment(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, p := range r.Environment {
		if p.Name == nil {
			continue
		}
		j[*p.Name] = aws.ToString(p.Value)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsEnvironmentFiles(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, p := range r.EnvironmentFiles {
		j[string(p.Type)] = aws.ToString(p.Value)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsExtraHosts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]interface{}{}
	for _, h := range r.ExtraHosts {
		if h.Hostname == nil {
			continue
		}
		j[*h.Hostname] = h.IpAddress
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]interface{}{}
	if r.LogConfiguration == nil {
		return nil
	}
	for _, s := range r.LogConfiguration.SecretOptions {
		j[*s.Name] = *s.ValueFrom
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsResourceRequirements(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, s := range r.ResourceRequirements {
		j[string(s.Type)] = aws.ToString(s.Value)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsSecrets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, s := range r.Secrets {
		if s.Name == nil {
			continue
		}
		j[*s.Name] = aws.ToString(s.ValueFrom)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsSystemControls(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, s := range r.SystemControls {
		if s.Namespace == nil {
			continue
		}
		j[*s.Namespace] = aws.ToString(s.Value)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveEcsTaskDefinitionContainerDefinitionsVolumesFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]interface{}{}
	for _, s := range r.VolumesFrom {
		if s.SourceContainer == nil {
			continue
		}
		j[*s.SourceContainer] = aws.ToBool(s.ReadOnly)
	}
	return diag.WrapError(resource.Set(c.Name, j))
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
	return diag.WrapError(resource.Set(c.Name, j))
}
