
# Table: aws_elbv2_listeners
Information about a listener.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|alpn_policy|text[]|[TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.|
|arn|text|The Amazon Resource Name (ARN) of the listener.|
|load_balancer_arn|text|The Amazon Resource Name (ARN) of the load balancer.|
|port|integer|The port on which the load balancer is listening.|
|protocol|text|The protocol for connections from clients to the load balancer.|
|ssl_policy|text|[HTTPS or TLS listener] The security policy that defines which protocols and ciphers are supported.|
