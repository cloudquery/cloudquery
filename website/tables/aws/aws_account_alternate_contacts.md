# Table: aws_account_alternate_contacts

This table shows data for Account Alternate Contacts.

https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html

The composite primary key for this table is (**account_id**, **alternate_contact_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|alternate_contact_type (PK)|`utf8`|
|email_address|`utf8`|
|name|`utf8`|
|phone_number|`utf8`|
|title|`utf8`|