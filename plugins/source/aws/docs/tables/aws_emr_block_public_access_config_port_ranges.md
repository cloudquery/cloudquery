
# Table: aws_emr_block_public_access_config_port_ranges
A list of port ranges that are permitted to allow inbound traffic from all public IP addresses
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|block_public_access_config_cq_id|uuid|Unique CloudQuery ID of aws_emr_block_public_access_configs table (FK)|
|min_range|integer|The smallest port number in a specified range of port numbers.|
|max_range|integer|The smallest port number in a specified range of port numbers.|
