# Table: aws_s3_access_grants

This table shows data for S3 Access Grants.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessGrantEntry.html

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
|access_grant_arn|`utf8`|
|access_grant_id|`utf8`|
|access_grants_location_configuration|`json`|
|access_grants_location_id|`utf8`|
|application_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|grant_scope|`utf8`|
|grantee|`json`|
|permission|`utf8`|