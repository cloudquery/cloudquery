# Table: aws_dms_replication_instances

This table shows data for Dms Replication Instances.

https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
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

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### AWS Database Migration Service replication instances should not be public

```sql
SELECT
  'AWS Database Migration Service replication instances should not be public'
    AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN publicly_accessible IS true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_dms_replication_instances;
```


