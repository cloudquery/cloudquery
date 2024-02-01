# Table: aws_iam_account_authorization_details

This table shows data for IAM Account Authorization Details.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountAuthorizationDetails.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|group_detail_list|`json`|
|policies|`json`|
|role_detail_list|`json`|
|user_detail_list|`json`|