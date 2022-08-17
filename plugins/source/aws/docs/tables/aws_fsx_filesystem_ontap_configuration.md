
# Table: aws_fsx_filesystem_ontap_configuration
Configuration for the FSx for NetApp ONTAP file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filesystem_cq_id|uuid|Unique CloudQuery ID of aws_fsx_filesystems table (FK)|
|automatic_backup_retention_days|bigint|The number of days to retain automatic backups|
|daily_automatic_backup_start_time|text|A recurring daily time, in the format HH:MM|
|deployment_type|text|Specifies the FSx for ONTAP file system deployment type in use in the file system.  * MULTI_AZ_1 - (Default) A high availability file system configured for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability.  * SINGLE_AZ_1 - A file system configured for Single-AZ redundancy.  For information about the use cases for Multi-AZ and Single-AZ deployments, refer to Choosing Multi-AZ or Single-AZ file system deployment (https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html).|
|disk_iops_configuration_iops|bigint|The total number of SSD IOPS provisioned for the file system.|
|disk_iops_configuration_mode|text|Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED).|
|endpoint_ip_address_range|text|(Multi-AZ only) The IP address range in which the endpoints to access your file system are created|
|endpoints_intercluster_dns_name|text|The Domain Name Service (DNS) name for the file system|
|endpoints_intercluster_ip_addresses|text[]|IP addresses of the file system endpoint.|
|endpoints_management_dns_name|text|The Domain Name Service (DNS) name for the file system|
|endpoints_management_ip_addresses|text[]|IP addresses of the file system endpoint.|
|preferred_subnet_id|text|The ID for a subnet|
|route_table_ids|text[]|(Multi-AZ only) The VPC route tables in which your file system's endpoints are created.|
|throughput_capacity|bigint|The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps).|
|weekly_maintenance_start_time|text|A recurring weekly time, in the format D:HH:MM|
