# Table: aws_iam_openid_connect_identity_providers

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetOpenIDConnectProvider.html

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
|tags|JSON|
|client_id_list|StringArray|
|create_date|Timestamp|
|thumbprint_list|StringArray|
|url|String|
|result_metadata|JSON|