# Table: aws_ses_contact_lists

This table shows data for Amazon Simple Email Service (SES) Contact Lists.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|name|`utf8`|
|tags|`json`|
|contact_list_name|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_updated_timestamp|`timestamp[us, tz=UTC]`|
|topics|`json`|