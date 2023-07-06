# Table: aws_glacier_vault_notifications

This table shows data for Glacier Vault Notifications.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-get.html

The primary key for this table is **vault_arn**.

## Relations

This table depends on [aws_glacier_vaults](aws_glacier_vaults).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|vault_arn (PK)|`utf8`|
|events|`list<item: utf8, nullable>`|
|sns_topic|`utf8`|