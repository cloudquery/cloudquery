# Table: aws_ecr_registry_policies

This table shows data for Amazon Elastic Container Registry (ECR) Registry Policies.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetRegistryPolicy.html

The composite primary key for this table is (**account_id**, **region**, **registry_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|registry_id (PK)|`utf8`|
|policy_text|`json`|
|result_metadata|`json`|