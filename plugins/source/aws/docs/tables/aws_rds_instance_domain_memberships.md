
# Table: aws_rds_instance_domain_memberships
An Active Directory Domain membership record associated with the DB instance or cluster. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|instance_id|text|The AWS Region-unique, immutable identifier for the DB instance|
|domain|text|The identifier of the Active Directory Domain.|
|fqdn|text|The fully qualified domain name of the Active Directory Domain.|
|iam_role_name|text|The name of the IAM role to be used when making API calls to the Directory Service.|
|status|text|The status of the Active Directory Domain membership for the DB instance or cluster|
