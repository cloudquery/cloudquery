# Table: gcp_baremetalsolution_networks

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.networks#Network

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|type|String|
|ip_address|String|
|mac_address|StringArray|
|state|String|
|vlan_id|String|
|cidr|String|
|vrf|JSON|
|labels|JSON|
|services_cidr|String|
|reservations|JSON|