
# Table: aws_inspector2_finding_resources
Details about the resource involved in a finding
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|finding_cq_id|uuid|Unique CloudQuery ID of aws_inspector2_findings table (FK)|
|id|text|The ID of the resource|
|type|text|The type of resource|
|aws_ec2_instance|jsonb|An object that contains details about the Amazon EC2 instance involved in the finding|
|aws_ecr_container_image|jsonb|An object that contains details about the Amazon ECR container image involved in the finding|
|partition|text|The partition of the resource|
|region|text|The Amazon Web Services Region the impacted resource is located in|
|tags|jsonb|The tags attached to the resource|
