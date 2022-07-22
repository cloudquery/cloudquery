
# Table: aws_lightsail_databases
Describes a database
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the database|
|backup_retention_enabled|boolean|A Boolean value indicating whether automated backup retention is enabled for the database|
|ca_certificate_identifier|text|The certificate associated with the database|
|created_at|timestamp without time zone|The timestamp when the database was created|
|engine|text|The database software (for example, MySQL)|
|engine_version|text|The database engine version (for example, 5723)|
|hardware_cpu_count|integer|The number of vCPUs for the database|
|hardware_disk_size_in_gb|integer|The size of the disk for the database|
|hardware_ram_size_in_gb|float|The amount of RAM in GB for the database|
|latest_restorable_time|timestamp without time zone|The latest point in time to which the database can be restored|
|availability_zone|text|The Availability Zone|
|master_database_name|text|The name of the master database created when the Lightsail database resource is created|
|master_endpoint_address|text|Specifies the DNS address of the database|
|master_endpoint_port|integer|Specifies the port that the database is listening on|
|master_username|text|The master user name of the database|
|name|text|The unique name of the database resource in Lightsail|
|parameter_apply_status|text|The status of parameter updates for the database|
|pending_modified_values_backup_retention_enabled|boolean|A Boolean value indicating whether automated backup retention is enabled|
|pending_modified_values_engine_version|text|The database engine version|
|pending_modified_values_master_user_password|text|The password for the master user of the database|
|preferred_backup_window|text|The daily time range during which automated backups are created for the database (for example, 16:00-16:30)|
|preferred_maintenance_window|text|The weekly time range during which system maintenance can occur on the database In the format ddd:hh24:mi-ddd:hh24:mi|
|publicly_accessible|boolean|A Boolean value indicating whether the database is publicly accessible|
|relational_database_blueprint_id|text|The blueprint ID for the database|
|relational_database_bundle_id|text|The bundle ID for the database|
|resource_type|text|The Lightsail resource type for the database (for example, RelationalDatabase)|
|secondary_availability_zone|text|Describes the secondary Availability Zone of a high availability database|
|state|text|Describes the current state of the database|
|support_code|text|The support code for the database|
|tags|jsonb|The tag keys and optional values for the resource|
