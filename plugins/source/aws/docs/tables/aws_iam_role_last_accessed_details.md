# Table: aws_iam_role_last_accessed_details

This table shows data for IAM Role Last Accessed Details.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **role_arn**, **service_namespace**).
## Relations

This table depends on [aws_iam_roles](aws_iam_roles.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|role_arn|`utf8`|
|job_id|`utf8`|
|service_name|`utf8`|
|service_namespace|`utf8`|
|last_authenticated|`timestamp[us, tz=UTC]`|
|last_authenticated_entity|`utf8`|
|last_authenticated_region|`utf8`|
|total_authenticated_entities|`int64`|
|tracked_actions_last_accessed|`json`|