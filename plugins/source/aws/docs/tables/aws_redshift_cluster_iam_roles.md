
# Table: aws_redshift_cluster_iam_roles
An AWS Identity and Access Management (IAM) role that can be used by the associated Amazon Redshift cluster to access other AWS services.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_redshift_clusters table (FK)|
|apply_status|text|A value that describes the status of the IAM role's association with an Amazon Redshift cluster.|
|iam_role_arn|text|The Amazon Resource Name (ARN) of the IAM role, for example, arn:aws:iam::123456789012:role/RedshiftCopyUnload.|
