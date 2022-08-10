
# Table: aws_elbv2_target_groups
Information about a target group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|health_check_enabled|boolean|Indicates whether health checks are enabled.|
|health_check_interval_seconds|integer|The approximate amount of time, in seconds, between health checks of an individual target.|
|health_check_path|text|The destination for health checks on the targets.|
|health_check_port|text|The port to use to connect with the target.|
|health_check_protocol|text|The protocol to use to connect with the target|
|health_check_timeout_seconds|integer|The amount of time, in seconds, during which no response means a failed health check.|
|healthy_threshold_count|integer|The number of consecutive health checks successes required before considering an unhealthy target healthy.|
|load_balancer_arns|text[]|The Amazon Resource Names (ARN) of the load balancers that route traffic to this target group.|
|matcher_grpc_code|text|You can specify values between 0 and 99|
|matcher_http_code|text|For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200|
|port|integer|The port on which the targets are listening|
|protocol|text|The protocol to use for routing traffic to the targets.|
|protocol_version|text|[HTTP/HTTPS protocol] The protocol version|
|arn|text|The Amazon Resource Name (ARN) of the target group.|
|name|text|The name of the target group.|
|target_type|text|The type of target that you must specify when registering targets with this target group|
|unhealthy_threshold_count|integer|The number of consecutive health check failures required before considering the target unhealthy.|
|vpc_id|text|The ID of the VPC for the targets.|
