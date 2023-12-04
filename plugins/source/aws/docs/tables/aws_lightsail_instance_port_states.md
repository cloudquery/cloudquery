# Table: aws_lightsail_instance_port_states

This table shows data for Lightsail Instance Port States.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html

The composite primary key for this table is (**instance_arn**, **allow_list**, **from_port**, **protocol**, **to_port**).

## Relations

This table depends on [aws_lightsail_instances](aws_lightsail_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|instance_arn (PK)|`utf8`|
|allow_list (PK)|`utf8`|
|cidr_list_aliases|`list<item: utf8, nullable>`|
|cidrs|`list<item: utf8, nullable>`|
|from_port (PK)|`int64`|
|ipv6_cidrs|`list<item: utf8, nullable>`|
|protocol (PK)|`utf8`|
|state|`utf8`|
|to_port (PK)|`int64`|