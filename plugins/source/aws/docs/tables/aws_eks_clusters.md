# Table: aws_eks_clusters



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|certificate_authority|JSON|
|client_request_token|String|
|connector_config|JSON|
|created_at|Timestamp|
|encryption_config|JSON|
|endpoint|String|
|identity|JSON|
|kubernetes_network_config|JSON|
|logging|JSON|
|name|String|
|platform_version|String|
|resources_vpc_config|JSON|
|role_arn|String|
|status|String|
|tags|JSON|
|version|String|