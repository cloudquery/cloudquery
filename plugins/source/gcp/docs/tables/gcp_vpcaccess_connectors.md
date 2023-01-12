# Table: gcp_vpcaccess_connectors

https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vpcaccess_locations](gcp_vpcaccess_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|network|String|
|ip_cidr_range|String|
|state|String|
|min_throughput|Int|
|max_throughput|Int|
|connected_projects|StringArray|
|subnet|JSON|
|machine_type|String|
|min_instances|Int|
|max_instances|Int|