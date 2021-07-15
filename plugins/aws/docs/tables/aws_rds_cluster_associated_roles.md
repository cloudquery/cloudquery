
# Table: aws_rds_cluster_associated_roles
Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid|Unique CloudQuery ID of aws_rds_clusters table (FK)|
|feature_name|text|The name of the feature associated with the AWS Identity and Access Management (IAM) role|
|role_arn|text|The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.|
|status|text|Describes the state of association between the IAM role and the DB cluster|
