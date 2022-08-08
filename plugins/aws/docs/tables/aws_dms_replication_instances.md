
# Table: aws_dms_replication_instances
Provides information that defines a replication instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|Any tags assigned to the resource|
|allocated_storage|integer|The amount of storage (in gigabytes) that is allocated for the replication instance.|
|auto_minor_version_upgrade|boolean|Boolean value indicating if minor version upgrades will be automatically applied to the instance.|
|availability_zone|text|The Availability Zone for the instance.|
|dns_name_servers|text|The DNS name servers supported for the replication instance to access your on-premise source or target database.|
|engine_version|text|The engine version number of the replication instance|
|free_until|timestamp without time zone|The expiration date of the free replication instance that is part of the Free DMS program.|
|instance_create_time|timestamp without time zone|The time the replication instance was created.|
|kms_key_id|text|An KMS key identifier that is used to encrypt the data on the replication instance|
|multi_az|boolean|Specifies whether the replication instance is a Multi-AZ deployment|
|pending_modified_values_allocated_storage|integer|The amount of storage (in gigabytes) that is allocated for the replication instance.|
|pending_modified_values_engine_version|text|The engine version number of the replication instance.|
|pending_modified_values_multi_az|boolean|Specifies whether the replication instance is a Multi-AZ deployment|
|pending_modified_values_class|text|The compute and memory capacity of the replication instance as defined for the specified replication instance class|
|preferred_maintenance_window|text|The maintenance window times for the replication instance|
|publicly_accessible|boolean|Specifies the accessibility options for the replication instance|
|arn|text|The Amazon Resource Name (ARN) of the replication instance.|
|class|text|The compute and memory capacity of the replication instance as defined for the specified replication instance class|
|identifier|text|The replication instance identifier is a required parameter|
|private_ip_address|inet|The private IP address of the replication instance.  Deprecated: This member has been deprecated.|
|private_ip_addresses|inet[]|One or more private IP addresses for the replication instance.|
|public_ip_address|inet|The public IP address of the replication instance.  Deprecated: This member has been deprecated.|
|public_ip_addresses|inet[]|One or more public IP addresses for the replication instance.|
|status|text|The status of the replication instance|
|replication_subnet_group_description|text|A description for the replication subnet group.|
|replication_subnet_group_identifier|text|The identifier of the replication instance subnet group.|
|replication_subnet_group_subnet_group_status|text|The status of the subnet group.|
|replication_subnet_group_vpc_id|text|The ID of the VPC.|
|secondary_availability_zone|text|The Availability Zone of the standby replication instance in a Multi-AZ deployment.|
