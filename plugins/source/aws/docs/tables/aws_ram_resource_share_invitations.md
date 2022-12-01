# Table: aws_ram_resource_share_invitations

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|invitation_timestamp|Timestamp|
|receiver_account_id|String|
|receiver_arn|String|
|resource_share_arn|String|
|resource_share_associations|JSON|
|arn (PK)|String|
|resource_share_name|String|
|sender_account_id|String|
|status|String|