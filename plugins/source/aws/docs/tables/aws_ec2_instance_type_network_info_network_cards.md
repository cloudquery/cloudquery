
# Table: aws_ec2_instance_type_network_info_network_cards
Describes the network card support of the instance type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_type_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instance_types table (FK)|
|maximum_network_interfaces|bigint|The maximum number of network interfaces for the network card.|
|network_card_index|bigint|The index of the network card.|
|network_performance|text|The network performance of the network card.|
