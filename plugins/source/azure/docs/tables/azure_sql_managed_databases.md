
# Table: azure_sql_managed_databases
ManagedDatabase a managed database resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|collation|text|Collation of the managed database.|
|status|text|Status of the database|
|creation_date_time|timestamp without time zone||
|earliest_restore_point_time|timestamp without time zone||
|restore_point_in_time|timestamp without time zone||
|default_secondary_location|text|Geo paired region.|
|catalog_collation|text|Collation of the metadata catalog|
|create_mode|text|Managed database create mode|
|storage_container_uri|text|Conditional|
|source_database_id|text|The resource identifier of the source database associated with create operation of this database.|
|restorable_dropped_database_id|text|The restorable dropped database resource id to restore when creating this database.|
|storage_container_sas_token|text|SAS token used to access resources|
|failover_group_id|text|Instance Failover Group resource identifier that this managed database belongs to.|
|recoverable_database_id|text|The resource identifier of the recoverable database associated with create operation of this database.|
|long_term_retention_backup_resource_id|text|The name of the Long Term Retention backup to be used for restore of this managed database.|
|auto_complete_restore|boolean|Whether to auto complete restore of this managed database.|
|last_backup_name|text|Last backup file name for restore of this managed database.|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
|id|text|Resource ID.|
|name|text|Resource name.|
|type|text|Resource type.|
