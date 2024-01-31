# Table: aws_ecr_registry_policies

This table shows data for Amazon Elastic Container Registry (ECR) Registry Policies.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetRegistryPolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **registry_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|policy_text|`json`|
|registry_id|`utf8`|