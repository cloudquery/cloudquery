# Table: gcp_baremetalsolution_networks

This table shows data for GCP Bare Metal Solution Networks.

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.networks#Network

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|type|`utf8`|
|ip_address|`utf8`|
|mac_address|`list<item: utf8, nullable>`|
|state|`utf8`|
|vlan_id|`utf8`|
|cidr|`utf8`|
|vrf|`json`|
|labels|`json`|
|services_cidr|`utf8`|
|reservations|`json`|