# Table: oracle_database_external_pluggable_databases

This table shows data for Oracle Database External Pluggable Databases.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|external_container_database_id|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|lifecycle_details|`utf8`|
|db_unique_name|`utf8`|
|db_id|`utf8`|
|database_version|`utf8`|
|database_edition|`utf8`|
|time_zone|`utf8`|
|character_set|`utf8`|
|ncharacter_set|`utf8`|
|db_packs|`utf8`|
|database_configuration|`utf8`|
|database_management_config|`json`|
|stack_monitoring_config|`json`|
|source_id|`utf8`|
|operations_insights_config|`json`|