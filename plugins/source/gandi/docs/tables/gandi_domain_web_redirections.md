# Table: gandi_domain_web_redirections



The composite primary key for this table is (**fqdn**, **host**, **type**).

## Relations
This table depends on [gandi_domains](gandi_domains.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|fqdn (PK)|String|
|host (PK)|String|
|type (PK)|String|
|url|String|
|cert_status|String|
|cert_uuid|String|
|created_at|Timestamp|
|protocol|String|
|updated_at|Timestamp|