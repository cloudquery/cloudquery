
# Table: aws_elbv1_load_balancer_listeners
The policies enabled for a listener.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|listener_instance_port|integer|The port on which the instance is listening.|
|listener_load_balancer_port|integer|The port on which the load balancer is listening.|
|listener_protocol|text|The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.|
|listener_instance_protocol|text|The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.|
|listener_ssl_certificate_id|text|The Amazon Resource Name (ARN) of the server certificate.|
|policy_names|text[]|The policies.|
