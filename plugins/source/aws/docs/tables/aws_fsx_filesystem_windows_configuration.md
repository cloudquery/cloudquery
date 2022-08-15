
# Table: aws_fsx_filesystem_windows_configuration
The configuration for this Microsoft Windows file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filesystem_cq_id|uuid|Unique CloudQuery ID of aws_fsx_filesystems table (FK)|
|active_directory_id|text|The ID for an existing Amazon Web Services Managed Microsoft Active Directory instance that the file system is joined to.|
|aliases|jsonb|An array of one or more DNS aliases that are currently associated with the Amazon FSx file system|
|audit_log_configuration_file_access_audit_log_level|text|Sets which attempt type is logged by Amazon FSx for file and folder accesses.  * SUCCESS_ONLY - only successful attempts to access files or folders are logged.  * FAILURE_ONLY - only failed attempts to access files or folders are logged.  * SUCCESS_AND_FAILURE - both successful attempts and failed attempts to access files or folders are logged.  * DISABLED - access auditing of files and folders is turned off.  This member is required.|
|audit_log_configuration_file_share_access_audit_log_level|text|Sets which attempt type is logged by Amazon FSx for file share accesses.  * SUCCESS_ONLY - only successful attempts to access file shares are logged.  * FAILURE_ONLY - only failed attempts to access file shares are logged.  * SUCCESS_AND_FAILURE - both successful attempts and failed attempts to access file shares are logged.  * DISABLED - access auditing of file shares is turned off.  This member is required.|
|audit_log_configuration_audit_log_destination|text|The Amazon Resource Name (ARN) for the destination of the audit logs|
|automatic_backup_retention_days|bigint|The number of days to retain automatic backups|
|copy_tags_to_backups|boolean|A boolean flag indicating whether tags on the file system should be copied to backups|
|daily_automatic_backup_start_time|text|The preferred time to take daily automatic backups, in the UTC time zone.|
|deployment_type|text|Specifies the file system deployment type, valid values are the following:  * MULTI_AZ_1 - Specifies a high availability file system that is configured for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability, and supports SSD and HDD storage.  * SINGLE_AZ_1 - (Default) Specifies a file system that is configured for single AZ redundancy, only supports SSD storage.  * SINGLE_AZ_2 - Latest generation Single AZ file system|
|maintenance_operations_in_progress|text[]|The list of maintenance operations in progress for this file system.|
|preferred_file_server_ip|text|For MULTI_AZ_1 deployment types, the IP address of the primary, or preferred, file server|
|preferred_subnet_id|text|For MULTI_AZ_1 deployment types, it specifies the ID of the subnet where the preferred file server is located|
|remote_administration_endpoint|text|For MULTI_AZ_1 deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell|
|self_managed_ad_config_dns_ips|text[]|A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.|
|self_managed_ad_config_domain_name|text|The fully qualified domain name of the self-managed AD directory.|
|self_managed_ad_config_file_system_administrators_group|text|The name of the domain group whose members have administrative privileges for the FSx file system.|
|self_managed_ad_config_organizational_unit_distinguished_name|text|The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined.|
|self_managed_ad_config_user_name|text|The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain.|
|throughput_capacity|bigint|The throughput of the Amazon FSx file system, measured in megabytes per second.|
|weekly_maintenance_start_time|text|The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone|
