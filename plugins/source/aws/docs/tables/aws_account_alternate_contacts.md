# Table: aws_account_alternate_contacts

https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html

The composite primary key for this table is (**account_id**, **alternate_contact_type**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|alternate_contact_type (PK)|String|
|email_address|String|
|name|String|
|phone_number|String|
|title|String|