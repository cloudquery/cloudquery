# Table: aws_glacier_vault_notifications

This table shows data for Glacier Vault Notifications.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-get.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **vault_arn**.
## Relations

This table depends on [aws_glacier_vaults](aws_glacier_vaults.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|vault_arn|`utf8`|
|events|`list<item: utf8, nullable>`|
|sns_topic|`utf8`|