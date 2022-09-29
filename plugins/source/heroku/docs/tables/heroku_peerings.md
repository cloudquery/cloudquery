# Table: heroku_peerings
https://devcenter.heroku.com/articles/platform-api-reference#peering-attributes

The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|aws_account_id|String|
|aws_region|String|
|aws_vpc_id|String|
|cidr_blocks|StringArray|
|expires|Timestamp|
|pcx_id|String|
|status|String|
|type|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|