
# Table: aws_ec2_route_table_associations
Describes an association between a route table and a subnet or gateway.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_cq_id|uuid|Unique CloudQuery ID of aws_ec2_route_tables table (FK)|
|id|text|The ID of the association.|
|association_state|text|The state of the association.|
|association_state_status_message|text|The status message, if applicable.|
|gateway_id|text|The ID of the internet gateway or virtual private gateway.|
|main|boolean|Indicates whether this is the main route table.|
|subnet_id|text|The ID of the subnet.|
