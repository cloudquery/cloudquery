# Table: gcp_compute_interconnect_locations

This table shows data for GCP Compute Interconnect Locations.

https://cloud.google.com/compute/docs/reference/rest/v1/interconnectLocations#InterconnectLocationsGetResponse

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|address|`utf8`|
|availability_zone|`utf8`|
|city|`utf8`|
|continent|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|facility_provider|`utf8`|
|facility_provider_facility_id|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|peeringdb_facility_id|`utf8`|
|region_infos|`json`|
|self_link (PK)|`utf8`|
|status|`utf8`|
|supports_pzs|`bool`|