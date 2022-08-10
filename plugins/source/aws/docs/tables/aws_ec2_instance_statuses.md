
# Table: aws_ec2_instance_statuses
Describes the status of an instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|availability_zone|text|The Availability Zone of the instance.|
|instance_id|text|The ID of the instance.|
|instance_state_code|integer|The state of the instance as a 16-bit unsigned integer.|
|instance_state_name|text|The current state of the instance.|
|details|jsonb|The system instance health or application instance health.|
|status|text|The instance status.|
|outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost.|
|system_status|text|The system status.|
|system_status_details|jsonb|The system instance health or application instance health.|
