# Table: aws_lightsail_databases

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabase.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_lightsail_databases:
  - [aws_lightsail_database_parameters](aws_lightsail_database_parameters.md)
  - [aws_lightsail_database_events](aws_lightsail_database_events.md)
  - [aws_lightsail_database_log_events](aws_lightsail_database_log_events.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|backup_retention_enabled|Bool|
|ca_certificate_identifier|String|
|created_at|Timestamp|
|engine|String|
|engine_version|String|
|hardware|JSON|
|latest_restorable_time|Timestamp|
|location|JSON|
|master_database_name|String|
|master_endpoint|JSON|
|master_username|String|
|name|String|
|parameter_apply_status|String|
|pending_maintenance_actions|JSON|
|pending_modified_values|JSON|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|publicly_accessible|Bool|
|relational_database_blueprint_id|String|
|relational_database_bundle_id|String|
|resource_type|String|
|secondary_availability_zone|String|
|state|String|
|support_code|String|
|tags|JSON|