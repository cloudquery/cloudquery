# Table: aws_iot_policies

This table shows data for AWS IoT Policies.

https://docs.aws.amazon.com/iot/latest/apireference/API_Policy.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|policy_arn|`utf8`|
|policy_name|`utf8`|