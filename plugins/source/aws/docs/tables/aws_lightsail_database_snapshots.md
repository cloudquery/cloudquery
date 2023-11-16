# Table: aws_lightsail_database_snapshots

This table shows data for Lightsail Database Snapshots.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseSnapshot.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|engine|`utf8`|
|engine_version|`utf8`|
|from_relational_database_arn|`utf8`|
|from_relational_database_blueprint_id|`utf8`|
|from_relational_database_bundle_id|`utf8`|
|from_relational_database_name|`utf8`|
|location|`json`|
|name|`utf8`|
|resource_type|`utf8`|
|size_in_gb|`int64`|
|state|`utf8`|
|support_code|`utf8`|