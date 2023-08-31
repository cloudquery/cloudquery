# Table: gcp_vpcaccess_connectors

This table shows data for GCP VPC Access Connectors.

https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vpcaccess_locations](gcp_vpcaccess_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|network|`utf8`|
|ip_cidr_range|`utf8`|
|state|`utf8`|
|min_throughput|`int64`|
|max_throughput|`int64`|
|connected_projects|`list<item: utf8, nullable>`|
|subnet|`json`|
|machine_type|`utf8`|
|min_instances|`int64`|
|max_instances|`int64`|