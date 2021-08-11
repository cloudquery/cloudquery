
# Table: aws_route53_health_checks
A complex type that contains information about one health check that is associated with the current AWS account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|cloud_watch_alarm_configuration_dimensions|jsonb|the metric that the CloudWatch alarm is associated with, a complex type that contains information about the dimensions for the metric.|
|tags|jsonb|The tags associated with the health check.|
|caller_reference|text|A unique string that you specified when you created the health check.|
|type|text|The type of health check that you want to create, which indicates how Amazon Route 53 determines whether an endpoint is healthy.|
|alarm_identifier_name|text|The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.|
|alarm_identifier_region|text|For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.|
|child_health_checks|text[]|(CALCULATED Health Checks Only) A complex type that contains one ChildHealthCheck element for each health check that you want to associate with a CALCULATED health check.|
|disabled|boolean|Stops Route 53 from performing health checks.|
|enable_sni|boolean|Specify whether you want Amazon Route 53 to send the value of FullyQualifiedDomainName to the endpoint in the client_hello message during TLS negotiation.|
|failure_threshold|integer|The number of consecutive health checks that an endpoint must pass or fail for Amazon Route 53 to change the current status of the endpoint from unhealthy to healthy or vice versa.|
|fully_qualified_domain_name|text|Amazon Route 53 behavior depends on whether you specify a value for IPAddress.|
|health_threshold|integer|The number of child health checks that are associated with a CALCULATED health check that Amazon Route 53 must consider healthy for the CALCULATED health check to be considered healthy.|
|ip_address|text|The IPv4 or IPv6 IP address of the endpoint that you want Amazon Route 53 to perform health checks on.|
|insufficient_data_health_status|text|When CloudWatch has insufficient data about the metric to determine the alarm state, the status that you want Amazon Route 53 to assign to the health check.|
|inverted|boolean|Specify whether you want Amazon Route 53 to invert the status of a health check, for example, to consider a health check unhealthy when it otherwise would be considered healthy.|
|measure_latency|boolean|Specify whether you want Amazon Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint, and to display CloudWatch latency graphs on the Health Checks page in the Route 53 console.|
|port|integer|The port on the endpoint that you want Amazon Route 53 to perform health checks on.|
|regions|text[]|A complex type that contains one Region element for each region from which you want Amazon Route 53 health checkers to check the specified endpoint.|
|request_interval|integer|The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health check request.|
|resource_path|text|The path, if any, that you want Amazon Route 53 to request when performing health checks.|
|search_string|text|If the value of Type is HTTP_STR_MATCH or HTTPS_STR_MATCH, the string that you want Amazon Route 53 to search for in the response body from the specified resource.|
|health_check_version|bigint|The version of the health check.|
|id|text|The identifier that Amazon Route 53 assigned to the health check when you created it.|
|cloud_watch_alarm_config_comparison_operator|text|For the metric that the CloudWatch alarm is associated with, the arithmetic operation that is used for the comparison.|
|cloud_watch_alarm_config_evaluation_periods|integer|For the metric that the CloudWatch alarm is associated with, the number of periods that the metric is compared to the threshold.|
|cloud_watch_alarm_config_metric_name|text|The name of the CloudWatch metric that the alarm is associated with.|
|cloud_watch_alarm_config_namespace|text|The namespace of the metric that the alarm is associated with.|
|cloud_watch_alarm_config_period|integer|For the metric that the CloudWatch alarm is associated with, the duration of one evaluation period in seconds.|
|cloud_watch_alarm_config_statistic|text|For the metric that the CloudWatch alarm is associated with, the statistic that is applied to the metric.|
|cloud_watch_alarm_config_threshold|float|For the metric that the CloudWatch alarm is associated with, the value the metric is compared with.|
|linked_service_description|text|If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.|
|linked_service_service_principal|text|If the health check or hosted zone was created by another service, the service that created the resource.|
|arn|text|The Amazon Resource Name (ARN) for the route 53 health check|
