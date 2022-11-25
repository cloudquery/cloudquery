# Table: gandi_domain_dnssec_keys



The composite primary key for this table is (**fqdn**, **id**).

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
|algorithm|Int|
|digest|String|
|digest_type|Int|
|id (PK)|Int|
|keytag|Int|
|type|String|
|public_key|String|