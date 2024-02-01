# Table: aws_elasticache_users

This table shows data for Elasticache Users.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_User.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|access_string|`utf8`|
|authentication|`json`|
|engine|`utf8`|
|minimum_engine_version|`utf8`|
|status|`utf8`|
|user_group_ids|`list<item: utf8, nullable>`|
|user_id|`utf8`|
|user_name|`utf8`|