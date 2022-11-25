# Table: gandi_simplehosting_instances



The primary key for this table is **id**.

## Relations

The following tables depend on gandi_simplehosting_instances:
  - [gandi_simplehosting_instance_vhosts](gandi_simplehosting_instance_vhosts.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|sharing_id|String|
|id (PK)|String|
|name|String|
|size|String|
|status|String|
|database|JSON|
|language|JSON|
|datacenter|JSON|