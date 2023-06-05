# Table: gandi_domain_web_redirections

This table shows data for Gandi Domain Web Redirections.

The composite primary key for this table is (**fqdn**, **host**, **type**).

## Relations

This table depends on [gandi_domains](gandi_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|fqdn (PK)|`utf8`|
|host (PK)|`utf8`|
|type (PK)|`utf8`|
|url|`utf8`|
|cert_status|`utf8`|
|cert_uuid|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|protocol|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|