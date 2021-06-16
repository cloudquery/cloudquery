
# Table: aws_ec2_route_table_associations
Describes an association between a route table and a subnet or gateway.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_id|uuid|Unique ID of aws_ec2_route_tables table (FK)|
|association_state|text|The state of the association.|
|association_state_status_message|text|The status message, if applicable.|
|gateway_id|text|The ID of the internet gateway or virtual private gateway.|
|main|boolean|Indicates whether this is the main route table.|
|route_table_association_id|text|The ID of the association.|
|subnet_id|text|The ID of the subnet.|
