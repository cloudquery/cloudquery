# Table: aws_eks_cluster_addons

This table shows data for Amazon Elastic Kubernetes Service (EKS) Cluster Addons.

https://docs.aws.amazon.com/eks/latest/APIReference/API_Addon.html

The composite primary key for this table is (**arn**, **cluster_arn**).

## Relations

This table depends on [aws_eks_clusters](aws_eks_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|cluster_arn (PK)|`utf8`|
|addon_arn|`utf8`|
|addon_name|`utf8`|
|addon_version|`utf8`|
|cluster_name|`utf8`|
|configuration_values|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|health|`json`|
|marketplace_information|`json`|
|modified_at|`timestamp[us, tz=UTC]`|
|owner|`utf8`|
|publisher|`utf8`|
|service_account_role_arn|`utf8`|
|status|`utf8`|
|tags|`json`|