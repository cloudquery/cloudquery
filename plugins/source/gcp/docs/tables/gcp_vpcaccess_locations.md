# Table: gcp_vpcaccess_locations

https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/Shared.Types/ListLocationsResponse#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_vpcaccess_locations:
  - [gcp_vpcaccess_connectors](gcp_vpcaccess_connectors.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|location_id|String|
|display_name|String|
|labels|JSON|
|metadata|JSON|