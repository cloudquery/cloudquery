# Table: aws_ram_resource_share_associations

This table shows data for RAM Resource Share Associations.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareAssociation.html

The composite primary key for this table is (**associated_entity**, **resource_share_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|associated_entity (PK)|`utf8`|
|association_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|external|`bool`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|resource_share_arn (PK)|`utf8`|
|resource_share_name|`utf8`|
|status|`utf8`|
|status_message|`utf8`|