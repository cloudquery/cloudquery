# Table: aws_eks_cluster_oidc_identity_provider_configs

This table shows data for Amazon Elastic Kubernetes Service (EKS) Cluster Oidc Identity Provider Configs.

https://docs.aws.amazon.com/eks/latest/APIReference/API_OidcIdentityProviderConfig.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**arn**, **cluster_arn**).
## Relations

This table depends on [aws_eks_clusters](aws_eks_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|cluster_arn|`utf8`|
|client_id|`utf8`|
|cluster_name|`utf8`|
|groups_claim|`utf8`|
|groups_prefix|`utf8`|
|identity_provider_config_arn|`utf8`|
|identity_provider_config_name|`utf8`|
|issuer_url|`utf8`|
|required_claims|`json`|
|status|`utf8`|
|tags|`json`|
|username_claim|`utf8`|
|username_prefix|`utf8`|