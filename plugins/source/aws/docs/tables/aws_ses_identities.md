# Table: aws_ses_identities

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetEmailIdentity.html

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
|arn (PK)|String|
|identity_name|String|
|sending_enabled|Bool|
|configuration_set_name|String|
|dkim_attributes|JSON|
|feedback_forwarding_status|Bool|
|identity_type|String|
|mail_from_attributes|JSON|
|policies|JSON|
|tags|JSON|
|verification_status|String|
|verified_for_sending_status|Bool|