# Table: aws_emr_studio_session_mappings

This table shows data for Amazon EMR Studio Session Mappings.

https://docs.aws.amazon.com/emr/latest/APIReference/API_GetStudioSessionMapping.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**studio_arn**, **identity_id**, **identity_type**).
## Relations

This table depends on [aws_emr_studios](aws_emr_studios.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|studio_arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|identity_id|`utf8`|
|identity_name|`utf8`|
|identity_type|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|session_policy_arn|`utf8`|
|studio_id|`utf8`|