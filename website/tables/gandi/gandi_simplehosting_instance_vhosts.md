# Table: gandi_simplehosting_instance_vhosts

This table shows data for Gandi Simple Hosting Instance Vhosts.

The composite primary key for this table is (**instance_id**, **fqdn**).

## Relations

This table depends on [gandi_simplehosting_instances](gandi_simplehosting_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|instance_id (PK)|`utf8`|
|created_at|`utf8`|
|fqdn (PK)|`utf8`|
|is_a_test_vhost|`bool`|
|linked_dns_zone|`json`|
|status|`utf8`|
|application|`json`|