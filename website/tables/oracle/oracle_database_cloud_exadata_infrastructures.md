# Table: oracle_database_cloud_exadata_infrastructures

This table shows data for Oracle Database Cloud Exadata Infrastructures.

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
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|shape|`utf8`|
|availability_domain|`utf8`|
|compute_count|`int64`|
|storage_count|`int64`|
|total_storage_size_in_g_bs|`int64`|
|available_storage_size_in_g_bs|`int64`|
|cpu_count|`int64`|
|max_cpu_count|`int64`|
|memory_size_in_g_bs|`int64`|
|max_memory_in_g_bs|`int64`|
|db_node_storage_size_in_g_bs|`int64`|
|max_db_node_storage_in_g_bs|`int64`|
|data_storage_size_in_t_bs|`float64`|
|max_data_storage_in_t_bs|`float64`|
|additional_storage_count|`int64`|
|activated_storage_count|`int64`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|maintenance_window|`json`|
|last_maintenance_run_id|`utf8`|
|next_maintenance_run_id|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|customer_contacts|`json`|
|storage_server_version|`utf8`|
|db_server_version|`utf8`|
|monthly_storage_server_version|`utf8`|
|monthly_db_server_version|`utf8`|