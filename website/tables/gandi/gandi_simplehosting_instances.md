# Table: gandi_simplehosting_instances

This table shows data for Gandi Simple Hosting Instances.

The primary key for this table is **id**.

## Relations

The following tables depend on gandi_simplehosting_instances:
  - [gandi_simplehosting_instance_vhosts](gandi_simplehosting_instance_vhosts)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|sharing_id|utf8|
|id (PK)|utf8|
|name|utf8|
|size|utf8|
|status|utf8|
|database|json|
|language|json|
|datacenter|json|