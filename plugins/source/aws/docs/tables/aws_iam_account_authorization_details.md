# Table: aws_iam_account_authorization_details

This table shows data for IAM Account Authorization Details.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountAuthorizationDetails.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|group_detail_list|`json`|
|policies|`json`|
|role_detail_list|`json`|
|user_detail_list|`json`|