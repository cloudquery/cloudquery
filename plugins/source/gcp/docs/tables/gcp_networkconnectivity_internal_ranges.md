# Table: gcp_networkconnectivity_internal_ranges

This table shows data for GCP Network Connectivity Internal Ranges.

https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/v1/projects.locations.internalRanges/list

The primary key for this table is **name**.

## Relations

This table depends on [gcp_networkconnectivity_locations](gcp_networkconnectivity_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|create_time|`utf8`|
|description|`utf8`|
|ip_cidr_range|`utf8`|
|labels|`json`|
|name (PK)|`utf8`|
|network|`utf8`|
|overlaps|`list<item: utf8, nullable>`|
|peering|`utf8`|
|prefix_length|`int64`|
|target_cidr_range|`list<item: utf8, nullable>`|
|update_time|`utf8`|
|usage|`utf8`|
|users|`list<item: utf8, nullable>`|