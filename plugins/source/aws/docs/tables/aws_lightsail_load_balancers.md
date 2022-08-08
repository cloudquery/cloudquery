
# Table: aws_lightsail_load_balancers
Describes a load balancer
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the load balancer|
|configuration_options|jsonb|A string to string map of the configuration options for your load balancer Valid values are listed below|
|created_at|timestamp without time zone|The date when your load balancer was created|
|dns_name|text|The DNS name of your Lightsail load balancer|
|health_check_path|text|The path you specified to perform your health checks|
|https_redirection_enabled|boolean|A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer|
|instance_port|bigint|The port where the load balancer will direct traffic to your Lightsail instances|
|ip_address_type|text|The IP address type of the load balancer|
|availability_zone|text|The Availability Zone|
|name|text|The name of the load balancer (eg, my-load-balancer)|
|protocol|text|The protocol you have enabled for your load balancer|
|public_ports|integer[]|An array of public port settings for your load balancer|
|resource_type|text|Type of the lightsail resource|
|state|text|The status of your load balancer|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
|tls_policy_name|text|The name of the TLS security policy for the load balancer|
