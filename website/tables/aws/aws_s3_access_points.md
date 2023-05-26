# Table: aws_s3_access_points

This table shows data for S3 Access Points.

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|bucket|utf8|
|name|utf8|
|network_origin|utf8|
|access_point_arn|utf8|
|alias|utf8|
|bucket_account_id|utf8|
|vpc_configuration|json|