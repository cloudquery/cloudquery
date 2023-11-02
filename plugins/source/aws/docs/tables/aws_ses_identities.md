# Table: aws_ses_identities

This table shows data for Amazon Simple Email Service (SES) Identities.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailIdentity.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|identity_name|`utf8`|
|sending_enabled|`bool`|
|configuration_set_name|`utf8`|
|dkim_attributes|`json`|
|feedback_forwarding_status|`bool`|
|identity_type|`utf8`|
|mail_from_attributes|`json`|
|policies|`json`|
|verification_status|`utf8`|
|verified_for_sending_status|`bool`|