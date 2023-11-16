# Table: gcp_compute_url_maps

This table shows data for GCP Compute URL Maps.

https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps#UrlMap

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|default_route_action|`json`|
|default_service|`utf8`|
|default_url_redirect|`json`|
|description|`utf8`|
|fingerprint|`utf8`|
|header_action|`json`|
|host_rules|`json`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|path_matchers|`json`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|tests|`json`|