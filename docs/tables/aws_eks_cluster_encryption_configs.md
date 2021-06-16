
# Table: aws_eks_cluster_encryption_configs
The encryption configuration for the cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid|Unique ID of aws_eks_clusters table (FK)|
|provider_key_arn|text|Amazon Resource Name (ARN) or alias of the customer master key (CMK).|
|resources|text[]|Specifies the resources to be encrypted.|
