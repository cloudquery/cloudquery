
# Table: aws_eks_cluster_loggings
An object representing the enabled or disabled Kubernetes control plane logs for your cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_eks_clusters table (FK)|
|enabled|boolean|If a log type is enabled, that log type exports its control plane logs to CloudWatch Logs.|
|types|text[]|The available cluster control plane log types.|
