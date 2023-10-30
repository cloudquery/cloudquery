# Table: atlas_database_federations

This table shows data for Atlas Database Federations.

The composite primary key for this table is (**group_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_provider_config|`json`|
|data_process_region|`json`|
|group_id (PK)|`utf8`|
|hostnames|`list<item: utf8, nullable>`|
|name (PK)|`utf8`|
|state|`utf8`|
|storage|`json`|