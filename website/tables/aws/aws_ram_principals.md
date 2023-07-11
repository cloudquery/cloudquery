# Table: aws_ram_principals

This table shows data for RAM Principals.

https://docs.aws.amazon.com/ram/latest/APIReference/API_Principal.html

The composite primary key for this table is (**account_id**, **region**, **id**, **resource_share_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|external|`bool`|
|id (PK)|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|resource_share_arn (PK)|`utf8`|