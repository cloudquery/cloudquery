
# Table: aws_workspaces_directories
Describes a directory that is used with Amazon WorkSpaces.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the workspaces directory|
|alias|text|The directory alias.|
|customer_user_name|text|The user name for the service account.|
|id|text|The directory identifier.|
|name|text|The name of the directory.|
|type|text|The directory type.|
|dns_ip_addresses|text[]|The IP addresses of the DNS servers for the directory.|
|iam_role_id|text|The identifier of the IAM role|
|ip_group_ids|text[]|The identifiers of the IP access control groups associated with the directory.|
|registration_code|text|The registration code for the directory|
|change_compute_type|text|Specifies whether users can change the compute type (bundle) for their WorkSpace.|
|increase_volume_size|text|Specifies whether users can increase the volume size of the drives on their WorkSpace.|
|rebuild_workspace|text|Specifies whether users can rebuild the operating system of a WorkSpace to its original state.|
|restart_workspace|text|Specifies whether users can restart their WorkSpace.|
|switch_running_mode|text|Specifies whether users can switch the running mode of their WorkSpace.|
|state|text|The state of the directory's registration with Amazon WorkSpaces|
|subnet_ids|text[]|The identifiers of the subnets used with the directory.|
|tenancy|text|Specifies whether the directory is dedicated or shared|
|device_type_android|text|Indicates whether users can use Android and Android-compatible Chrome OS devices to access their WorkSpaces.|
|device_type_chrome_os|text|Indicates whether users can use Chromebooks to access their WorkSpaces.|
|device_type_ios|text|Indicates whether users can use iOS devices to access their WorkSpaces.|
|device_type_linux|text|Indicates whether users can use Linux clients to access their WorkSpaces.|
|device_type_osx|text|Indicates whether users can use macOS clients to access their WorkSpaces.|
|device_type_web|text|Indicates whether users can access their WorkSpaces through a web browser.|
|device_type_windows|text|Indicates whether users can use Windows clients to access their WorkSpaces.|
|device_type_zero_client|text|Indicates whether users can use zero client devices to access their WorkSpaces.|
|custom_security_group_id|text|The identifier of the default security group to apply to WorkSpaces when they are created|
|default_ou|text|The organizational unit (OU) in the directory for the WorkSpace machine accounts.|
|enable_internet_access|boolean|Specifies whether to automatically assign an Elastic public IP address to WorkSpaces in this directory by default|
|enable_maintenance_mode|boolean|Specifies whether maintenance mode is enabled for WorkSpaces|
|enable_work_docs|boolean|Specifies whether the directory is enabled for Amazon WorkDocs.|
|user_enabled_as_local_administrator|boolean|Specifies whether WorkSpace users are local administrators on their WorkSpaces.|
|workspace_security_group_id|text|The identifier of the security group that is assigned to new WorkSpaces.|
