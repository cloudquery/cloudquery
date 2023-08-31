# Table: aws_ram_resource_share_invitations

This table shows data for RAM Resource Share Invitations.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html

The composite primary key for this table is (**account_id**, **region**, **arn**, **receiver_combined**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn (PK)|`utf8`|
|receiver_combined (PK)|`utf8`|
|invitation_timestamp|`timestamp[us, tz=UTC]`|
|receiver_account_id|`utf8`|
|receiver_arn|`utf8`|
|resource_share_arn|`utf8`|
|resource_share_associations|`json`|
|resource_share_invitation_arn|`utf8`|
|resource_share_name|`utf8`|
|sender_account_id|`utf8`|
|status|`utf8`|