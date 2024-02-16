# Table: aws_xray_resource_policies

This table shows data for AWS X-Ray Resource Policies.

https://docs.aws.amazon.com/xray/latest/api/API_ResourcePolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **policy_name**, **policy_revision_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|policy_name|`utf8`|
|policy_revision_id|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|policy_document|`utf8`|