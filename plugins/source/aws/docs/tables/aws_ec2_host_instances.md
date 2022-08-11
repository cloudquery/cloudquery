
# Table: aws_ec2_host_instances
Describes an instance running on a Dedicated Host.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|host_cq_id|uuid|Unique CloudQuery ID of aws_ec2_hosts table (FK)|
|instance_id|text|The ID of instance that is running on the Dedicated Host.|
|instance_type|text|The instance type (for example, m3.medium) of the running instance.|
|owner_id|text|The ID of the Amazon Web Services account that owns the instance.|
