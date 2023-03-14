# Table: aws_iam_saml_identity_providers

This table shows data for IAM Saml Identity Providers.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|create_date|Timestamp|
|saml_metadata_document|String|
|tags|JSON|
|valid_until|Timestamp|