# Table: gcp_compute_target_http_proxies

This table shows data for GCP Compute Target HTTP Proxies.

https://cloud.google.com/compute/docs/reference/rest/v1/targetHttpProxies#TargetHttpProxy

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|fingerprint|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|proxy_bind|`bool`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|url_map|`utf8`|