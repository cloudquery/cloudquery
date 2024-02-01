# Table: aws_eks_clusters

This table shows data for Amazon Elastic Kubernetes Service (EKS) Clusters.

https://docs.aws.amazon.com/eks/latest/APIReference/API_Cluster.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_eks_clusters:
  - [aws_eks_cluster_addons](aws_eks_cluster_addons.md)
  - [aws_eks_cluster_node_groups](aws_eks_cluster_node_groups.md)
  - [aws_eks_cluster_oidc_identity_provider_configs](aws_eks_cluster_oidc_identity_provider_configs.md)
  - [aws_eks_fargate_profiles](aws_eks_fargate_profiles.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|certificate_authority|`json`|
|client_request_token|`utf8`|
|connector_config|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|encryption_config|`json`|
|endpoint|`utf8`|
|health|`json`|
|id|`utf8`|
|identity|`json`|
|kubernetes_network_config|`json`|
|logging|`json`|
|name|`utf8`|
|outpost_config|`json`|
|platform_version|`utf8`|
|resources_vpc_config|`json`|
|role_arn|`utf8`|
|status|`utf8`|
|tags|`json`|
|version|`utf8`|