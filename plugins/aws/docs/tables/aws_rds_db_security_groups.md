
# Table: aws_rds_db_security_groups
Contains the details for an Amazon RDS DB security group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the DB security group.|
|description|text|Provides the description of the DB security group.|
|name|text|Specifies the name of the DB security group.|
|ec2_security_groups|jsonb|Contains a list of EC2 Security Group elements.|
|ip_ranges|jsonb|Contains a list of IP range elements.|
|owner_id|text|Provides the AWS ID of the owner of a specific DB security group.|
|vpc_id|text|Provides the VpcId of the DB security group.|
|tags|jsonb|List of tags|
