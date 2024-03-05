# Table: aws_s3_access_grant_instances

This table shows data for S3 Access Grant Instances.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessGrantsInstanceEntry.html

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
|access_grants_instance_arn|`utf8`|
|access_grants_instance_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|identity_center_arn|`utf8`|