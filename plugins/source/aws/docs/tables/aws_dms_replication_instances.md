# Table: aws_dms_replication_instances

This table shows data for Dms Replication Instances.

https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|allocated_storage|`int64`|
|auto_minor_version_upgrade|`bool`|
|availability_zone|`utf8`|
|dns_name_servers|`utf8`|
|engine_version|`utf8`|
|free_until|`timestamp[us, tz=UTC]`|
|instance_create_time|`timestamp[us, tz=UTC]`|
|kms_key_id|`utf8`|
|multi_az|`bool`|
|network_type|`utf8`|
|pending_modified_values|`json`|
|preferred_maintenance_window|`utf8`|
|publicly_accessible|`bool`|
|replication_instance_arn|`utf8`|
|replication_instance_class|`utf8`|
|replication_instance_identifier|`utf8`|
|replication_instance_ipv6_addresses|`list<item: utf8, nullable>`|
|replication_instance_private_ip_address|`utf8`|
|replication_instance_private_ip_addresses|`list<item: utf8, nullable>`|
|replication_instance_public_ip_address|`utf8`|
|replication_instance_public_ip_addresses|`list<item: utf8, nullable>`|
|replication_instance_status|`utf8`|
|replication_subnet_group|`json`|
|secondary_availability_zone|`utf8`|
|vpc_security_groups|`json`|
|tags|`json`|