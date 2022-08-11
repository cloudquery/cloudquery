
# Table: aws_ec2_host_available_instance_capacity
Information about the number of instances that can be launched onto the Dedicated Host.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|host_cq_id|uuid|Unique CloudQuery ID of aws_ec2_hosts table (FK)|
|available_capacity|integer|The number of instances that can be launched onto the Dedicated Host based on the host's available capacity.|
|instance_type|text|The instance type supported by the Dedicated Host.|
|total_capacity|integer|The total number of instances that can be launched onto the Dedicated Host if there are no instances running on it.|
