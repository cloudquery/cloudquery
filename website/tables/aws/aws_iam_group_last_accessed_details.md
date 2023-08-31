# Table: aws_iam_group_last_accessed_details

This table shows data for IAM Group Last Accessed Details.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html

The composite primary key for this table is (**account_id**, **group_arn**, **service_namespace**).

## Relations

This table depends on [aws_iam_groups](aws_iam_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|group_arn (PK)|`utf8`|
|job_id|`utf8`|
|service_name|`utf8`|
|service_namespace (PK)|`utf8`|
|last_authenticated|`timestamp[us, tz=UTC]`|
|last_authenticated_entity|`utf8`|
|last_authenticated_region|`utf8`|
|total_authenticated_entities|`int64`|
|tracked_actions_last_accessed|`json`|