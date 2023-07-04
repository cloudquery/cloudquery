# Table: gcp_compute_addresses

This table shows data for GCP Compute Addresses.

https://cloud.google.com/compute/docs/reference/rest/v1/addresses#Address

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|address|`utf8`|
|address_type|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|id|`int64`|
|ip_version|`utf8`|
|ipv6_endpoint_type|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|network|`utf8`|
|network_tier|`utf8`|
|prefix_length|`int64`|
|purpose|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|status|`utf8`|
|subnetwork|`utf8`|
|users|`list<item: utf8, nullable>`|