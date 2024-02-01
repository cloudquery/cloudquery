# Table: aws_fsx_snapshots

This table shows data for Amazon FSx Snapshots.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Snapshot.html

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
|administrative_actions|`json`|
|tags|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|lifecycle|`utf8`|
|lifecycle_transition_reason|`json`|
|name|`utf8`|
|resource_arn|`utf8`|
|snapshot_id|`utf8`|
|volume_id|`utf8`|