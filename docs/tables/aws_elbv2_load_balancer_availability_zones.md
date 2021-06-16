
# Table: aws_elbv2_load_balancer_availability_zones
Information about an Availability Zone.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv2_load_balancers table (FK)|
|outpost_id|text|[Application Load Balancers on Outposts] The ID of the Outpost.|
|subnet_id|text|The ID of the subnet.|
|zone_name|text|The name of the Availability Zone.|
