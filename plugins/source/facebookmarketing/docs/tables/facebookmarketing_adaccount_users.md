# Table: facebookmarketing_adaccount_users

This table shows data for Facebook Marketing Adaccount Users.

https://developers.facebook.com/docs/marketing-api/reference/ad-account-user#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|tasks|`list<item: utf8, nullable>`|