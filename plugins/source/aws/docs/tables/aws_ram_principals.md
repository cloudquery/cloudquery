# Table: aws_ram_principals

This table shows data for RAM Principals.

https://docs.aws.amazon.com/ram/latest/APIReference/API_Principal.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**, **resource_share_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|external|`bool`|
|id|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|resource_share_arn|`utf8`|