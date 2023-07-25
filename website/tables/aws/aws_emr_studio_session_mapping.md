# Table: aws_emr_studio_session_mapping

This table shows data for Amazon EMR Studio Session Mapping.

https://docs.aws.amazon.com/emr/latest/APIReference/API_SessionMappingSummary.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_emr_studios](aws_emr_studios).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|creation_time|`timestamp[us, tz=UTC]`|
|identity_id|`utf8`|
|identity_name|`utf8`|
|identity_type|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|session_policy_arn|`utf8`|
|studio_id|`utf8`|