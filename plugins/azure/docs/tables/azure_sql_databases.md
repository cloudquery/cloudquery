
# Table: azure_sql_databases
Azure sql database
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|The name of the SKU, typically, a letter + Number code, eg P3|
|sku_tier|text|The tier or edition of the particular SKU, eg Basic, Premium|
|sku_size|text|Size of the particular SKU|
|sku_family|text|If the service has different generations of hardware, for the same SKU, then that can be captured here|
|sku_capacity|integer|Capacity of the particular SKU|
|kind|text|Kind of database This is metadata used for the Azure portal experience|
|managed_by|text|Resource that manages the database|
|create_mode|text|Specifies the mode of database creation.|
|collation|text|The collation of the database|
|max_size_bytes|bigint|The max size of the database expressed in bytes|
|sample_name|text|The name of the sample schema to apply when creating this database.|
|elastic_pool_id|text|The resource identifier of the elastic pool containing this database|
|source_database_id|text|The resource identifier of the source database associated with create operation of this database|
|status|text|The status of the database.|
|database_id|uuid|The ID of the database|
|creation_date_time|timestamp without time zone|The creation date of the database (ISO8601 format)|
|current_service_objective_name|text|The current service level objective name of the database|
|requested_service_objective_name|text|The requested service level objective name of the database|
|default_secondary_location|text|The default secondary region for this database|
|failover_group_id|text|Failover Group resource identifier that this database belongs to|
|restore_point_in_time|timestamp without time zone|Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database|
|source_database_deletion_date_time|timestamp without time zone|Specifies the time that the database was deleted|
|recovery_services_recovery_point_id|text|The resource identifier of the recovery point associated with create operation of this database|
|long_term_retention_backup_resource_id|text|The resource identifier of the long term retention backup associated with create operation of this database|
|recoverable_database_id|text|The resource identifier of the recoverable database associated with create operation of this database|
|restorable_dropped_database_id|text|The resource identifier of the restorable dropped database associated with create operation of this database|
|catalog_collation|text|Collation of the metadata catalog.|
|zone_redundant|boolean|Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones|
|license_type|text|The license type to apply for this database.|
|max_log_size_bytes|bigint|The max log size for this database|
|earliest_restore_date_time|timestamp without time zone|This records the earliest start date and time that restore is available for this database (ISO8601 format)|
|read_scale|text|The state of read-only routing If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.|
|high_availability_replica_count|integer|The number of secondary replicas associated with the database that are used to provide high availability|
|secondary_type|text|The secondary type of the database if it is a secondary.|
|current_sku_name|text|The name of the SKU, typically, a letter + Number code, eg P3|
|current_sku_tier|text|The tier or edition of the particular SKU, eg Basic, Premium|
|current_sku_size|text|Size of the particular SKU|
|current_sku_family|text|If the service has different generations of hardware, for the same SKU, then that can be captured here|
|current_sku_capacity|integer|Capacity of the particular SKU|
|auto_pause_delay|integer|Time in minutes after which database is automatically paused.|
|storage_account_type|text|The storage account type used to store backups for this database.|
|min_capacity|float|Minimal capacity that database will always have allocated, if not paused|
|paused_date_time|timestamp without time zone||
|resumed_date_time|timestamp without time zone||
|maintenance_configuration_id|text|Maintenance configuration id assigned to the database.|
|location|text|Resource location|
|tags|jsonb|Resource tags|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
