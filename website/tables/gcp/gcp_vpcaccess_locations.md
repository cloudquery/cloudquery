# Table: gcp_vpcaccess_locations

This table shows data for GCP VPC Access Locations.

https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/Shared.Types/ListLocationsResponse#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_vpcaccess_locations:
  - [gcp_vpcaccess_connectors](gcp_vpcaccess_connectors)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|location_id|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|metadata|`json`|