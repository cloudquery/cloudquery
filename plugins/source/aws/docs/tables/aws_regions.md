# Table: aws_regions

This table shows data for Regions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|enabled|`bool`|
|partition|`utf8`|
|region|`utf8`|
|endpoint|`utf8`|
|opt_in_status|`utf8`|
|region_name|`utf8`|