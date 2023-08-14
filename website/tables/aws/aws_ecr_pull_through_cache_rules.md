# Table: aws_ecr_pull_through_cache_rules

This table shows data for Amazon Elastic Container Registry (ECR) Pull Through Cache Rules.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_DescribePullThroughCacheRules.html

The composite primary key for this table is (**account_id**, **region**, **ecr_repository_prefix**, **registry_id**, **upstream_registry_url**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|ecr_repository_prefix (PK)|`utf8`|
|registry_id (PK)|`utf8`|
|upstream_registry_url (PK)|`utf8`|