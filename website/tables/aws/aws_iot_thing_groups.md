# Table: aws_iot_thing_groups

This table shows data for AWS IoT Thing Groups.

https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeThingGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|things_in_group|`list<item: utf8, nullable>`|
|policies|`list<item: utf8, nullable>`|
|tags|`json`|
|arn (PK)|`utf8`|
|index_name|`utf8`|
|query_string|`utf8`|
|query_version|`utf8`|
|status|`utf8`|
|thing_group_arn|`utf8`|
|thing_group_id|`utf8`|
|thing_group_metadata|`json`|
|thing_group_name|`utf8`|
|thing_group_properties|`json`|
|version|`int64`|
|result_metadata|`json`|