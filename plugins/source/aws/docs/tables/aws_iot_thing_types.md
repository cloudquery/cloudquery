# Table: aws_iot_thing_types

This table shows data for AWS IoT Thing Types.

https://docs.aws.amazon.com/iot/latest/apireference/API_ThingTypeDefinition.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|thing_type_arn|`utf8`|
|thing_type_metadata|`json`|
|thing_type_name|`utf8`|
|thing_type_properties|`json`|