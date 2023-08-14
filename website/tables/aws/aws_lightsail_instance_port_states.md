# Table: aws_lightsail_instance_port_states

This table shows data for Lightsail Instance Port States.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_instances](aws_lightsail_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|instance_arn|`utf8`|
|cidr_list_aliases|`list<item: utf8, nullable>`|
|cidrs|`list<item: utf8, nullable>`|
|from_port|`int64`|
|ipv6_cidrs|`list<item: utf8, nullable>`|
|protocol|`utf8`|
|state|`utf8`|
|to_port|`int64`|