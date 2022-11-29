# Table: heroku_peerings

https://devcenter.heroku.com/articles/platform-api-reference#peering

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|aws_account_id|String|
|aws_region|String|
|aws_vpc_id|String|
|cidr_blocks|StringArray|
|expires|Timestamp|
|pcx_id|String|
|status|String|
|type|String|