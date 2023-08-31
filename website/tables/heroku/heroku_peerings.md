# Table: heroku_peerings

This table shows data for Heroku Peerings.

https://devcenter.heroku.com/articles/platform-api-reference#peering

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|aws_account_id|`utf8`|
|aws_region|`utf8`|
|aws_vpc_id|`utf8`|
|cidr_blocks|`list<item: utf8, nullable>`|
|expires|`timestamp[us, tz=UTC]`|
|pcx_id|`utf8`|
|status|`utf8`|
|type|`utf8`|