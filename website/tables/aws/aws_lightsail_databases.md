# Table: aws_lightsail_databases

This table shows data for Lightsail Databases.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabase.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_databases:
  - [aws_lightsail_database_events](aws_lightsail_database_events)
  - [aws_lightsail_database_log_events](aws_lightsail_database_log_events)
  - [aws_lightsail_database_parameters](aws_lightsail_database_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|backup_retention_enabled|`bool`|
|ca_certificate_identifier|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|engine|`utf8`|
|engine_version|`utf8`|
|hardware|`json`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|location|`json`|
|master_database_name|`utf8`|
|master_endpoint|`json`|
|master_username|`utf8`|
|name|`utf8`|
|parameter_apply_status|`utf8`|
|pending_maintenance_actions|`json`|
|pending_modified_values|`json`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|publicly_accessible|`bool`|
|relational_database_blueprint_id|`utf8`|
|relational_database_bundle_id|`utf8`|
|resource_type|`utf8`|
|secondary_availability_zone|`utf8`|
|state|`utf8`|
|support_code|`utf8`|