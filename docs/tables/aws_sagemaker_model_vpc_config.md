
# Table: aws_sagemaker_model_vpc_config
Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|model_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_model table (FK)|
|security_group_ids|text[]|The VPC security group IDs, in the form sg-xxxxxxxx|
|subnets|text[]|The ID of the subnets in the VPC to which you want to connect your training job or model|
