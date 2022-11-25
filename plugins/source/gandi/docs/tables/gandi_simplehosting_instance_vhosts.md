# Table: gandi_simplehosting_instance_vhosts



The composite primary key for this table is (**instance_id**, **fqdn**).

## Relations
This table depends on [gandi_simplehosting_instances](gandi_simplehosting_instances.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|instance_id (PK)|String|
|created_at|String|
|fqdn (PK)|String|
|is_a_test_vhost|Bool|
|linked_dns_zone|JSON|
|status|String|
|application|JSON|