# Table: gcp_compute_target_ssl_proxies

This table shows data for GCP Compute Target SSL Proxies.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|certificate_map|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|proxy_header|`utf8`|
|self_link (PK)|`utf8`|
|service|`utf8`|
|ssl_certificates|`list<item: utf8, nullable>`|
|ssl_policy|`utf8`|