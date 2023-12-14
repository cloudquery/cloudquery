# Table: gcp_compute_interconnect_remote_locations

This table shows data for GCP Compute Interconnect Remote Locations.

 https://cloud.google.com/compute/docs/reference/rest/v1/interconnectRemoteLocations#InterconnectRemoteLocationsGetResponse

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|address|`utf8`|
|attachment_configuration_constraints|`json`|
|city|`utf8`|
|constraints|`json`|
|continent|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|facility_provider|`utf8`|
|facility_provider_facility_id|`utf8`|
|id|`int64`|
|kind|`utf8`|
|lacp|`utf8`|
|max_lag_size100_gbps|`int64`|
|max_lag_size10_gbps|`int64`|
|name|`utf8`|
|peeringdb_facility_id|`utf8`|
|permitted_connections|`json`|
|remote_service|`utf8`|
|self_link (PK)|`utf8`|
|status|`utf8`|