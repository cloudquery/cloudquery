# Table: aws_iot_things

This table shows data for AWS IoT Things.

https://docs.aws.amazon.com/iot/latest/apireference/API_ThingAttribute.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|principals|`list<item: utf8, nullable>`|
|arn (PK)|`utf8`|
|attributes|`json`|
|thing_arn|`utf8`|
|thing_name|`utf8`|
|thing_type_name|`utf8`|
|version|`int64`|