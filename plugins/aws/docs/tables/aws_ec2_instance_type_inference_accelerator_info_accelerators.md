
# Table: aws_ec2_instance_type_inference_accelerator_info_accelerators
Describes the Inference accelerators for the instance type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_type_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instance_types table (FK)|
|count|bigint|The number of Inference accelerators for the instance type.|
|manufacturer|text|The manufacturer of the Inference accelerator.|
|name|text|The name of the Inference accelerator.|
