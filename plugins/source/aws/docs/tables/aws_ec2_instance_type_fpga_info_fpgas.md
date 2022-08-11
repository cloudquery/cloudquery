
# Table: aws_ec2_instance_type_fpga_info_fpgas
Describes the FPGA accelerator for the instance type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_type_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instance_types table (FK)|
|count|bigint|The count of FPGA accelerators for the instance type.|
|manufacturer|text|The manufacturer of the FPGA accelerator.|
|memory_info_size_in_mi_b|bigint|The size of the memory available to the FPGA accelerator, in MiB.|
|name|text|The name of the FPGA accelerator.|
