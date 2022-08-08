
# Table: aws_rds_instance_associated_roles
Describes an AWS Identity and Access Management (IAM) role that is associated with a DB instance. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|instance_id|text|The AWS Region-unique, immutable identifier for the DB instance|
|feature_name|text|The name of the feature associated with the AWS Identity and Access Management (IAM) role|
|role_arn|text|The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.|
|status|text|Describes the state of association between the IAM role and the DB instance|
