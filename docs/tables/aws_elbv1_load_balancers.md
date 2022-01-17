
# Table: aws_elbv1_load_balancers
Information about a load balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|attributes_access_log_enabled|boolean||
|attributes_access_log_s3_bucket_name|text||
|attributes_access_log_s3_bucket_prefix|text||
|attributes_access_log_emit_interval|integer||
|attributes_connection_settings_idle_timeout|integer||
|attributes_cross_zone_load_balancing_enabled|boolean||
|attributes_connection_draining_enabled|boolean||
|attributes_connection_draining_timeout|integer||
|attributes_additional_attributes|jsonb||
|tags|jsonb||
|availability_zones|text[]|The Availability Zones for the load balancer.|
|canonical_hosted_zone_name|text|The DNS name of the load balancer.|
|canonical_hosted_zone_name_id|text|The ID of the Amazon Route 53 hosted zone for the load balancer.|
|created_time|timestamp without time zone|The date and time the load balancer was created.|
|dns_name|text|The DNS name of the load balancer.|
|health_check_healthy_threshold|integer|The number of consecutive health checks successes required before moving the instance to the Healthy state.|
|health_check_interval|integer|The approximate interval, in seconds, between health checks of an individual instance.|
|health_check_target|text|The instance being checked.|
|health_check_timeout|integer|The amount of time, in seconds, during which no response means a failed health check.|
|health_check_unhealthy_threshold|integer|The number of consecutive health check failures required before moving the instance to the Unhealthy state.|
|instances|text[]|The IDs of the instances for the load balancer.|
|name|text|The name of the load balancer.|
|other_policies|text[]|The policies other than the stickiness policies.|
|scheme|text|The type of load balancer.|
|security_groups|text[]|The security groups for the load balancer.|
|source_security_group_name|text|The name of the security group.|
|source_security_group_owner_alias|text|The owner of the security group.|
|subnets|text[]|The IDs of the subnets for the load balancer.|
|vpc_id|text|The ID of the VPC for the load balancer.|
