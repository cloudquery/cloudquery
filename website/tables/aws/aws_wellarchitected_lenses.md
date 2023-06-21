# Table: aws_wellarchitected_lenses

This table shows data for Wellarchitected Lenses.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Lens.html

The composite primary key for this table is (**account_id**, **region**, **alias**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|alias (PK)|`utf8`|
|arn|`utf8`|
|name|`utf8`|
|status|`utf8`|
|type|`utf8`|
|version|`utf8`|
|owner|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|share_invitation_id|`utf8`|
|tags|`json`|