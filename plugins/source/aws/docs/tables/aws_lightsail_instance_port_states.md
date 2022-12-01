# Table: aws_lightsail_instance_port_states

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lightsail_instances](aws_lightsail_instances.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|instance_arn|String|
|cidr_list_aliases|StringArray|
|cidrs|StringArray|
|from_port|Int|
|ipv6_cidrs|StringArray|
|protocol|String|
|state|String|
|to_port|Int|