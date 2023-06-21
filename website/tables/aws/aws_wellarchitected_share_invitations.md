# Table: aws_wellarchitected_share_invitations

This table shows data for Wellarchitected Share Invitations.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Workload.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|lens_arn|`utf8`|
|lens_name|`utf8`|
|permission_type|`utf8`|
|profile_arn|`utf8`|
|profile_name|`utf8`|
|id (PK)|`utf8`|
|share_resource_type|`utf8`|
|shared_by|`utf8`|
|shared_with|`utf8`|
|workload_id|`utf8`|
|workload_name|`utf8`|