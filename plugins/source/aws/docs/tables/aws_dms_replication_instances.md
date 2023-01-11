# Table: aws_dms_replication_instances

https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|allocated_storage|Int|
|auto_minor_version_upgrade|Bool|
|availability_zone|String|
|dns_name_servers|String|
|engine_version|String|
|free_until|Timestamp|
|instance_create_time|Timestamp|
|kms_key_id|String|
|multi_az|Bool|
|network_type|String|
|pending_modified_values|JSON|
|preferred_maintenance_window|String|
|publicly_accessible|Bool|
|replication_instance_arn|String|
|replication_instance_class|String|
|replication_instance_identifier|String|
|replication_instance_ipv6_addresses|StringArray|
|replication_instance_private_ip_address|String|
|replication_instance_private_ip_addresses|StringArray|
|replication_instance_public_ip_address|String|
|replication_instance_public_ip_addresses|StringArray|
|replication_instance_status|String|
|replication_subnet_group|JSON|
|secondary_availability_zone|String|
|vpc_security_groups|JSON|
|tags|JSON|