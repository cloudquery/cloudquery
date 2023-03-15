# Table: aws_ecr_registry_policies

This table shows data for Amazon Elastic Container Registry (ECR) Registry Policies.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetRegistryPolicy.html

The composite primary key for this table is (**account_id**, **region**, **registry_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|registry_id (PK)|String|
|policy_text|JSON|
|result_metadata|JSON|