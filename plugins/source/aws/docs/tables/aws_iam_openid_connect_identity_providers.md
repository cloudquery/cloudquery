# Table: aws_iam_openid_connect_identity_providers


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|client_id_list|StringArray|
|create_date|Timestamp|
|thumbprint_list|StringArray|
|url|String|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|