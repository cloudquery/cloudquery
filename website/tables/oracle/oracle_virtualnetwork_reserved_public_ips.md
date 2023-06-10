# Table: oracle_virtualnetwork_reserved_public_ips

This table shows data for Oracle Virtual Network Reserved Public IPs.

Reserved public IPs.
See https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/PublicIp/ListPublicIps for more details.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|