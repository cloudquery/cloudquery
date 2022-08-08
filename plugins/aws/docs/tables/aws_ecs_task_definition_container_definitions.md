
# Table: aws_ecs_task_definition_container_definitions
Container definitions are used in task definitions to describe the different containers that are launched as part of a task.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|task_definition_cq_id|uuid|Unique CloudQuery ID of aws_ecs_task_definitions table (FK)|
|command|text[]|The command that is passed to the container|
|cpu|integer|The number of cpu units reserved for the container|
|depends_on|jsonb|The dependencies defined for container startup and shutdown|
|disable_networking|boolean|When this parameter is true, networking is disabled within the container|
|dns_search_domains|text[]|A list of DNS search domains that are presented to the container|
|dns_servers|text[]|A list of DNS servers that are presented to the container|
|docker_labels|jsonb|A key/value map of labels to add to the container|
|docker_security_options|text[]|A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems|
|entry_point|text[]|Early versions of the Amazon ECS container agent do not properly handle entryPoint parameters|
|environment|jsonb|The environment variables to pass to a container|
|environment_files|jsonb|A list of files containing the environment variables to pass to a container. This parameter maps to the --env-file option to docker run (https://docs.docker.com/engine/reference/run/#security-configuration)|
|essential|boolean|If the essential parameter of a container is marked as true, and that container fails or stops for any reason, all other containers that are part of the task are stopped|
|extra_hosts|jsonb|A list of hostnames and IP address mappings to append to the /etc/hosts file on the container|
|firelens_configuration_type|text|The log router to use|
|firelens_configuration_options|jsonb|The options to use when configuring the log router|
|health_check_command|text[]|A string array representing the command that the container runs to determine if it is healthy|
|health_check_interval|integer|The time period in seconds between each health check execution|
|health_check_retries|integer|The number of times to retry a failed health check before the container is considered unhealthy|
|health_check_start_period|integer|The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries|
|health_check_timeout|integer|The time period in seconds to wait for a health check to succeed before it is considered a failure|
|hostname|text|The hostname to use for your container|
|image|text|The image used to start a container|
|interactive|boolean|When this parameter is true, this allows you to deploy containerized applications that require stdin or a tty to be allocated|
|links|text[]|The links parameter allows containers to communicate with each other without the need for port mappings|
|linux_parameters_capabilities_add|text[]|The Linux capabilities for the container that have been added to the default configuration provided by Docker|
|linux_parameters_capabilities_drop|text[]|The Linux capabilities for the container that have been removed from the default configuration provided by Docker|
|linux_parameters_devices|jsonb|Any host devices to expose to the container|
|linux_parameters_init_process_enabled|boolean|Run an init process inside the container that forwards signals and reaps processes|
|linux_parameters_max_swap|integer|The total amount of swap memory (in MiB) a container can use|
|linux_parameters_shared_memory_size|integer|The value for the size (in MiB) of the /dev/shm volume|
|linux_parameters_swappiness|integer|This allows you to tune a container's memory swappiness behavior|
|linux_parameters_tmpfs|jsonb|The container path, mount options, and size (in MiB) of the tmpfs mount|
|log_configuration_log_driver|text|The log driver to use for the container|
|log_configuration_options|jsonb|The configuration options to send to the log driver|
|log_configuration_secret_options|jsonb|The secrets to pass to the log configuration|
|memory|integer|The amount (in MiB) of memory to present to the container|
|memory_reservation|integer|The soft limit (in MiB) of memory to reserve for the container|
|mount_points|jsonb|The mount points for data volumes in your container|
|name|text|The name of a container|
|port_mappings|jsonb|The list of port mappings for the container|
|privileged|boolean|When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user)|
|pseudo_terminal|boolean|When this parameter is true, a TTY is allocated|
|readonly_root_filesystem|boolean|When this parameter is true, the container is given read-only access to its root file system|
|repository_credentials_parameter|text|The Amazon Resource Name (ARN) of the secret containing the private repository credentials|
|resource_requirements|jsonb|The type and amount of a resource to assign to a container|
|secrets|jsonb|The secrets to pass to the container|
|start_timeout|integer|Time duration (in seconds) to wait before giving up on resolving dependencies for a container|
|stop_timeout|integer|Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own|
|system_controls|jsonb|A list of namespaced kernel parameters to set in the container|
|ulimits|jsonb|A list of ulimits to set in the container|
|user|text|The user to use inside the container|
|volumes_from|jsonb|Data volumes to mount from another container|
|working_directory|text|The working directory in which to run commands inside the container|
