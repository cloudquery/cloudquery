
# Table: aws_ecs_task_definition_volumes
A data volume used in a task definition
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|task_definition_cq_id|uuid|Unique CloudQuery ID of aws_ecs_task_definitions table (FK)|
|docker_autoprovision|boolean|If this value is true, the Docker volume is created if it does not already exist|
|docker_driver|text|The Docker volume driver to use|
|docker_driver_opts|jsonb|A map of Docker driver-specific options passed through|
|docker_labels|jsonb|Custom metadata to add to your Docker volume|
|docker_scope|text|The scope for the Docker volume that determines its lifecycle|
|efs_file_system_id|text|The Amazon EFS file system ID to use.|
|efs_authorization_config_access_point_id|text|The Amazon EFS access point ID to use|
|efs_authorization_config_iam|text|Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system|
|efs_root_directory|text|The directory within the Amazon EFS file system to mount as the root directory inside the host|
|efs_volume_configuration_transit_encryption|text|Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server|
|efs_transit_encryption_port|integer|The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server|
|fsx_wfs_authorization_config_credentials_parameter|text|The authorization credential option to use|
|fsx_wfs_authorization_config_domain|text|A fully qualified domain name hosted by an AWS Directory Service (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html) Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.|
|fsx_wfs_file_system_id|text|The Amazon FSx for Windows File Server file system ID to use.|
|fsx_wfs_root_directory|text|The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.|
|host_source_path|text|When the host parameter is used, specify a sourcePath to declare the path on the host container instance that is presented to the container|
|name|text|The name of the volume|
