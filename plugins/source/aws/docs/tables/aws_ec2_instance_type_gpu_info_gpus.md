
# Table: aws_ec2_instance_type_gpu_info_gpus
Describes the GPU accelerators for the instance type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_type_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instance_types table (FK)|
|count|bigint|The number of GPUs for the instance type.|
|manufacturer|text|The manufacturer of the GPU accelerator.|
|memory_info_size_in_mi_b|bigint|The size of the memory available to the GPU accelerator, in MiB.|
|name|text|The name of the GPU accelerator.|
