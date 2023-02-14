# Table: aws_ram_resource_share_invitations

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html

The composite primary key for this table is (**account_id**, **region**, **arn**, **receiver_combined**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|arn (PK)|String|
|receiver_combined (PK)|String|
|invitation_timestamp|Timestamp|
|receiver_account_id|String|
|receiver_arn|String|
|resource_share_arn|String|
|resource_share_associations|JSON|
|resource_share_invitation_arn|String|
|resource_share_name|String|
|sender_account_id|String|
|status|String|