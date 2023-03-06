# Table: aws_s3_access_points

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|bucket|String|
|name|String|
|network_origin|String|
|access_point_arn|String|
|alias|String|
|bucket_account_id|String|
|vpc_configuration|JSON|