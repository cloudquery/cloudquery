
# Table: digitalocean_load_balancer_forwarding_rules
ForwardingRule represents load balancer forwarding rules.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of digitalocean_load_balancers table (FK)|
|entry_protocol|text|The protocol used for traffic to the load balancer. The possible values are: `http`, `https`, `http2`, or `tcp`. |
|entry_port|bigint|An integer representing the port on which the load balancer instance will listen.|
|target_protocol|text|The protocol used for traffic from the load balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, or `tcp`. |
|target_port|bigint|An integer representing the port on the backend Droplets to which the load balancer will send traffic.|
|certificate_id|text|The ID of the TLS certificate used for SSL termination if enabled.|
|tls_passthrough|boolean|A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets.|
