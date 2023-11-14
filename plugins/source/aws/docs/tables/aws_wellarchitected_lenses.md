# Table: aws_wellarchitected_lenses

This table shows data for AWS Well-Architected Lenses.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Lens.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|lens_alias|`utf8`|
|lens_arn|`utf8`|
|lens_name|`utf8`|
|lens_status|`utf8`|
|lens_type|`utf8`|
|lens_version|`utf8`|
|owner|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|share_invitation_id|`utf8`|
|tags|`json`|