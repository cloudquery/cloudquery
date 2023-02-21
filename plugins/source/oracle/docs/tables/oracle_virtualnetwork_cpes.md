# Table: oracle_virtualnetwork_cpes

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|ip_address|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|cpe_device_shape_id|String|
|time_created|Timestamp|