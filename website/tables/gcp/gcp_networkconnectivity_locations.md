# Table: gcp_networkconnectivity_locations

This table shows data for GCP NetworkConnectivity Locations.

https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/v1/projects.locations/list

The primary key for this table is **name**.

## Relations

The following tables depend on gcp_networkconnectivity_locations:
- [gcp_networkconnectivity_internal_ranges](gcp_networkconnectivity_internal_ranges)

## Columns

| Name              | Type          |
|-------------------| ------------- |
| _cq_id            |`uuid`|
| _cq_parent_id     |`uuid`|
| project_id        |`utf8`|
| name              |`utf8`|
| display_name      |`utf8`|
| labels            |`json`|
| location_id       |`utf8`|
| metadata          |`json`|
