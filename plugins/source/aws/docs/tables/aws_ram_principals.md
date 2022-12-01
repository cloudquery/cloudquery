# Table: aws_ram_principals

https://docs.aws.amazon.com/ram/latest/APIReference/API_Principal.html

The composite primary key for this table is (**account_id**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|creation_time|Timestamp|
|external|Bool|
|id (PK)|String|
|last_updated_time|Timestamp|
|resource_share_arn|String|