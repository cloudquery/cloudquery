# Table: aws_emr_studio_session_mapping

This table shows data for Amazon EMR Studio Session Mapping.

https://docs.aws.amazon.com/emr/latest/APIReference/API_GetStudioSessionMapping.html

The composite primary key for this table is (**account_id**, **region**, **identity_id**, **identity_type**, **studio_id**).

## Relations

This table depends on [aws_emr_studios](aws_emr_studios).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|identity_id (PK)|`utf8`|
|identity_name|`utf8`|
|identity_type (PK)|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|session_policy_arn|`utf8`|
|studio_id (PK)|`utf8`|