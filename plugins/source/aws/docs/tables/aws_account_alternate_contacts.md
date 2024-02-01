# Table: aws_account_alternate_contacts

This table shows data for Account Alternate Contacts.

https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **alternate_contact_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|alternate_contact_type|`utf8`|
|email_address|`utf8`|
|name|`utf8`|
|phone_number|`utf8`|
|title|`utf8`|