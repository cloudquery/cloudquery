# Table: aws_account_contacts

This table shows data for Account Contacts.

https://docs.aws.amazon.com/accounts/latest/reference/API_ContactInformation.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|address_line1|`utf8`|
|city|`utf8`|
|country_code|`utf8`|
|full_name|`utf8`|
|phone_number|`utf8`|
|postal_code|`utf8`|
|address_line2|`utf8`|
|address_line3|`utf8`|
|company_name|`utf8`|
|district_or_county|`utf8`|
|state_or_region|`utf8`|
|website_url|`utf8`|