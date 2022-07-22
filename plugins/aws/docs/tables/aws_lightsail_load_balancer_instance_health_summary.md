
# Table: aws_lightsail_load_balancer_instance_health_summary
Describes information about the health of the instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_load_balancers table (FK)|
|instance_health|text|Describes the overall instance health|
|instance_health_reason|text|More information about the instance health|
|instance_name|text|The name of the Lightsail instance for which you are requesting health check data|
