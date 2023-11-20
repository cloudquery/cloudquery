# Table: aws_regions

This table shows data for Regions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|enabled|`bool`|
|partition|`utf8`|
|region (PK)|`utf8`|
|endpoint|`utf8`|
|opt_in_status|`utf8`|
|region_name|`utf8`|