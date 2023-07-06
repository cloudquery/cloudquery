# Table: aws_ses_contact_lists

This table shows data for Amazon Simple Email Service (SES) Contact Lists.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|tags|`json`|
|contact_list_name|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_updated_timestamp|`timestamp[us, tz=UTC]`|
|topics|`json`|